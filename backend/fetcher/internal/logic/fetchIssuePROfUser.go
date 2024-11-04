package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/contribution/model"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/google/go-github/v66/github"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"strings"
	"time"
)

const issuePRFetcherTopic = "issuePR"

func FetchIssuePROfUser(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	err = fetchWithRetry(ctx, svcContext, userId, issuePRFetcherTopic, doFetchIssuePROfUser)
	return
}

func doFetchIssuePROfUser(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	var (
		githubClient *github.Client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_API_TOKEN"))
		githubUser   *github.User
		allIssuePR   []*github.Issue
	)

	if githubUser, _, err = getGithubUserById(ctx, githubClient, userId); err != nil {
		return
	}

	if allIssuePR, err = getAllGithubIssuePRByLogin(ctx, githubClient, githubUser.GetLogin(), RoleAuthor); err != nil {
		return
	}

	if err = delAllOldContributionInCategory(ctx, svcContext, userId, model.CategoryOpenIssue); err != nil {
		return
	}

	if err = delAllOldContributionInCategory(ctx, svcContext, userId, model.CategoryOpenPullRequest); err != nil {
		return
	}

	if err = delAllOldContributionInCategory(ctx, svcContext, userId, model.CategoryMerge); err != nil {
		return
	}

	for _, githubIssuePR := range allIssuePR {
		var (
			merged bool
			repo   *github.Repository
		)
		if repo, err = getRepoByUrl(ctx, githubClient, githubIssuePR.GetRepositoryURL()); err != nil {
			return
		}

		if merged, err = checkIfMerged(ctx, githubClient, githubIssuePR, repo); err != nil {
			return
		}
		if err = pushContribution(ctx, svcContext, buildIssuePR(githubIssuePR, userId, merged, repo)); err != nil {
			return
		}
	}

	return
}

func checkIfMerged(ctx context.Context, githubClient *github.Client, githubIssuePR *github.Issue, repo *github.Repository) (merged bool, err error) {
	if githubIssuePR.IsPullRequest() {
		pr, _, err := githubClient.PullRequests.Get(ctx, repo.GetOwner().GetLogin(), repo.GetName(), githubIssuePR.GetNumber())
		if err != nil {
			logx.Error("Unexpected error when getting PR: " + err.Error())
			return false, err
		}

		merged = pr.GetMerged()
	}
	return
}

func getAllGithubIssuePRByLogin(ctx context.Context, githubClient *github.Client, login string, role string) (allIssuePR []*github.Issue, err error) {
	logx.Info("Start to fetch " + role + " " + login + "'s issues and PRs")

	opt := &github.SearchOptions{ListOptions: github.ListOptions{PerPage: 100}}
	for {
		issueSearchResult, resp, err := githubClient.Search.Issues(ctx, role+":"+login, opt)
		if err != nil {
			logx.Error("Unexpected error when searching issues: " + err.Error())
			return nil, err
		}

		issues := issueSearchResult.Issues

		allIssuePR = append(allIssuePR, issues...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	logx.Info("Successfully get all " + role + " " + login + "'s issues and PRs, size: " + string(rune(len(allIssuePR))))
	return
}

func getRepoByUrl(ctx context.Context, githubClient *github.Client, repoUrl string) (repo *github.Repository, err error) {
	var (
		split    []string
		owner    string
		repoName string
	)
	split = strings.Split(repoUrl, "/")
	owner = split[len(split)-2]
	repoName = split[len(split)-1]

	if repo, _, err = githubClient.Repositories.Get(ctx, owner, repoName); err != nil {
		logx.Error("Unexpected error when getting repo info: " + err.Error())
		return
	}

	return
}

func buildIssuePR(githubIssuePR *github.Issue, userId int64, merged bool, repo *github.Repository) (newIssuePR *model.Contribution) {
	var category string
	if githubIssuePR.IsPullRequest() {
		if merged {
			category = model.CategoryMerge
		} else {
			category = model.CategoryOpenPullRequest
		}
	} else {
		category = model.CategoryOpenIssue
	}
	newIssuePR = &model.Contribution{
		DataCreatedAt:  time.Now(),
		DataUpdatedAt:  time.Now(),
		UserId:         userId,
		RepoId:         repo.GetID(),
		Category:       category,
		Content:        githubIssuePR.GetTitle() + " " + githubIssuePR.GetBody(),
		CreatedAt:      githubIssuePR.GetCreatedAt().Time,
		UpdatedAt:      githubIssuePR.GetUpdatedAt().Time,
		ContributionId: githubIssuePR.GetID(),
	}
	return
}

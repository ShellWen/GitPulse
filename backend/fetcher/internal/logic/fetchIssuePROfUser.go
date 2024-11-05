package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/contribution/model"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/google/go-github/v66/github"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"strconv"
	"strings"
	"time"
)

const issuePRFetcherTopic = "issuePR"

func FetchIssuePROfUser(ctx context.Context, svcContext *svc.ServiceContext, userId int64, createAfter string, serachLimit int64) (err error) {
	err = doFetchIssuePROfUser(ctx, svcContext, userId, createAfter, serachLimit)
	return
}

func doFetchIssuePROfUser(ctx context.Context, svcContext *svc.ServiceContext, userId int64, createAfter string, serachLimit int64) (err error) {
	var (
		githubClient *github.Client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_API_TOKEN"))
		githubUser   *github.User
		allIssuePR   []*github.Issue
	)

	if githubUser, _, err = getGithubUserById(ctx, githubClient, userId); err != nil {
		return
	}

	logx.Info("Start fetching issues and PRs of user: ", githubUser.GetLogin())
	if allIssuePR, err = getAllGithubIssuePRByLogin(ctx, githubClient, githubUser.GetLogin(), RoleAuthor, createAfter, serachLimit); err != nil {
		return
	}
	logx.Info("Finish fetching issues and PRs of user: ", githubUser.GetLogin()+", total issues and PRs: "+string(rune(len(allIssuePR))))

	if err = delAllOldContributionInCategory(ctx, svcContext, userId, model.CategoryOpenIssue); err != nil {
		return
	}

	if err = delAllOldContributionInCategory(ctx, svcContext, userId, model.CategoryOpenPullRequest); err != nil {
		return
	}

	if err = delAllOldContributionInCategory(ctx, svcContext, userId, model.CategoryMerge); err != nil {
		return
	}

	logx.Info("Start pushing issues and PRs of user: ", githubUser.GetLogin())
	var repos map[int64]*github.Repository = make(map[int64]*github.Repository)
	for _, githubIssuePR := range allIssuePR {
		var (
			merged bool
			repo   *github.Repository
		)
		if repo, err = getRepoByUrl(ctx, githubClient, githubIssuePR.GetRepositoryURL()); err != nil {
			return
		}

		if merged, err = checkIfMerged(ctx, githubClient, githubIssuePR, repo); err != nil {
			continue
		}
		if err = pushContribution(ctx, svcContext, buildIssuePR(githubIssuePR, userId, merged, repo)); err != nil {
			continue
		}

		repos[repo.GetID()] = repo
	}

	for _, repo := range repos {
		if err = buildAndPushRepoByGithubRepo(ctx, svcContext, githubClient, repo); err != nil {
			continue
		}
	}

	logx.Info("Successfully push all update tasks of issues and PRs")
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

func getAllGithubIssuePRByLogin(ctx context.Context, githubClient *github.Client, login string, role string, createAfter string, searchLimit int64) (issues []*github.Issue, err error) {
	logx.Info("Start to fetch " + role + " " + login + "'s issues and PRs")

	var issueSearchResult *github.IssuesSearchResult

	opt := &github.SearchOptions{
		Sort:        "updated",
		Order:       "desc",
		ListOptions: github.ListOptions{PerPage: int(searchLimit), Page: 1},
	}

	if issueSearchResult, _, err = githubClient.Search.Issues(ctx, role+":"+login+" created:>="+createAfter, opt); err != nil {
		logx.Error("Unexpected error when searching issues: " + err.Error())
		return nil, err
	}

	issues = issueSearchResult.Issues

	logx.Info("Successfully get all " + role + " " + login + "'s issues and PRs, size: " + strconv.Itoa(len(issues)))
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

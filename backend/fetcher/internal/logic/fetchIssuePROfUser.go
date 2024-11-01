package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/contribution/model"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/google/go-github/v66/github"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
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

	for _, githubIssuePR := range allIssuePR {
		var merged bool
		if merged, err = checkIfMerged(ctx, githubClient, githubIssuePR); err != nil {
			return
		}
		if err = pushContribution(ctx, svcContext, buildIssuePR(ctx, svcContext, githubIssuePR, userId, merged)); err != nil {
			return
		}
	}

	return
}

func checkIfMerged(ctx context.Context, githubClient *github.Client, githubIssuePR *github.Issue) (merged bool, err error) {
	if githubIssuePR.IsPullRequest() {
		pr, _, err := githubClient.PullRequests.Get(ctx, *githubIssuePR.Repository.Owner.Login, *githubIssuePR.Repository.Name, *githubIssuePR.Number)
		if err != nil {
			logx.Error("Unexpected error when getting PR: " + err.Error())
			return false, err
		}

		merged = pr.GetMerged()
	}
	return
}

func getAllGithubIssuePRByLogin(ctx context.Context, githubClient *github.Client, login string, role string) (allIssuePR []*github.Issue, err error) {
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

	return
}

func buildIssuePR(ctx context.Context, svcContext *svc.ServiceContext, githubIssuePR *github.Issue, userId int64, merged bool) (newIssuePR *model.Contribution) {
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
		UserId:         userId,
		RepoId:         *githubIssuePR.Repository.ID,
		Category:       category,
		Content:        githubIssuePR.GetTitle() + " " + githubIssuePR.GetBody(),
		CreatedAt:      githubIssuePR.CreatedAt.Time,
		UpdatedAt:      githubIssuePR.UpdatedAt.Time,
		ContributionId: githubIssuePR.GetID(),
	}
	return
}

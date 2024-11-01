package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/ShellWen/GitPulse/repo/model"
	"github.com/google/go-github/v66/github"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

func FetchRepo(ctx context.Context, svcContext *svc.ServiceContext, repoId int64) (err error) {
	err = fetchWithRetry(ctx, svcContext, repoId, "repo", doFetchRepo)
	return
}

func doFetchRepo(ctx context.Context, svcContext *svc.ServiceContext, repoId int64) (err error) {
	var (
		githubClient *github.Client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_API_TOKEN"))
		githubRepo   *github.Repository
	)

	if githubRepo, _, err = githubClient.Repositories.GetByID(ctx, repoId); err != nil {
		logx.Error(errors.New("Unexpected error when fetching repo: " + err.Error()))
		return
	}

	if err = buildAndPushRepoByGithubRepo(ctx, svcContext, githubClient, githubRepo); err != nil {
		return
	}

	return
}

func buildAndPushRepoByGithubRepo(ctx context.Context, svcContext *svc.ServiceContext, githubClient *github.Client, githubRepo *github.Repository) (err error) {
	var (
		issueCount  int64
		prCount     int64
		commitCount int64
	)

	if issueCount, prCount, err = getGithubIssuePrCountByRepo(ctx, githubClient, githubRepo.GetOwner().GetLogin(), githubRepo.GetName()); err != nil {
		return
	}

	if commitCount, err = getGithubCommitCountByRepo(ctx, githubClient, githubRepo.GetOwner().GetLogin(), githubRepo.GetName()); err != nil {
		return
	}

	if err = pushRepo(ctx, svcContext, buildRepo(ctx, svcContext, githubRepo, issueCount, prCount, commitCount)); err != nil {
		return
	}

	return
}

func getGithubIssuePrCountByRepo(ctx context.Context, githubClient *github.Client, owner string, name string) (issueCount int64, prCount int64, err error) {
	var githubPrResp *github.Response
	var githubIssueResp *github.Response

	if _, githubPrResp, err = githubClient.PullRequests.List(ctx, owner, name, &github.PullRequestListOptions{ListOptions: github.ListOptions{PerPage: 1}}); err != nil {
		logx.Error(errors.New("Unexpected error when fetching issue count: " + err.Error()))
		return
	}

	if _, githubIssueResp, err = githubClient.Issues.ListByRepo(ctx, owner, name, &github.IssueListByRepoOptions{ListOptions: github.ListOptions{PerPage: 1}}); err != nil {
		logx.Error(errors.New("Unexpected error when fetching issue count: " + err.Error()))
		return
	}

	prCount = int64(githubPrResp.LastPage)
	issueCount = int64(githubIssueResp.LastPage) - prCount // issueCount contains prCount

	return
}

func getGithubCommitCountByRepo(ctx context.Context, githubClient *github.Client, owner string, name string) (commitCount int64, err error) {
	var githubCommitResp *github.Response

	if _, githubCommitResp, err = githubClient.Repositories.ListCommits(ctx, owner, name, &github.CommitsListOptions{ListOptions: github.ListOptions{PerPage: 1}}); err != nil {
		logx.Error(errors.New("Unexpected error when fetching commit count: " + err.Error()))
		return
	}

	commitCount = int64(githubCommitResp.LastPage)

	return
}

func buildRepo(ctx context.Context, svcContext *svc.ServiceContext, githubRepo *github.Repository, issueCount int64, prCount int64, commitCount int64) (newRepo *model.Repo) {
	newRepo = &model.Repo{
		Id:          githubRepo.GetID(),
		Name:        githubRepo.GetName(),
		StarCount:   int64(githubRepo.GetStargazersCount()),
		ForkCount:   int64(githubRepo.GetForksCount()),
		IssueCount:  issueCount,
		PrCount:     prCount,
		CommitCount: commitCount,
		Language:    githubRepo.GetLanguage(),
		Description: githubRepo.GetDescription(),
	}
	return
}

func pushRepo(ctx context.Context, svcContext *svc.ServiceContext, newRepo *model.Repo) (err error) {
	var jsonStr string

	if jsonStr, err = jsonx.MarshalToString(newRepo); err != nil {
		logx.Error(errors.New("Unexpected error when marshalling repo: " + err.Error()))
		return
	}

	if err = svcContext.KqRepoPusher.Push(ctx, jsonStr); err != nil {
		logx.Error(errors.New("Unexpected error when pushing repo: " + err.Error()))
		return
	}

	return
}

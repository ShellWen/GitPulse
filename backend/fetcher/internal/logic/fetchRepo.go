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
	"time"
)

func FetchRepo(ctx context.Context, svcContext *svc.ServiceContext, repoId int64) (err error) {
	err = doFetchRepo(ctx, svcContext, repoId)
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

	logx.Info("Successfully push repo: " + githubRepo.GetFullName())
	return
}

func buildAndPushRepoByGithubRepo(ctx context.Context, svcContext *svc.ServiceContext, githubClient *github.Client, githubRepo *github.Repository) (err error) {
	var (
		issueCount    int64
		prCount       int64
		commitCount   int64
		mergedPrCount int64
		openPrCount   int64
		commentCount  int64
		reviewCount   int64
		languages     string
	)

	if issueCount, prCount, err = getGithubIssuePrCountByRepo(ctx, githubClient, githubRepo.GetOwner().GetLogin(), githubRepo.GetName()); err != nil {
		return
	}

	if commitCount, err = getGithubCommitCountByRepo(ctx, githubClient, githubRepo.GetOwner().GetLogin(), githubRepo.GetName()); err != nil {
		return
	}

	if mergedPrCount, err = getGithubMergedPrCountByRepo(ctx, githubClient, githubRepo.GetOwner().GetLogin(), githubRepo.GetName()); err != nil {
		return
	}

	if openPrCount, err = getGithubOpenPrCountByRepo(ctx, githubClient, githubRepo.GetOwner().GetLogin(), githubRepo.GetName()); err != nil {
		return
	}

	if commentCount, err = getGithubCommentCountByRepo(ctx, githubClient, githubRepo.GetOwner().GetLogin(), githubRepo.GetName()); err != nil {
		return
	}

	if reviewCount, err = getGithubReviewCountByRepo(ctx, githubClient, githubRepo.GetOwner().GetLogin(), githubRepo.GetName()); err != nil {
		return
	}

	if languages, err = getGithubLanguagesByRepo(ctx, githubClient, githubRepo.GetOwner().GetLogin(), githubRepo.GetName()); err != nil {
		return
	}

	if err = pushRepo(ctx, svcContext, buildRepo(ctx, svcContext, githubRepo, issueCount, prCount, commitCount, mergedPrCount, openPrCount, commentCount, reviewCount, languages)); err != nil {
		return
	}

	return
}

func getGithubIssuePrCountByRepo(ctx context.Context, githubClient *github.Client, owner string, name string) (issueCount int64, prCount int64, err error) {
	var githubPrResp *github.Response
	var githubIssueResp *github.Response

	if _, githubPrResp, err = githubClient.PullRequests.List(ctx, owner, name, &github.PullRequestListOptions{ListOptions: github.ListOptions{PerPage: 1}}); err != nil {
		logx.Error(errors.New("Unexpected error when fetching issue count: " + err.Error()))
		return 0, 0, nil
	}

	if _, githubIssueResp, err = githubClient.Issues.ListByRepo(ctx, owner, name, &github.IssueListByRepoOptions{ListOptions: github.ListOptions{PerPage: 1}}); err != nil {
		logx.Error(errors.New("Unexpected error when fetching issue count: " + err.Error()))
		return 0, 0, nil
	}

	prCount = int64(githubPrResp.LastPage)
	issueCount = int64(githubIssueResp.LastPage) - prCount // issueCount contains prCount

	return
}

func getGithubOpenPrCountByRepo(ctx context.Context, githubClient *github.Client, owner string, name string) (openMergePrCount int64, err error) {
	var githubPrResp *github.Response

	if _, githubPrResp, err = githubClient.PullRequests.List(ctx, owner, name, &github.PullRequestListOptions{State: "open", ListOptions: github.ListOptions{PerPage: 1}}); err != nil {
		logx.Error(errors.New("Unexpected error when fetching open merge pr count: " + err.Error()))
		return 0, nil
	}

	openMergePrCount = int64(githubPrResp.LastPage)

	return
}

func getGithubMergedPrCountByRepo(ctx context.Context, githubClient *github.Client, owner string, name string) (mergedPrCount int64, err error) {
	var githubPrResp *github.Response

	if _, githubPrResp, err = githubClient.Search.Issues(ctx, "repo:"+owner+"/"+name+" is:pr is:merged", &github.SearchOptions{ListOptions: github.ListOptions{PerPage: 1}}); err != nil {
		logx.Error(errors.New("Unexpected error when fetching merged pr count: " + err.Error()))
		return 0, nil
	}

	mergedPrCount = int64(githubPrResp.LastPage)

	return
}

func getGithubCommentCountByRepo(ctx context.Context, githubClient *github.Client, owner string, name string) (commentCount int64, err error) {
	var githubIssueResp *github.Response

	if _, githubIssueResp, err = githubClient.Issues.ListComments(ctx, owner, name, 0, &github.IssueListCommentsOptions{ListOptions: github.ListOptions{PerPage: 1}}); err != nil {
		logx.Error(errors.New("Unexpected error when fetching comment count: " + err.Error()))
		return 0, nil
	}

	commentCount = int64(githubIssueResp.LastPage)

	return
}

func getGithubReviewCountByRepo(ctx context.Context, githubClient *github.Client, owner string, name string) (reviewCount int64, err error) {
	var githubPrResp *github.Response

	if _, githubPrResp, err = githubClient.PullRequests.ListComments(ctx, owner, name, 0, &github.PullRequestListCommentsOptions{ListOptions: github.ListOptions{PerPage: 1}}); err != nil {
		logx.Error(errors.New("Unexpected error when fetching comment count: " + err.Error()))
		return 0, nil
	}

	reviewCount = int64(githubPrResp.LastPage)

	return
}

func getGithubCommitCountByRepo(ctx context.Context, githubClient *github.Client, owner string, name string) (commitCount int64, err error) {
	var githubCommitResp *github.Response

	if _, githubCommitResp, err = githubClient.Repositories.ListCommits(ctx, owner, name, &github.CommitsListOptions{ListOptions: github.ListOptions{PerPage: 1}}); err != nil {
		logx.Error(errors.New("Unexpected error when fetching commit count: " + err.Error()))
		return 0, nil
	}

	commitCount = int64(githubCommitResp.LastPage)

	return
}

func getGithubLanguagesByRepo(ctx context.Context, githubClient *github.Client, owner string, name string) (languages string, err error) {
	var githubLanguages map[string]int

	if githubLanguages, _, err = githubClient.Repositories.ListLanguages(ctx, owner, name); err != nil {
		logx.Error(errors.New("Unexpected error when fetching languages: " + err.Error()))
		return "{}", nil
	}

	if languages, err = jsonx.MarshalToString(githubLanguages); err != nil {
		logx.Error(errors.New("Unexpected error when marshalling languages: " + err.Error()))
		return "{}", nil
	}

	return
}

func buildRepo(ctx context.Context, svcContext *svc.ServiceContext, githubRepo *github.Repository, issueCount int64, prCount int64, commitCount int64, mergedPrCount int64, openPrCount int64, commentCount int64, reviewCount int64, languages string) (newRepo *model.Repo) {
	newRepo = &model.Repo{
		DataCreatedAt: time.Now(),
		DataUpdatedAt: time.Now(),
		Id:            githubRepo.GetID(),
		Name:          githubRepo.GetName(),
		StarCount:     int64(githubRepo.GetStargazersCount()),
		ForkCount:     int64(githubRepo.GetForksCount()),
		IssueCount:    issueCount,
		PrCount:       prCount,
		CommitCount:   commitCount,
		Language:      languages,
		Description:   githubRepo.GetDescription(),
		MergedPrCount: mergedPrCount,
		OpenPrCount:   openPrCount,
		CommentCount:  commentCount,
		ReviewCount:   reviewCount,
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

	logx.Info("Successfully push repo: " + newRepo.Name)
	return
}

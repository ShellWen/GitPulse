package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/developer/model"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/google/go-github/v66/github"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"time"
)

const developerProfileFetcherTopic = "developer profile"

func FetchDeveloper(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	err = doFetchDeveloper(ctx, svcContext, userId)
	return
}

func doFetchDeveloper(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	var (
		githubClient *github.Client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_API_TOKEN"))
		githubUser   *github.User
	)

	if githubUser, _, err = getGithubUserById(ctx, githubClient, userId); err != nil {
		return
	}

	if err = buildAndPushDeveloperByGithubUser(ctx, svcContext, githubClient, githubUser); err != nil {
		return
	}

	logx.Info("Successfully push developer profile: " + githubUser.GetLogin())
	return
}

func buildAndPushDeveloperByGithubUser(ctx context.Context, svcContext *svc.ServiceContext, githubClient *github.Client, githubUser *github.User) (err error) {
	var starredRepoCount int64

	if starredRepoCount, err = getGithubStarredRepoCountByLogin(ctx, githubClient, githubUser.GetLogin()); err != nil {
		return
	}

	if err = pushDeveloperProfile(ctx, svcContext, buildDeveloperProfile(ctx, svcContext, githubUser, starredRepoCount)); err != nil {
		return
	}

	return
}

func getGithubUserById(ctx context.Context, githubClient *github.Client, userId int64) (githubUser *github.User, githubResp *github.Response, err error) {
	if githubUser, githubResp, err = githubClient.Users.GetByID(ctx, userId); err != nil {
		logx.Error(errors.New("Unexpected error when get login: " + err.Error()))
	} else {
		logx.Info("Successfully get developer profile: " + githubUser.GetLogin())
	}

	return
}

func getGithubStarredRepoCountByLogin(ctx context.Context, githubClient *github.Client, login string) (starredRepoCount int64, err error) {
	var githubStarredRepoResp *github.Response
	if _, githubStarredRepoResp, err = githubClient.Activity.ListStarred(ctx, login, &github.ActivityListStarredOptions{ListOptions: github.ListOptions{PerPage: 1}}); err != nil {
		logx.Error(errors.New("Unexpected error when fetching starred repo: " + err.Error()))
		return 0, nil
	}

	logx.Info("Successfully get starred repo count: " + login)
	starredRepoCount = int64(githubStarredRepoResp.LastPage)
	return
}

func buildDeveloperProfile(ctx context.Context, svcContext *svc.ServiceContext, githubUser *github.User, starredRepoCount int64) (newDeveloper *model.Developer) {
	newDeveloper = &model.Developer{
		DataCreatedAt:   time.Now(),
		DataUpdatedAt:   time.Now(),
		Id:              githubUser.GetID(),
		Name:            githubUser.GetName(),
		Login:           githubUser.GetLogin(),
		AvatarUrl:       githubUser.GetAvatarURL(),
		Company:         githubUser.GetCompany(),
		Location:        githubUser.GetLocation(),
		Bio:             githubUser.GetBio(),
		Blog:            githubUser.GetBlog(),
		Email:           githubUser.GetEmail(),
		CreatedAt:       githubUser.GetCreatedAt().Time,
		UpdatedAt:       githubUser.GetUpdatedAt().Time,
		TwitterUsername: githubUser.GetTwitterUsername(),
		Repos:           int64(githubUser.GetPublicRepos()),
		Following:       int64(githubUser.GetFollowing()),
		Followers:       int64(githubUser.GetFollowers()),
		Gists:           int64(githubUser.GetPublicGists()),
		Stars:           starredRepoCount,
	}
	return
}

func pushDeveloperProfile(ctx context.Context, svcContext *svc.ServiceContext, newDeveloper *model.Developer) (err error) {
	var jsonStr string

	if jsonStr, err = jsonx.MarshalToString(newDeveloper); err != nil {
		logx.Error(errors.New("Unexpected error when fetching developer profile: " + err.Error()))
		return
	}

	if err = svcContext.KqDeveloperPusher.Push(ctx, jsonStr); err != nil {
		logx.Error(errors.New("Unexpected error when fetching developer profile: " + err.Error()))
		return
	}

	return
}

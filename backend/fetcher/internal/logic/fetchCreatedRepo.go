package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/relation"
	"github.com/ShellWen/GitPulse/relation/model"
	"github.com/google/go-github/v66/github"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"os"
	"strconv"
	"time"
)

func FetchCreatedRepo(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	if err = doFetchCreatedRepo(ctx, svcContext, userId); err != nil {
		return
	}

	if err = updateCreateRepoFetchTimeOfDeveloper(ctx, svcContext, userId); err != nil {
		return
	}

	return
}

func doFetchCreatedRepo(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	var (
		githubClient *github.Client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_API_TOKEN"))
		githubUser   *github.User
		allRepos     []*github.Repository
	)

	if githubUser, _, err = getGithubUserById(ctx, githubClient, userId); err != nil {
		return
	}

	logx.Info("Start fetching created repos of user: ", githubUser.GetLogin())
	if allRepos, err = getAllGithubReposByLogin(ctx, githubClient, githubUser.GetLogin(), &github.RepositoryListByUserOptions{
		Type:        "owner",
		ListOptions: github.ListOptions{PerPage: 100},
	}); err != nil {
		return
	}
	logx.Info("Finish fetching created repos of user: ", githubUser.GetLogin()+", total created repos: "+strconv.Itoa(len(allRepos)))

	if err = delAllCreatedRepo(ctx, svcContext, userId); err != nil {
		return
	}

	logx.Info("Start pushing created repos of user: ", githubUser.GetLogin())
	for _, githubRepo := range allRepos {
		if githubRepo.GetFork() == true {
			logx.Info("Skipped a fork...")
			continue
		}

		if err = pushCreatedRepo(ctx, svcContext, buildCreatedRepo(ctx, svcContext, githubRepo, userId)); err != nil {
			continue
		}

		if err = buildAndPushRepoByGithubRepo(ctx, svcContext, githubClient, githubRepo); err != nil {
			continue
		}
	}

	logx.Info("Successfully push all update tasks of created repos")
	return
}

func delAllCreatedRepo(ctx context.Context, svcContext *svc.ServiceContext, userId int64) error {
	relationZrpcClient := svcContext.RelationRpcClient

	if delAllCreatedRepoResp, err := relationZrpcClient.DelAllCreatedRepo(ctx, &relation.DelAllCreatedRepoReq{DeveloperId: userId}); err != nil {
		logx.Error("Unexpected error when deleting all created repos: " + err.Error())
		return err
	} else if delAllCreatedRepoResp.Code != http.StatusOK {
		errMsg := "Unexpected error when deleting all created repos: " + delAllCreatedRepoResp.Message
		err = errors.New(errMsg)
		logx.Error(errMsg)
		return err
	}

	logx.Info("Successfully delete all created repos")
	return nil
}

func getLastFetchTimeOfCreatedRepo(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (lastModified string, err error) {
	var existingDeveloperResp *developer.GetDeveloperByIdResp

	developerZrpcClient := svcContext.DeveloperRpcClient
	if existingDeveloperResp, err = developerZrpcClient.GetDeveloperById(ctx, &developer.GetDeveloperByIdReq{Id: userId}); err != nil {
		logx.Error(errors.New("Unexpected error when fetching developer profile: " + err.Error()))
	} else if existingDeveloperResp.Code == http.StatusOK {
		lastModified = time.Unix(existingDeveloperResp.Developer.LastFetchCreateRepoAt, 0).Format(http.TimeFormat)
	} else {
		lastModified = ""
	}

	return
}

func getAllGithubReposByLogin(ctx context.Context, githubClient *github.Client, login string, opt *github.RepositoryListByUserOptions) (allRepos []*github.Repository, err error) {
	for {
		repos, resp, err := githubClient.Repositories.ListByUser(ctx, login, opt)
		if err != nil {
			logx.Error("Unexpected error when getting github repos: " + err.Error())
			return nil, nil
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	logx.Info("Successfully get all github repos, size: " + string(rune(len(allRepos))))
	return
}

func buildCreatedRepo(ctx context.Context, svcContext *svc.ServiceContext, githubRepo *github.Repository, userId int64) (newCreatedRepo *model.CreateRepo) {
	newCreatedRepo = &model.CreateRepo{
		DeveloperId: userId,
		RepoId:      githubRepo.GetID(),
	}

	return
}

func pushCreatedRepo(ctx context.Context, svcContext *svc.ServiceContext, newCreatedRepos *model.CreateRepo) (err error) {
	var jsonStr string

	if jsonStr, err = jsonx.MarshalToString(newCreatedRepos); err != nil {
		logx.Error(errors.New("Unexpected error when marshalling created repo: " + err.Error()))
		return
	}

	if err = svcContext.KqCreateRepoPusher.Push(ctx, jsonStr); err != nil {
		logx.Error(errors.New("Unexpected error when pushing created repo: " + err.Error()))
		return
	}

	return
}

func updateCreateRepoFetchTimeOfDeveloper(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	developerZrpcClient := svcContext.DeveloperRpcClient
	var resp *developer.GetDeveloperByIdResp
	var theDeveloper *developer.Developer

	if resp, err = developerZrpcClient.GetDeveloperById(ctx, &developer.GetDeveloperByIdReq{Id: userId}); err != nil {
		return
	}

	switch resp.Code {
	case http.StatusOK:
		theDeveloper = resp.Developer
	case http.StatusNotFound:
		err = errors.New("Developer not found")
		return
	default:
		err = errors.New("Unexpected error when getting developer: " + resp.Message)
		return
	}

	theDeveloper.LastFetchCreateRepoAt = time.Now().Unix()
	if _, err = developerZrpcClient.UpdateDeveloper(ctx, &developer.UpdateDeveloperReq{
		Id:                      userId,
		Name:                    theDeveloper.Name,
		Login:                   theDeveloper.Login,
		AvatarUrl:               theDeveloper.AvatarUrl,
		Company:                 theDeveloper.Company,
		Location:                theDeveloper.Location,
		Bio:                     theDeveloper.Bio,
		Blog:                    theDeveloper.Blog,
		Email:                   theDeveloper.Email,
		CreatedAt:               theDeveloper.CreatedAt,
		UpdatedAt:               theDeveloper.UpdatedAt,
		TwitterUsername:         theDeveloper.TwitterUsername,
		Repos:                   theDeveloper.Repos,
		Following:               theDeveloper.Following,
		Followers:               theDeveloper.Followers,
		Gists:                   theDeveloper.Gists,
		Stars:                   theDeveloper.Stars,
		LastFetchContributionAt: theDeveloper.LastFetchContributionAt,
		LastFetchFollowAt:       theDeveloper.LastFetchFollowAt,
		LastFetchStarAt:         theDeveloper.LastFetchStarAt,
		LastFetchCreateRepoAt:   theDeveloper.LastFetchCreateRepoAt,
	}); err != nil {
		return
	}

	return
}

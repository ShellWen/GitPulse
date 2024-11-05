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
	"time"
)

func FetchStarredRepo(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	if err = doFetchStarredRepo(ctx, svcContext, userId); err != nil {
		return
	}

	if err = updateStarredRepoFetchTimeOfDeveloper(ctx, svcContext, userId); err != nil {
		return
	}

	logx.Info("Successfully update fetch time of starred repo")
	return
}

func doFetchStarredRepo(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	var (
		githubClient *github.Client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_API_TOKEN"))
		githubUser   *github.User
		allRepos     []*github.StarredRepository
	)

	if githubUser, _, err = getGithubUserById(ctx, githubClient, userId); err != nil {
		return
	}

	logx.Info("Start fetching starred repo of user: ", githubUser.GetLogin())
	if allRepos, err = getAllGithubStarredReposByLogin(ctx, githubClient, githubUser.GetLogin()); err != nil {
		return
	}
	logx.Info("Finish fetching starred repo of user: ", githubUser.GetLogin()+", total starred repos: "+string(rune(len(allRepos))))

	if err = delAllOldStars(ctx, svcContext, userId); err != nil {
		return
	}

	logx.Info("Start pushing starred repo of user: ", githubUser.GetLogin())
	for _, githubRepo := range allRepos {
		if err = pushStarredRepo(ctx, svcContext, buildStarredRepo(ctx, svcContext, githubRepo, userId)); err != nil {
			continue
		}

		if err = buildAndPushRepoByGithubRepo(ctx, svcContext, githubClient, githubRepo.Repository); err != nil {
			continue
		}
	}

	logx.Info("Successfully push all update tasks of starred repo")
	return
}

func delAllOldStars(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	relationZrpcClient := svcContext.RelationRpcClient

	if delAllStarredRepoResp, err := relationZrpcClient.DelAllStarredRepo(ctx, &relation.DelAllStarredRepoReq{DeveloperId: userId}); err != nil {
		logx.Error(errors.New("Unexpected error when deleting old stars: " + err.Error()))
		return err
	} else if delAllStarredRepoResp.Code != http.StatusOK {
		logx.Error(errors.New("Unexpected error when deleting old stars: " + delAllStarredRepoResp.Message))
		return errors.New("Unexpected error when deleting old stars: " + delAllStarredRepoResp.Message)
	}

	return
}

func getAllGithubStarredReposByLogin(ctx context.Context, githubClient *github.Client, login string) (allRepos []*github.StarredRepository, err error) {
	opt := &github.ActivityListStarredOptions{ListOptions: github.ListOptions{PerPage: 100}}
	for {
		repos, resp, err := githubClient.Activity.ListStarred(ctx, login, opt)
		if err != nil {
			logx.Error(errors.New("Unexpected error when fetching starred repo: " + err.Error()))
			return nil, nil
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return
}

func getLastFetchTimeOfStarredRepo(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (lastModified string, err error) {
	var (
		existingDeveloperResp *developer.GetDeveloperByIdResp
		developerZrpcClient   = svcContext.DeveloperRpcClient
	)

	if existingDeveloperResp, err = developerZrpcClient.GetDeveloperById(ctx, &developer.GetDeveloperByIdReq{Id: userId}); err != nil {
		logx.Error(errors.New("Unexpected error when fetching developer profile: " + err.Error()))
	} else if existingDeveloperResp.Code == http.StatusOK {
		lastModified = time.Unix(existingDeveloperResp.Developer.LastFetchStarAt, 0).Format(http.TimeFormat)
	} else {
		lastModified = ""
	}

	return
}

func buildStarredRepo(ctx context.Context, svcContext *svc.ServiceContext, githubStarredRepo *github.StarredRepository, userId int64) (newStarredRepo *model.Star) {
	newStarredRepo = &model.Star{
		DeveloperId: userId,
		RepoId:      githubStarredRepo.Repository.GetID(),
	}
	return
}

func pushStarredRepo(ctx context.Context, svcContext *svc.ServiceContext, newStarredRepo *model.Star) (err error) {
	var jsonStr string

	if jsonStr, err = jsonx.MarshalToString(newStarredRepo); err != nil {
		logx.Error(errors.New("Unexpected error when fetching starred repo: " + err.Error()))
		return
	}

	if err = svcContext.KqStarPusher.Push(ctx, jsonStr); err != nil {
		logx.Error(errors.New("Unexpected error when fetching starred repo: " + err.Error()))
		return
	}

	return
}

func updateStarredRepoFetchTimeOfDeveloper(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
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

	theDeveloper.LastFetchStarAt = time.Now().Unix()
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

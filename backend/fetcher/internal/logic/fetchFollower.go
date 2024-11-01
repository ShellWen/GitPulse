package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/relation"
	"github.com/google/go-github/v66/github"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"os"
)

func FetchFollower(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	err = fetchWithRetry(ctx, svcContext, userId, "follower", doFetchFollower)
	return
}

func doFetchFollower(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	var (
		githubClient *github.Client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_API_TOKEN"))
		githubUser   *github.User
		allFollowers []*github.User
	)

	if githubUser, _, err = getGithubUserById(ctx, githubClient, userId); err != nil {
		return
	}

	if allFollowers, err = getAllGithubFollowersByLogin(ctx, githubClient, githubUser.GetLogin()); err != nil {
		return
	}

	if err = delAllOldFollowers(ctx, svcContext, userId); err != nil {
		return
	}

	for _, githubRepo := range allFollowers {
		if err = pushFollow(ctx, svcContext, buildFollow(ctx, svcContext, githubRepo.GetID(), userId)); err != nil {
			return
		}

		if err = buildAndPushDeveloperByGithubUser(ctx, svcContext, githubClient, githubRepo); err != nil {
			return
		}
	}

	return
}

func getAllGithubFollowersByLogin(ctx context.Context, githubClient *github.Client, login string) (allFollowers []*github.User, err error) {
	opt := &github.ListOptions{PerPage: 100}
	for {
		followers, resp, err := githubClient.Users.ListFollowers(ctx, login, opt)
		if err != nil {
			logx.Error(errors.New("Unexpected error when fetching followers: " + err.Error()))
			return nil, err
		}
		allFollowers = append(allFollowers, followers...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return
}

func delAllOldFollowers(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	relationZrpcClient := relation.NewRelation(svcContext.RpcClient)

	if delAllFollowersResp, err := relationZrpcClient.DelAllFollower(ctx, &relation.DelAllFollowerReq{DeveloperId: userId}); err != nil {
		logx.Error(errors.New("Unexpected error when deleting old followers: " + err.Error()))
		return err
	} else if delAllFollowersResp.Code != http.StatusOK {
		logx.Error(errors.New("Unexpected error when deleting old followers: " + delAllFollowersResp.Message))
		return errors.New("Unexpected error when deleting old followers: " + delAllFollowersResp.Message)
	}

	return
}

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
	if err = doFetchFollower(ctx, svcContext, userId); err != nil {
		return
	}

	if err = updateFollowFetchTimeOfDeveloper(ctx, svcContext, userId); err != nil {
		return
	}

	logx.Info("Successfully update fetch time of follower")
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

	logx.Info("Start fetching followers of user: ", githubUser.GetLogin())
	if allFollowers, err = getAllGithubFollowersByLogin(ctx, githubClient, githubUser.GetLogin()); err != nil {
		return
	}
	logx.Info("Finish fetching followers of user: ", githubUser.GetLogin()+", total followers: "+string(rune(len(allFollowers))))

	if err = delAllOldFollowers(ctx, svcContext, userId); err != nil {
		return
	}

	logx.Info("Start pushing followers of user: ", githubUser.GetLogin())
	for _, githubRepo := range allFollowers {
		if err = pushFollow(ctx, svcContext, buildFollow(ctx, svcContext, githubRepo.GetID(), userId)); err != nil {
			continue
		}

		if err = buildAndPushDeveloperByGithubUser(ctx, svcContext, githubClient, githubRepo); err != nil {
			continue
		}
	}

	logx.Info("Successfully push all update tasks of followers")
	return
}

func getAllGithubFollowersByLogin(ctx context.Context, githubClient *github.Client, login string) (allFollowers []*github.User, err error) {
	opt := &github.ListOptions{PerPage: 100}
	for {
		followers, resp, err := githubClient.Users.ListFollowers(ctx, login, opt)
		if err != nil {
			logx.Error(errors.New("Unexpected error when fetching followers: " + err.Error()))
			return nil, nil
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
	relationZrpcClient := svcContext.RelationRpcClient

	if delAllFollowersResp, err := relationZrpcClient.DelAllFollower(ctx, &relation.DelAllFollowerReq{DeveloperId: userId}); err != nil {
		logx.Error(errors.New("Unexpected error when deleting old followers: " + err.Error()))
		return err
	} else if delAllFollowersResp.Code != http.StatusOK {
		logx.Error(errors.New("Unexpected error when deleting old followers: " + delAllFollowersResp.Message))
		return errors.New("Unexpected error when deleting old followers: " + delAllFollowersResp.Message)
	}

	return
}

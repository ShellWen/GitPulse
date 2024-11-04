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

func FetchFollowing(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	err = fetchWithRetry(ctx, svcContext, userId, "following", doFetchFollowing)
	return
}

func doFetchFollowing(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	var (
		githubClient *github.Client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_API_TOKEN"))
		githubUser   *github.User
		allFollowing []*github.User
	)

	if githubUser, _, err = getGithubUserById(ctx, githubClient, userId); err != nil {
		return
	}

	logx.Info("Start fetching following of user: ", githubUser.GetLogin())
	if allFollowing, err = getAllGithubFollowingByLogin(ctx, githubClient, githubUser.GetLogin()); err != nil {
		return
	}
	logx.Info("Finish fetching following of user: ", githubUser.GetLogin()+", total following: "+string(rune(len(allFollowing))))

	if err = delAllOldFollowing(ctx, svcContext, userId); err != nil {
		return
	}

	logx.Info("Start pushing following of user: ", githubUser.GetLogin())
	for _, githubRepo := range allFollowing {
		if err = pushFollow(ctx, svcContext, buildFollow(ctx, svcContext, githubRepo.GetID(), userId)); err != nil {
			return
		}

		if err = buildAndPushDeveloperByGithubUser(ctx, svcContext, githubClient, githubRepo); err != nil {
			return
		}
	}

	logx.Info("Successfully push all update tasks of following")
	return
}

func getAllGithubFollowingByLogin(ctx context.Context, githubClient *github.Client, login string) (allFollowing []*github.User, err error) {
	opt := &github.ListOptions{PerPage: 100}
	for {
		following, resp, err := githubClient.Users.ListFollowing(ctx, login, opt)
		if err != nil {
			logx.Error(errors.New("Unexpected error when fetching following: " + err.Error()))
			return nil, err
		}
		allFollowing = append(allFollowing, following...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return
}

func delAllOldFollowing(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	relationZrpcClient := svcContext.RelationRpcClient

	if delAllFollowingResp, err := relationZrpcClient.DelAllFollowing(ctx, &relation.DelAllFollowingReq{DeveloperId: userId}); err != nil {
		logx.Error(errors.New("Unexpected error when deleting old following: " + err.Error()))
		return err
	} else if delAllFollowingResp.Code != http.StatusOK {
		logx.Error(errors.New("Unexpected error when deleting old following: " + delAllFollowingResp.Message))
		return errors.New("Unexpected error when deleting old following: " + delAllFollowingResp.Message)
	}

	return
}

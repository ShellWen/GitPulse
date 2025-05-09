package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/common/tasks"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/relation"
	"github.com/ShellWen/GitPulse/relation/model"
	"github.com/google/go-github/v66/github"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"os"
	"strconv"
)

func FetchFork(ctx context.Context, svcContext *svc.ServiceContext, repoId int64) (err error) {
	err = doFetchFork(ctx, svcContext, repoId)
	return
}

func doFetchFork(ctx context.Context, svcContext *svc.ServiceContext, repoId int64) (err error) {
	var (
		githubClient *github.Client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_API_TOKEN"))
		originalRepo *github.Repository
		allForks     []*github.Repository
	)

	if originalRepo, _, err = githubClient.Repositories.GetByID(ctx, repoId); err != nil {
		return
	}

	if err = delAllOldForks(ctx, svcContext, repoId); err != nil {
		return
	}

	logx.Info("Start fetching forks of repo: ", originalRepo.GetFullName())
	if allForks, err = getAllGithubForksByRepo(ctx, githubClient, originalRepo.GetOwner().GetLogin(), originalRepo.GetName()); err != nil {
		return
	}
	logx.Info("Finish fetching forks of repo: ", originalRepo.GetFullName()+", total forks: "+strconv.Itoa(len(allForks)))

	logx.Info("Start pushing forks of repo: ", originalRepo.GetFullName())
	for _, githubRepo := range allForks {
		if err = pushFork(ctx, svcContext, buildFork(ctx, svcContext, githubRepo, repoId)); err != nil {
			continue
		}

		if err = buildAndPushRepoByGithubRepo(ctx, svcContext, githubClient, githubRepo); err != nil {
			continue
		}
	}

	if err = pushFetchForkCompleted(ctx, svcContext, repoId); err != nil {
		return
	}

	logx.Info("Successfully push all update tasks of forks")
	return
}

func getAllGithubForksByRepo(ctx context.Context, githubClient *github.Client, login string, repoName string) (allForks []*github.Repository, err error) {
	opt := &github.RepositoryListForksOptions{ListOptions: github.ListOptions{PerPage: 100}}
	for {
		repos, resp, err := githubClient.Repositories.ListForks(ctx, login, repoName, opt)
		if err != nil {
			logx.Error(errors.New("Unexpected error when fetching forks: " + err.Error()))
			return nil, nil
		}
		allForks = append(allForks, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return
}

func delAllOldForks(ctx context.Context, svcContext *svc.ServiceContext, repoId int64) (err error) {
	relationZrpcClient := svcContext.RelationRpcClient

	if delAllForkResp, err := relationZrpcClient.DelAllFork(ctx, &relation.DelAllForkReq{OriginalRepoId: repoId}); err != nil {
		logx.Error(errors.New("Unexpected error when deleting old forks: " + err.Error()))
		return err
	} else if delAllForkResp.Code != http.StatusOK {
		logx.Error(errors.New("Unexpected error when deleting old forks: " + delAllForkResp.Message))
		return errors.New("Unexpected error when deleting old forks: " + delAllForkResp.Message)
	}

	logx.Info("Successfully delete all old forks of repo: " + strconv.FormatInt(repoId, 10))
	return
}

func buildFork(ctx context.Context, svcContext *svc.ServiceContext, githubRepo *github.Repository, originalRepoId int64) (newFork *model.Fork) {
	newFork = &model.Fork{
		OriginalRepoId: originalRepoId,
		ForkRepoId:     githubRepo.GetID(),
	}
	return
}

func pushFork(ctx context.Context, svcContext *svc.ServiceContext, newFork *model.Fork) (err error) {
	var jsonStr string

	if jsonStr, err = jsonx.MarshalToString(newFork); err != nil {
		logx.Error(errors.New("Unexpected error when marshalling fork: " + err.Error()))
		return
	}

	if err = svcContext.KqForkPusher.Push(ctx, jsonStr); err != nil {
		logx.Error(errors.New("Unexpected error when pushing fork: " + err.Error()))
		return
	}

	return
}

func pushFetchForkCompleted(ctx context.Context, svcContext *svc.ServiceContext, repoId int64) (err error) {
	if err = pushFork(ctx, svcContext, &model.Fork{
		DataId:         tasks.FetchForkCompletedDataId,
		OriginalRepoId: repoId,
	}); err != nil {
		logx.Error("Push fetch fork completed error: ", err)
		return
	}

	return
}

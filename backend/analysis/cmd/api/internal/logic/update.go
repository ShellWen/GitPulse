package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/svc"
	"github.com/ShellWen/GitPulse/common/tasks"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/contribution"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/relation"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strconv"
	"time"
)

func checkIfNeedUpdateContribution(ctx context.Context, svcCtx *svc.ServiceContext, id int64) (needUpdate bool, err error) {
	var (
		developerRpcClient = svcCtx.DeveloperRpcClient
		resp               *developer.GetDeveloperByIdResp
	)

	if resp, err = developerRpcClient.GetDeveloperById(ctx, &developer.GetDeveloperByIdReq{Id: id}); err != nil {
		return
	}

	switch resp.Code {
	case http.StatusOK:
		if time.Now().Unix()-resp.Developer.LastFetchContributionAt > int64(time.Hour.Seconds()*24) {
			needUpdate = true
		}
	case http.StatusNotFound:
		needUpdate = true
	default:
		err = errors.New("Unexpected error when getting developer: " + resp.Message)
	}

	return
}

func checkIfNeedUpdateCreatedRepo(ctx context.Context, svcCtx *svc.ServiceContext, id int64) (needUpdate bool, err error) {
	var (
		developerRpcClient = svcCtx.DeveloperRpcClient
		resp               *developer.GetDeveloperByIdResp
	)

	if resp, err = developerRpcClient.GetDeveloperById(ctx, &developer.GetDeveloperByIdReq{Id: id}); err != nil {
		return
	}

	switch resp.Code {
	case http.StatusOK:
		if time.Now().Unix()-resp.Developer.LastFetchCreateRepoAt > int64(time.Hour.Seconds()*24) {
			needUpdate = true
		}
	case http.StatusNotFound:
		needUpdate = true
	default:
		err = errors.New("Unexpected error when getting developer: " + resp.Message)
	}

	return
}

func checkIfNeedUpdateDeveloper(ctx context.Context, svcCtx *svc.ServiceContext, id int64) (needUpdate bool, err error) {
	var (
		developerRpcClient = svcCtx.DeveloperRpcClient
		resp               *developer.GetDeveloperByIdResp
	)

	if resp, err = developerRpcClient.GetDeveloperById(ctx, &developer.GetDeveloperByIdReq{Id: id}); err != nil {
		return
	}

	switch resp.Code {
	case http.StatusOK:
		if time.Now().Unix()-resp.Developer.UpdatedAt > int64(time.Hour.Seconds()*24) {
			needUpdate = true
		}
	case http.StatusNotFound:
		needUpdate = true
	default:
		err = errors.New("Unexpected error when getting developer: " + resp.Message)
	}

	return
}

func updateDeveloper(ctx context.Context, svcCtx *svc.ServiceContext, id int64) (err error) {
	var (
		task      *asynq.Task
		jsonStr   string
		blockResp *developer.BlockUntilDeveloperUpdatedResp
	)

	if task, err = tasks.NewFetcherTask(tasks.FetchDeveloper, id); err != nil {
		return
	}
	if jsonStr, err = jsonx.MarshalToString(task); err != nil {
		return
	}

	if _, err = svcCtx.AsynqClient.Enqueue(task, asynq.TaskID(strconv.Itoa(tasks.FetchDeveloper)+","+strconv.Itoa(int(id)))); err != nil {
		return
	}
	logx.Info("Successfully pushed task ", jsonStr, " to fetcher, waiting for developer updated")

	if blockResp, err = svcCtx.DeveloperRpcClient.BlockUntilDeveloperUpdated(ctx, &developer.BlockUntilDeveloperUpdatedReq{Id: id}); err != nil {
		return
	}

	if blockResp.Code != http.StatusOK {
		err = errors.New(blockResp.Message)
	}

	logx.Info("Developer successfully updated")

	return
}

func updateContribution(ctx context.Context, svcCtx *svc.ServiceContext, id int64) (err error) {
	var (
		contributionRpcClient = svcCtx.ContributionRpcClient
		fetcherTask           *asynq.Task
		blockResp             *contribution.BlockUntilAllUpdatedResp
	)

	if fetcherTask, err = tasks.NewFetcherTask(tasks.FetchContributionOfUser, id); err != nil {
		return
	}

	if _, err = svcCtx.AsynqClient.Enqueue(fetcherTask, asynq.TaskID(strconv.Itoa(tasks.FetchContributionOfUser)+","+strconv.Itoa(int(id)))); err != nil {
		return
	}

	logx.Info("Successfully pushed task ", fetcherTask.Payload(), " to fetcher, waiting for contribution updated")

	if blockResp, err = contributionRpcClient.BlockUntilAllUpdated(ctx, &contribution.BlockUntilAllUpdatedReq{
		UserId: id,
	}); err != nil {
		return
	}

	if blockResp.Code != http.StatusOK {
		err = errors.New(blockResp.Message)
	}

	logx.Info("Contribution successfully updated")

	return
}

func updateCreatedRepo(ctx context.Context, svcCtx *svc.ServiceContext, id int64) (err error) {
	var (
		relationRpcClient = svcCtx.RelationRpcClient
		fetcherTask       *asynq.Task
		blockResp         *relation.BlockUntilCreatedRepoUpdatedResp
	)

	if fetcherTask, err = tasks.NewFetcherTask(tasks.FetchCreatedRepo, id); err != nil {
		return
	}

	if _, err = svcCtx.AsynqClient.Enqueue(fetcherTask, asynq.TaskID(strconv.Itoa(tasks.FetchCreatedRepo)+","+strconv.Itoa(int(id)))); err != nil {
		return
	}

	if blockResp, err = relationRpcClient.BlockUntilCreatedRepoUpdated(ctx, &relation.BlockUntilCreatedRepoUpdatedReq{
		Id: id,
	}); err != nil {
		return
	}

	if blockResp.Code != http.StatusOK {
		err = errors.New(blockResp.Message)
	}

	return
}

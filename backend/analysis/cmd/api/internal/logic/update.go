package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/svc"
	"github.com/ShellWen/GitPulse/common/message"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/contribution"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/relation"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
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
		task      *message.FetcherTask
		jsonStr   string
		blockResp *developer.BlockUntilDeveloperUpdatedResp
	)

	task = &message.FetcherTask{Type: message.FetchDeveloper, Id: id}
	if jsonStr, err = jsonx.MarshalToString(task); err != nil {
		return
	}

	if err = svcCtx.KqFetcherTaskPusher.Push(ctx, jsonStr); err != nil {
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
		fetcherTask           = message.FetcherTask{Type: message.FetchContributionOfUser, Id: id}
		taskStr               string
		blockResp             *contribution.BlockUntilAllUpdatedResp
	)

	if taskStr, err = jsonx.MarshalToString(fetcherTask); err != nil {
		return
	}

	if err = svcCtx.KqFetcherTaskPusher.Push(ctx, taskStr); err != nil {
		return
	}

	logx.Info("Successfully pushed task ", taskStr, " to fetcher, waiting for contribution updated")

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
		fetcherTask       = message.FetcherTask{Type: message.FetchCreatedRepo, Id: id}
		taskStr           string
		blockResp         *relation.BlockUntilCreatedRepoUpdatedResp
	)

	if taskStr, err = jsonx.MarshalToString(fetcherTask); err != nil {
		return
	}

	if err = svcCtx.KqFetcherTaskPusher.Push(ctx, taskStr); err != nil {
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

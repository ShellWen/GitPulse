package logic

import (
	"context"
	"encoding/json"
	"github.com/ShellWen/GitPulse/api_processor/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"github.com/zeromicro/x/errors"
	"net/http"
)

func GetDeveloper(ctx context.Context, svcCtx *svc.ServiceContext, developerId int64) (string, error) {
	updateResp, err := svcCtx.DeveloperRpcClient.UpdateDeveloper(ctx, &developer.UpdateDeveloperReq{Id: developerId})
	if err != nil {
		return MustBuildErrData(http.StatusInternalServerError, err.Error()), err
	} else if updateResp.Code != http.StatusOK {
		return MustBuildErrData(int(updateResp.Code), updateResp.Message), errors.New(int(updateResp.Code), updateResp.Message)
	}

	getResp, err := svcCtx.DeveloperRpcClient.GetDeveloperById(ctx, &developer.GetDeveloperByIdReq{Id: developerId})
	if err != nil {
		return MustBuildErrData(http.StatusInternalServerError, err.Error()), err
	} else if getResp.Code != http.StatusOK {
		return MustBuildErrData(int(getResp.Code), getResp.Message), errors.New(int(getResp.Code), getResp.Message)
	}

	data, err := json.Marshal(getResp.Developer)
	if err != nil {
		return MustBuildErrData(http.StatusInternalServerError, err.Error()), err
	}

	return string(data), nil
}

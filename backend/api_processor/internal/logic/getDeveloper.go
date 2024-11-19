package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ShellWen/GitPulse/api_processor/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"net/http"
)

func GetDeveloper(ctx context.Context, svcCtx *svc.ServiceContext, developerId int64) ([]byte, error) {
	updateResp, err := svcCtx.DeveloperRpcClient.UpdateDeveloper(ctx, &developer.UpdateDeveloperReq{Id: developerId})
	if err != nil {
		return nil, err
	} else if updateResp.Code != http.StatusOK {
		return nil, errors.New(updateResp.Message)
	}

	getResp, err := svcCtx.DeveloperRpcClient.GetDeveloperById(ctx, &developer.GetDeveloperByIdReq{Id: developerId})
	if err != nil {
		return nil, err
	} else if getResp.Code != http.StatusOK {
		return nil, errors.New(getResp.Message)
	}

	data, err := json.Marshal(getResp.Developer)
	if err != nil {
		return nil, err
	}

	return data, nil
}

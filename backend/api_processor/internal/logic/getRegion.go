package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/analysis"
	"github.com/ShellWen/GitPulse/api_processor/internal/svc"
	"net/http"
)

func GetRegion(ctx context.Context, svcCtx *svc.ServiceContext, developerId int64) ([]byte, error) {
	updateResp, err := svcCtx.AnalysisRpcClient.UpdateRegion(ctx, &analysis.UpdateAnalysisReq{DeveloperId: developerId})
	if err != nil {
		return nil, err
	} else if updateResp.Code != http.StatusOK {
		return nil, errors.New(updateResp.Message)
	}

	getResp, err := svcCtx.AnalysisRpcClient.GetRegion(ctx, &analysis.GetAnalysisReq{DeveloperId: developerId})
	if err != nil {
		return nil, err
	} else if getResp.Code != http.StatusOK {
		return nil, errors.New(getResp.Message)
	}

	data, err := json.Marshal(getResp.Region)
	if err != nil {
		return nil, err
	}

	return data, nil
}

package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/analysis"
	"github.com/ShellWen/GitPulse/api_processor/internal/svc"
	"net/http"
)

func GetSummary(ctx context.Context, svcCtx *svc.ServiceContext, developerId int64) ([]byte, error) {
	updateResp, err := svcCtx.AnalysisRpcClient.UpdateSummary(ctx, &analysis.UpdateAnalysisReq{DeveloperId: developerId})
	if err != nil {
		return nil, err
	} else if updateResp.Code != http.StatusOK {
		return nil, errors.New(updateResp.Message)
	}

	getResp, err := svcCtx.AnalysisRpcClient.GetSummary(ctx, &analysis.GetAnalysisReq{DeveloperId: developerId})
	if err != nil {
		return nil, err
	} else if getResp.Code != http.StatusOK {
		return nil, errors.New(getResp.Message)
	}

	data, err := json.Marshal(getResp.Summary)
	if err != nil {
		return nil, err
	}

	return data, nil
}

package logic

import (
	"context"
	"encoding/json"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/analysis"
	"github.com/ShellWen/GitPulse/api_processor/internal/svc"
	"github.com/zeromicro/x/errors"
	"net/http"
)

func GetLanguage(ctx context.Context, svcCtx *svc.ServiceContext, developerId int64) (string, error) {
	updateResp, err := svcCtx.AnalysisRpcClient.UpdateLanguage(ctx, &analysis.UpdateAnalysisReq{DeveloperId: developerId})
	if err != nil {
		return MustBuildErrData(http.StatusInternalServerError, err.Error()), err
	} else if updateResp.Code != http.StatusOK {
		return MustBuildErrData(int(updateResp.Code), updateResp.Message), errors.New(int(updateResp.Code), updateResp.Message)
	}

	getResp, err := svcCtx.AnalysisRpcClient.GetLanguages(ctx, &analysis.GetAnalysisReq{DeveloperId: developerId})
	if err != nil {
		return MustBuildErrData(http.StatusInternalServerError, err.Error()), err
	} else if getResp.Code != http.StatusOK {
		return MustBuildErrData(int(getResp.Code), updateResp.Message), errors.New(int(getResp.Code), getResp.Message)
	}

	data, err := json.Marshal(getResp.Languages)
	if err != nil {
		return MustBuildErrData(http.StatusInternalServerError, err.Error()), err
	}

	return string(data), nil
}

package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/types"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/analysis"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRegionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRegionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRegionLogic {
	return &GetRegionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRegionLogic) GetRegion(req *types.GetRegionReq) (resp *types.GetRegionResp, err error) {
	if resp, err = l.doGetRegion(req); err != nil {
		logx.Error(err)
		return
	}
	return
}

func (l *GetRegionLogic) doGetRegion(req *types.GetRegionReq) (resp *types.GetRegionResp, err error) {
	var (
		id         int64
		region     string
		confidence float64
	)

	if id, err = customGithub.GetIdByLogin(l.ctx, req.Login); err != nil {
		return
	}

	if region, confidence, err = l.getRegionFromRpc(id); err != nil {
		return
	}

	resp = &types.GetRegionResp{
		Region: types.Region{
			Id:         id,
			Region:     region,
			Confidence: confidence,
		},
	}

	return
}

func (l *GetRegionLogic) getRegionFromRpc(id int64) (region string, confidence float64, err error) {
	var (
		analysisRpcClient = l.svcCtx.AnalysisRpcClient
		rpcUpdateResp     *analysis.UpdateAnalysisResp
		rpcGetResp        *analysis.GetRegionResp
	)

	if rpcUpdateResp, err = analysisRpcClient.UpdateRegion(l.ctx, &analysis.UpdateAnalysisReq{
		DeveloperId: id,
	}); err != nil || rpcUpdateResp.Code != http.StatusOK {
		return
	}

	if rpcGetResp, err = analysisRpcClient.GetRegion(l.ctx, &analysis.GetAnalysisReq{
		DeveloperId: id,
	}); err != nil || rpcGetResp.Code != http.StatusOK {
		return
	}

	region = rpcGetResp.Region.GetRegion()
	confidence = rpcGetResp.Region.GetConfidence()

	return
}

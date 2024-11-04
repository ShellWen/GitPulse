package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/analysis"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	"net/http"
	"time"

	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/types"

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
		analysisRpcClient = analysis.NewAnalysis(l.svcCtx.RpcClient)
		rpcResp           *analysis.GetRegionResp
	)

	if rpcResp, err = analysisRpcClient.GetRegion(l.ctx, &analysis.GetAnalysisReq{
		DeveloperId: id,
	}); err != nil {
		return
	}

	switch rpcResp.Code {
	case http.StatusOK:
		logx.Info("Found in local cache")
		if time.Now().Unix()-rpcResp.Region.GetDataUpdatedAt() < int64(time.Hour.Seconds()*24) {
			break
		}
		logx.Info("Local cache expired, fetching from github")
		fallthrough
	case http.StatusNotFound:
		if err = l.updateRegion(id); err != nil {
			return
		}
		if rpcResp, err = analysisRpcClient.GetRegion(l.ctx, &analysis.GetAnalysisReq{
			DeveloperId: id,
		}); err != nil {
			return
		}
		fallthrough
	default:
		if rpcResp.Code != http.StatusOK {
			err = errors.New(rpcResp.Message)
			return
		}
	}

	region = rpcResp.Region.GetRegion()
	confidence = rpcResp.Region.GetConfidence()

	return
}

func (l *GetRegionLogic) updateRegion(id int64) (err error) {
	var (
		analysisRpcClient  = analysis.NewAnalysis(l.svcCtx.RpcClient)
		updateAnalysisResp *analysis.UpdateAnalysisResp
	)

	if updateAnalysisResp, err = analysisRpcClient.UpdateRegion(l.ctx, &analysis.UpdateAnalysisReq{
		DeveloperId: id,
	}); err != nil {
		return
	}

	switch updateAnalysisResp.Code {
	case http.StatusOK:
	default:
		err = errors.New(updateAnalysisResp.Message)
	}

	return
}

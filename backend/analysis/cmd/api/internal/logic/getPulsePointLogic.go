package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/analysis"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	"net/http"
	"time"

	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const UpdateContributionLimit = 10

type GetPulsePointLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPulsePointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPulsePointLogic {
	return &GetPulsePointLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPulsePointLogic) GetPulsePoint(req *types.GetPulsePointReq) (resp *types.GetPulsePointResp, err error) {
	if resp, err = l.doGetPulsePoint(req); err != nil {
		logx.Error(err)
		return
	}
	return
}

func (l *GetPulsePointLogic) doGetPulsePoint(req *types.GetPulsePointReq) (resp *types.GetPulsePointResp, err error) {
	var (
		id         int64
		pulsePoint float64
		updatedAt  time.Time
	)

	if id, err = customGithub.GetIdByLogin(l.ctx, req.Login); err != nil {
		return
	}

	if pulsePoint, updatedAt, err = l.getPulsePointFromRpc(id); err != nil {
		return
	}

	resp = &types.GetPulsePointResp{
		PulsePoint: types.PulsePoint{
			Id:         id,
			PulsePoint: pulsePoint,
			UpdatedAt:  updatedAt.Format(time.RFC3339),
		},
	}

	return
}

func (l *GetPulsePointLogic) getPulsePointFromRpc(id int64) (pulsePoint float64, updatedAt time.Time, err error) {
	var (
		analysisRpcClient = l.svcCtx.AnalysisRpcClient
		rpcUpdateResp     *analysis.UpdateAnalysisResp
		rpcGetResp        *analysis.GetPulsePointResp
	)

	if rpcUpdateResp, err = analysisRpcClient.UpdatePulsePoint(l.ctx, &analysis.UpdateAnalysisReq{
		DeveloperId: id,
	}); err != nil || rpcUpdateResp.Code != http.StatusOK {
		return
	}

	if rpcGetResp, err = analysisRpcClient.GetPulsePoint(l.ctx, &analysis.GetAnalysisReq{
		DeveloperId: id,
	}); err != nil || rpcGetResp.Code != http.StatusOK {
		return
	}

	pulsePoint = rpcGetResp.PulsePoint.PulsePoint
	updatedAt = time.Unix(rpcGetResp.PulsePoint.DataUpdatedAt, 0)

	return
}

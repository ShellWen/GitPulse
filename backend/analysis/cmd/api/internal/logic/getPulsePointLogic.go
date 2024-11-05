package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/analysis"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	"github.com/ShellWen/GitPulse/common/message"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/contribution"
	"github.com/zeromicro/go-zero/core/jsonx"
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
		rpcResp           *analysis.GetPulsePointResp
	)

	if rpcResp, err = analysisRpcClient.GetPulsePoint(l.ctx, &analysis.GetAnalysisReq{
		DeveloperId: id,
	}); err != nil {
		return
	}

	switch rpcResp.Code {
	case http.StatusOK:
		logx.Info("Found in local cache")
		if time.Now().Unix()-rpcResp.PulsePoint.GetDataUpdatedAt() < int64(time.Hour.Seconds()*24) {
			break
		}
		logx.Info("Local cache expired, fetching from github")
		fallthrough
	case http.StatusNotFound:
		if err = l.updatePulsePoint(id); err != nil {
			return
		}
		if rpcResp, err = analysisRpcClient.GetPulsePoint(l.ctx, &analysis.GetAnalysisReq{
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

	pulsePoint = rpcResp.PulsePoint.PulsePoint
	updatedAt = time.Unix(rpcResp.PulsePoint.DataUpdatedAt, 0)

	return
}

func (l *GetPulsePointLogic) updatePulsePoint(id int64) (err error) {
	var (
		analysisRpcClient  = l.svcCtx.AnalysisRpcClient
		updateAnalysisResp *analysis.UpdateAnalysisResp
	)

	// contribution needed for pulse point
	if err = l.updateContribution(id); err != nil {
		return
	}

	if updateAnalysisResp, err = analysisRpcClient.UpdatePulsePoint(l.ctx, &analysis.UpdateAnalysisReq{
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

func (l *GetPulsePointLogic) updateContribution(id int64) (err error) {
	var (
		contributionRpcClient = l.svcCtx.ContributionRpcClient
		fetcherTask           = message.FetcherTask{Type: message.FetchContributionOfUser, Id: id}
		taskStr               string
		blockResp             *contribution.BlockUntilAllUpdatedResp
	)

	if taskStr, err = jsonx.MarshalToString(fetcherTask); err != nil {
		return
	}

	if err = l.svcCtx.KqFetcherTaskPusher.Push(l.ctx, taskStr); err != nil {
		return
	}

	if blockResp, err = contributionRpcClient.BlockUntilAllUpdated(l.ctx, &contribution.BlockUntilAllUpdatedReq{
		UserId: id,
	}); err != nil {
		return
	}

	if blockResp.Code != http.StatusOK {
		err = errors.New(blockResp.Message)
	}

	return
}

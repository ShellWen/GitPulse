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
		needUpdate         bool
		analysisRpcClient  = l.svcCtx.AnalysisRpcClient
		updateAnalysisResp *analysis.UpdateAnalysisResp
	)

	if l.svcCtx.PulsePointUpdatedChan[id] == nil {
		l.svcCtx.PulsePointUpdatedChan[id] = make(chan struct{})
	}

	if l.svcCtx.PulsePointUpdating {
		<-l.svcCtx.PulsePointUpdatedChan[id]
		return
	} else {
		l.svcCtx.PulsePointUpdating = true
		defer func() {
			l.svcCtx.PulsePointUpdating = false
			for stillHasBlock := true; stillHasBlock; {
				select {
				case l.svcCtx.PulsePointUpdatedChan[id] <- struct{}{}:
				default:
					stillHasBlock = false
				}
			}
		}()
	}

	if needUpdate, err = checkIfNeedUpdateContribution(l.ctx, l.svcCtx, id); err != nil {
		return
	} else if needUpdate {
		if err = updateContribution(l.ctx, l.svcCtx, id); err != nil {
			return
		}
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

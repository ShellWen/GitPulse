package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/model"
	"net/http"

	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPulsePointLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPulsePointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPulsePointLogic {
	return &GetPulsePointLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPulsePointLogic) GetPulsePoint(in *pb.GetAnalysisReq) (resp *pb.GetPulsePointResp, err error) {
	var pulsePoint *model.PulsePoint
	if pulsePoint, err = l.svcCtx.PulsePointModel.FindOneByDeveloperId(l.ctx, in.DeveloperId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetPulsePointResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetPulsePointResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetPulsePointResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			PulsePoint: &pb.PulsePoint{
				DataId:        pulsePoint.DataId,
				DataCreatedAt: pulsePoint.DataCreatedAt.Unix(),
				DataUpdatedAt: pulsePoint.DataUpdatedAt.Unix(),
				PulsePoint:    pulsePoint.PulsePoint,
			},
		}
	}

	err = nil
	return
}

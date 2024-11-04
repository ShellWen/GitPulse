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

type GetRegionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRegionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRegionLogic {
	return &GetRegionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRegionLogic) GetRegion(in *pb.GetAnalysisReq) (resp *pb.GetRegionResp, err error) {
	var region *model.Region
	if region, err = l.svcCtx.RegionModel.FindOneByDeveloperId(l.ctx, in.DeveloperId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetRegionResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetRegionResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetRegionResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Region: &pb.Region{
				DataId:        region.DataId,
				DataCreatedAt: region.DataCreatedAt.Unix(),
				DataUpdatedAt: region.DataUpdatedAt.Unix(),
				Region:        region.Region,
				Confidence:    region.Confidence,
			},
		}
	}

	err = nil
	return
}

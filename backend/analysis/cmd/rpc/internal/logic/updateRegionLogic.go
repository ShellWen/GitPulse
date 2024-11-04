package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/analysis/model"
	"net/http"

	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRegionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRegionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRegionLogic {
	return &UpdateRegionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRegionLogic) UpdateRegion(in *pb.UpdateAnalysisReq) (resp *pb.UpdateAnalysisResp, err error) {
	if err = l.doUpdateRegion(in.DeveloperId); err != nil {
		resp = &pb.UpdateAnalysisResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.UpdateAnalysisResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}
	return
}

func (l *UpdateRegionLogic) doUpdateRegion(id int64) (err error) {
	var (
		regionItem *model.Region
		region     string
		confidence float64
	)

	// TODO: Get region

	if regionItem, err = l.svcCtx.RegionModel.FindOneByDeveloperId(l.ctx, id); err != nil {
		return
	}
	regionItem.Region = region
	regionItem.Confidence = confidence
	if err = l.svcCtx.RegionModel.Update(l.ctx, regionItem); err != nil {
		return
	}

	return
}

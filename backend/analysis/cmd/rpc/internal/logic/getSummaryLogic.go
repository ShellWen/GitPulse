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

type GetSummaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSummaryLogic {
	return &GetSummaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSummaryLogic) GetSummary(in *pb.GetAnalysisReq) (resp *pb.GetSummaryResp, err error) {
	var summary *model.Summary
	if summary, err = l.svcCtx.SummaryModel.FindOneByDeveloperId(l.ctx, in.DeveloperId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetSummaryResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetSummaryResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetSummaryResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Summary: &pb.Summary{
				DataId:        summary.DataId,
				DataCreatedAt: summary.DataCreatedAt.Unix(),
				DataUpdatedAt: summary.DataUpdatedAt.Unix(),
				Summary:       summary.Summary,
			},
		}
	}

	err = nil
	return
}

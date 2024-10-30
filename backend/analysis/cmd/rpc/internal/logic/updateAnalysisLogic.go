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

type UpdateAnalysisLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAnalysisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAnalysisLogic {
	return &UpdateAnalysisLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAnalysisLogic) UpdateAnalysis(in *pb.UpdateAnalysisReq) (resp *pb.UpdateAnalysisResp, err error) {
	var analysis *model.Analysis
	if analysis, err = l.svcCtx.AnalysisModel.FindOneByDeveloperId(l.ctx, in.DeveloperId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.UpdateAnalysisResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.UpdateAnalysisResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else if err = l.doUpdateAnalysis(analysis, in); err != nil {
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

	err = nil
	return
}

func (l *UpdateAnalysisLogic) doUpdateAnalysis(analysis *model.Analysis, in *pb.UpdateAnalysisReq) (err error) {
	analysis.Languages = in.Languages
	analysis.TalentRank = in.TalentRank
	analysis.Nation = in.Nation

	err = l.svcCtx.AnalysisModel.Update(l.ctx, analysis)
	return
}

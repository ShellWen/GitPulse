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

type GetAnalysisLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAnalysisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnalysisLogic {
	return &GetAnalysisLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAnalysisLogic) GetAnalysis(in *pb.GetAnalysisReq) (resp *pb.GetAnalysisResp, err error) {
	var analysis *model.Analysis
	if analysis, err = l.svcCtx.AnalysisModel.FindOneByDeveloperId(l.ctx, in.DeveloperId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetAnalysisResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetAnalysisResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetAnalysisResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Analysis: &pb.Analysis{
				DataId:       analysis.DataId,
				DataCreateAt: analysis.DataCreateAt.Unix(),
				DataUpdateAt: analysis.DataUpdateAt.Unix(),
				DeveloperId:  analysis.DeveloperId,
				Languages:    analysis.Languages,
				TalentRank:   analysis.TalentRank,
				Nation:       analysis.Nation,
			},
		}
	}

	err = nil
	return
}

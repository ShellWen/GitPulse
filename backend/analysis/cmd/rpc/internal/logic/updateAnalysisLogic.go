package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/analysis/model"

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
		return
	}

	analysis.Languages = in.Languages
	analysis.TalentRank = in.TalentRank
	analysis.Nation = in.Nation

	if err = l.svcCtx.AnalysisModel.Update(l.ctx, analysis); err != nil {
		return
	}

	resp = &pb.UpdateAnalysisResp{}

	return
}

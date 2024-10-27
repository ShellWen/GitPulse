package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/analysis/model"

	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAnalysisLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAnalysisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAnalysisLogic {
	return &DelAnalysisLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAnalysisLogic) DelAnalysis(in *pb.DelAnalysisReq) (resp *pb.DelAnalysisResp, err error) {
	var analysis *model.Analysis
	if analysis, err = l.svcCtx.AnalysisModel.FindOneByDeveloperId(l.ctx, in.DeveloperId); err != nil {
		return
	}

	if err = l.svcCtx.AnalysisModel.Delete(l.ctx, analysis.DataId); err != nil {
		return
	}

	resp = &pb.DelAnalysisResp{}

	return
}

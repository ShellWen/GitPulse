package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/pb"
	"github.com/ShellWen/GitPulse/analysis/model"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAnalysisLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAnalysisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAnalysisLogic {
	return &AddAnalysisLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------analysis-----------------------
func (l *AddAnalysisLogic) AddAnalysis(in *pb.AddAnalysisReq) (resp *pb.AddAnalysisResp, err error) {
	analysis := &model.Analysis{
		DataCreateAt: time.Now(),
		DataUpdateAt: time.Now(),
		DeveloperId:  in.DeveloperId,
		Languages:    in.Languages,
		TalentRank:   in.TalentRank,
		Nation:       in.Nation,
	}

	if _, err = l.svcCtx.AnalysisModel.Insert(l.ctx, analysis); err != nil {
		return
	}

	resp = &pb.AddAnalysisResp{}

	return
}

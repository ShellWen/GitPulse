package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelStarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelStarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelStarLogic {
	return &DelStarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelStarLogic) DelStar(in *pb.DelStarReq) (resp *pb.DelStarResp, err error) {
	var star *model.Star
	star, err = l.svcCtx.StarModel.FindOneByDeveloperIdRepoId(l.ctx, in.DeveloperId, in.RepoId)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.StarModel.Delete(l.ctx, star.RepoId)
	if err != nil {
		return nil, err
	}

	resp = &pb.DelStarResp{}

	return
}

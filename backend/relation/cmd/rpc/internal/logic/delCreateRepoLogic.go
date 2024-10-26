package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCreateRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCreateRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCreateRepoLogic {
	return &DelCreateRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCreateRepoLogic) DelCreateRepo(in *pb.DelCreateRepoReq) (resp *pb.DelCreateRepoResp, err error) {
	var createRepo *model.CreateRepo
	createRepo, err = l.svcCtx.CreateRepoModel.FindOneByRepoId(l.ctx, in.RepoId)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.CreateRepoModel.Delete(l.ctx, createRepo.RepoId)
	if err != nil {
		return nil, err
	}

	resp = &pb.DelCreateRepoResp{}

	return
}

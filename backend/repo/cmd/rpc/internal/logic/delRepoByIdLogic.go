package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/repo/model"

	"github.com/ShellWen/GitPulse/repo/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelRepoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelRepoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelRepoByIdLogic {
	return &DelRepoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelRepoByIdLogic) DelRepoById(in *pb.DelRepoByIdReq) (resp *pb.DelRepoByIdResp, err error) {
	var repo *model.Repo

	id := in.Id
	repo, err = l.svcCtx.RepoModel.FindOneById(l.ctx, id)
	if err != nil {
		return nil, err
	}

	dataId := repo.DataId
	err = l.svcCtx.RepoModel.Delete(l.ctx, dataId)
	if err != nil {
		return nil, err
	}

	resp = &pb.DelRepoByIdResp{}

	return
}

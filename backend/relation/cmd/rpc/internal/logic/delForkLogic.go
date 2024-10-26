package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelForkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelForkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelForkLogic {
	return &DelForkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelForkLogic) DelFork(in *pb.DelForkReq) (resp *pb.DelForkResp, err error) {
	var fork *model.Fork
	fork, err = l.svcCtx.ForkModel.FindOneByForkRepoId(l.ctx, in.ForkRepoId)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.ForkModel.Delete(l.ctx, fork.ForkRepoId)
	if err != nil {
		return nil, err
	}

	resp = &pb.DelForkResp{}

	return
}

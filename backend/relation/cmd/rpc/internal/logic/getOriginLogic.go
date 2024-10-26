package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOriginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOriginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOriginLogic {
	return &GetOriginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOriginLogic) GetOrigin(in *pb.GetOriginReq) (resp *pb.GetOriginResp, err error) {
	var fork *model.Fork
	fork, err = l.svcCtx.ForkModel.FindOneByForkRepoId(l.ctx, in.ForkRepoId)
	if err != nil {
		return nil, err
	}

	resp = &pb.GetOriginResp{
		OriginalRepoId: fork.OriginalRepoId,
	}

	return
}

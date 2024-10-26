package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"time"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddForkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddForkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddForkLogic {
	return &AddForkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------fork-----------------------
func (l *AddForkLogic) AddFork(in *pb.AddForkReq) (resp *pb.AddForkResp, err error) {
	fork := &model.Fork{
		DataCreateAt:   time.Now(),
		DataUpdateAt:   time.Now(),
		OriginalRepoId: in.OriginalRepoId,
		ForkRepoId:     in.ForkRepoId,
	}

	_, err = l.svcCtx.ForkModel.Insert(l.ctx, fork)
	if err != nil {
		return nil, err
	}

	resp = &pb.AddForkResp{}

	return
}

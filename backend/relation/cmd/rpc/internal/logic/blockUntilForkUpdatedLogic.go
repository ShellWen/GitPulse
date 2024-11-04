package logic

import (
	"context"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlockUntilForkUpdatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBlockUntilForkUpdatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlockUntilForkUpdatedLogic {
	return &BlockUntilForkUpdatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BlockUntilForkUpdatedLogic) BlockUntilForkUpdated(in *pb.BlockUntilForkUpdatedReq) (*pb.BlockUntilForkUpdatedResp, error) {
	if l.svcCtx.ForkUpdatedChan[in.Id] == nil {
		l.svcCtx.ForkUpdatedChan[in.Id] = make(chan struct{})
	}

	<-l.svcCtx.ForkUpdatedChan[in.Id]

	return &pb.BlockUntilForkUpdatedResp{}, nil
}

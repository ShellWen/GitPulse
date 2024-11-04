package logic

import (
	"context"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlockUntilFollowingUpdatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBlockUntilFollowingUpdatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlockUntilFollowingUpdatedLogic {
	return &BlockUntilFollowingUpdatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BlockUntilFollowingUpdatedLogic) BlockUntilFollowingUpdated(in *pb.BlockUntilFollowingUpdatedReq) (*pb.BlockUntilFollowingUpdatedResp, error) {
	if l.svcCtx.FollowingUpdatedChan[in.Id] == nil {
		l.svcCtx.FollowingUpdatedChan[in.Id] = make(chan struct{})
	}

	<-l.svcCtx.FollowingUpdatedChan[in.Id]

	return &pb.BlockUntilFollowingUpdatedResp{}, nil
}

package logic

import (
	"context"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlockUntilStarredRepoUpdatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBlockUntilStarredRepoUpdatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlockUntilStarredRepoUpdatedLogic {
	return &BlockUntilStarredRepoUpdatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BlockUntilStarredRepoUpdatedLogic) BlockUntilStarredRepoUpdated(in *pb.BlockUntilStarredRepoUpdatedReq) (*pb.BlockUntilStarredRepoUpdatedResp, error) {
	if l.svcCtx.StarredRepoUpdatedChan[in.Id] == nil {
		l.svcCtx.StarredRepoUpdatedChan[in.Id] = make(chan struct{})
	}

	<-l.svcCtx.StarredRepoUpdatedChan[in.Id]

	return &pb.BlockUntilStarredRepoUpdatedResp{}, nil
}

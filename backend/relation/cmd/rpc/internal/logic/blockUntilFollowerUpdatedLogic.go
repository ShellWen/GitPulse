package logic

import (
	"context"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlockUntilFollowerUpdatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBlockUntilFollowerUpdatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlockUntilFollowerUpdatedLogic {
	return &BlockUntilFollowerUpdatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BlockUntilFollowerUpdatedLogic) BlockUntilFollowerUpdated(in *pb.BlockUntilFollowerUpdatedReq) (*pb.BlockUntilFollowerUpdatedResp, error) {
	if l.svcCtx.FollowerUpdatedChan[in.Id] == nil {
		l.svcCtx.FollowerUpdatedChan[in.Id] = make(chan struct{})
	}

	<-l.svcCtx.FollowerUpdatedChan[in.Id]

	return &pb.BlockUntilFollowerUpdatedResp{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}, nil
}

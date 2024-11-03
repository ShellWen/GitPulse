package logic

import (
	"context"
	"net/http"

	"github.com/ShellWen/GitPulse/repo/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlockUntilRepoUpdatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBlockUntilRepoUpdatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlockUntilRepoUpdatedLogic {
	return &BlockUntilRepoUpdatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BlockUntilRepoUpdatedLogic) BlockUntilRepoUpdated(in *pb.BlockUntilRepoUpdatedReq) (resp *pb.BlockUntilRepoUpdatedResp, err error) {
	if l.svcCtx.RepoUpdatedChan[in.Id] == nil {
		l.svcCtx.RepoUpdatedChan[in.Id] = make(chan struct{})
	}

	<-l.svcCtx.RepoUpdatedChan[in.Id]

	return &pb.BlockUntilRepoUpdatedResp{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}, nil
}

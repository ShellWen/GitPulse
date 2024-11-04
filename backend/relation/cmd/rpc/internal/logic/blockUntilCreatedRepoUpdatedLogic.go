package logic

import (
	"context"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlockUntilCreatedRepoUpdatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBlockUntilCreatedRepoUpdatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlockUntilCreatedRepoUpdatedLogic {
	return &BlockUntilCreatedRepoUpdatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BlockUntilCreatedRepoUpdatedLogic) BlockUntilCreatedRepoUpdated(in *pb.BlockUntilCreatedRepoUpdatedReq) (resp *pb.BlockUntilCreatedRepoUpdatedResp, err error) {
	if l.svcCtx.CreatedRepoUpdatedChan[in.Id] == nil {
		l.svcCtx.CreatedRepoUpdatedChan[in.Id] = make(chan struct{})
	}

	<-l.svcCtx.CreatedRepoUpdatedChan[in.Id]

	return &pb.BlockUntilCreatedRepoUpdatedResp{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}, nil
}

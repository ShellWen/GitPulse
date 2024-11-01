package logic

import (
	"context"
	"net/http"

	"github.com/ShellWen/GitPulse/developer/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlockUntilDeveloperUpdatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBlockUntilDeveloperUpdatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlockUntilDeveloperUpdatedLogic {
	return &BlockUntilDeveloperUpdatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BlockUntilDeveloperUpdatedLogic) BlockUntilDeveloperUpdated(in *pb.BlockUntilDeveloperUpdatedReq) (*pb.BlockUntilDeveloperUpdatedResp, error) {
	if l.svcCtx.DeveloperUpdatedChan[in.Id] == nil {
		l.svcCtx.DeveloperUpdatedChan[in.Id] = make(chan struct{})
	}

	<-l.svcCtx.DeveloperUpdatedChan[in.Id]

	return &pb.BlockUntilDeveloperUpdatedResp{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}, nil
}

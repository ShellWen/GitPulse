package logic

import (
	"context"
	"net/http"

	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlockUntilAllUpdatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBlockUntilAllUpdatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlockUntilAllUpdatedLogic {
	return &BlockUntilAllUpdatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BlockUntilAllUpdatedLogic) BlockUntilAllUpdated(in *pb.BlockUntilAllUpdatedReq) (*pb.BlockUntilAllUpdatedResp, error) {
	if l.svcCtx.AllUpdatedChan[in.UserId] == nil {
		l.svcCtx.AllUpdatedChan[in.UserId] = make(chan struct{})
	}

	<-l.svcCtx.AllUpdatedChan[in.UserId]

	return &pb.BlockUntilAllUpdatedResp{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}, nil
}

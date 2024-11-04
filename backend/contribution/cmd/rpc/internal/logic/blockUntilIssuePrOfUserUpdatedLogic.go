package logic

import (
	"context"
	"net/http"

	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlockUntilIssuePrOfUserUpdatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBlockUntilIssuePrOfUserUpdatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlockUntilIssuePrOfUserUpdatedLogic {
	return &BlockUntilIssuePrOfUserUpdatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BlockUntilIssuePrOfUserUpdatedLogic) BlockUntilIssuePrOfUserUpdated(in *pb.BlockUntilIssuePrOfUserUpdatedReq) (*pb.BlockUntilIssuePrOfUserUpdatedResp, error) {
	if l.svcCtx.IssuePrUpdatedChan[in.UserId] == nil {
		l.svcCtx.IssuePrUpdatedChan[in.UserId] = make(chan struct{})
	}

	<-l.svcCtx.IssuePrUpdatedChan[in.UserId]

	return &pb.BlockUntilIssuePrOfUserUpdatedResp{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}, nil
}

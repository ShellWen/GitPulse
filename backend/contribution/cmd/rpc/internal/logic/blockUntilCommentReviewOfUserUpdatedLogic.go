package logic

import (
	"context"
	"net/http"

	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlockUntilCommentReviewOfUserUpdatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBlockUntilCommentReviewOfUserUpdatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlockUntilCommentReviewOfUserUpdatedLogic {
	return &BlockUntilCommentReviewOfUserUpdatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BlockUntilCommentReviewOfUserUpdatedLogic) BlockUntilCommentReviewOfUserUpdated(in *pb.BlockUntilCommentReviewOfUserUpdatedReq) (*pb.BlockUntilCommentReviewOfUserUpdatedResp, error) {
	if l.svcCtx.CommentReviewUpdatedChan[in.UserId] == nil {
		l.svcCtx.CommentReviewUpdatedChan[in.UserId] = make(chan struct{})
	}

	<-l.svcCtx.CommentReviewUpdatedChan[in.UserId]

	return &pb.BlockUntilCommentReviewOfUserUpdatedResp{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}, nil
}

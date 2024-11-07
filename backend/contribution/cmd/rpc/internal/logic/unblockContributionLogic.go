package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/common/tasks"
	"net/http"

	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnblockContributionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnblockContributionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnblockContributionLogic {
	return &UnblockContributionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnblockContributionLogic) UnblockContribution(in *pb.UnblockContributionReq) (*pb.UnblockContributionResp, error) {
	switch in.FetchType {
	case tasks.FetchContributionOfUser:
		l.syncAllContribution(in.Id)
	case tasks.FetchIssuePROfUser:
		l.syncIssuePr(in.Id)
	case tasks.FetchCommentOfUser:
		l.syncCommentReview(in.Id)
	default:
		logx.Error("Invalid message type: ", in.FetchType)
		return &pb.UnblockContributionResp{
			Code:    http.StatusInternalServerError,
			Message: "Invalid message type",
		}, nil
	}

	return &pb.UnblockContributionResp{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}

func (l *UnblockContributionLogic) syncIssuePr(id int64) {
	if l.svcCtx.IssuePrUpdatedChan[id] == nil {
		l.svcCtx.IssuePrUpdatedChan[id] = make(chan struct{})
	}

	for stillHasBlock := true; stillHasBlock; {
		select {
		case l.svcCtx.IssuePrUpdatedChan[id] <- struct{}{}:
			stillHasBlock = true
		default:
			stillHasBlock = false
		}
	}
}

func (l *UnblockContributionLogic) syncCommentReview(id int64) {
	if l.svcCtx.CommentReviewUpdatedChan[id] == nil {
		l.svcCtx.CommentReviewUpdatedChan[id] = make(chan struct{})
	}

	for stillHasBlock := true; stillHasBlock; {
		select {
		case l.svcCtx.CommentReviewUpdatedChan[id] <- struct{}{}:
			stillHasBlock = true
		default:
			stillHasBlock = false
		}
	}

}

func (l *UnblockContributionLogic) syncAllContribution(id int64) {
	if l.svcCtx.AllUpdatedChan[id] == nil {
		l.svcCtx.AllUpdatedChan[id] = make(chan struct{})
	}

	for stillHasBlock := true; stillHasBlock; {
		select {
		case l.svcCtx.AllUpdatedChan[id] <- struct{}{}:
			stillHasBlock = true
		default:
			stillHasBlock = false
		}
	}

	return
}

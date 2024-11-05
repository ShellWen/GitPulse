package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/common/tasks"
	"net/http"

	"github.com/ShellWen/GitPulse/repo/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnblockRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnblockRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnblockRepoLogic {
	return &UnblockRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnblockRepoLogic) UnblockRepo(in *pb.UnblockRepoReq) (*pb.UnblockRepoResp, error) {
	if in.FetchType != tasks.FetchRepo {
		logx.Error("Invalid message type: ", in.FetchType)
		return &pb.UnblockRepoResp{
			Code:    http.StatusInternalServerError,
			Message: "Invalid message type",
		}, nil
	}

	if l.svcCtx.RepoUpdatedChan[in.Id] == nil {
		l.svcCtx.RepoUpdatedChan[in.Id] = make(chan struct{})
	}

	select {
	case l.svcCtx.RepoUpdatedChan[in.Id] <- struct{}{}:
	default:
	}

	return &pb.UnblockRepoResp{
		Code:    http.StatusOK,
		Message: "UnblockRepoed",
	}, nil
}

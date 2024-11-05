package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/common/tasks"
	"net/http"

	"github.com/ShellWen/GitPulse/developer/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnblockDeveloperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnblockDeveloperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnblockDeveloperLogic {
	return &UnblockDeveloperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnblockDeveloperLogic) UnblockDeveloper(in *pb.UnblockDeveloperReq) (*pb.UnblockDeveloperResp, error) {
	if in.FetchType != tasks.FetchDeveloper {
		logx.Error("Invalid message type: ", in.FetchType)
		return &pb.UnblockDeveloperResp{
			Code:    http.StatusInternalServerError,
			Message: "Invalid message type",
		}, nil
	}

	if l.svcCtx.DeveloperUpdatedChan[in.Id] == nil {
		l.svcCtx.DeveloperUpdatedChan[in.Id] = make(chan struct{})
	}

	select {
	case l.svcCtx.DeveloperUpdatedChan[in.Id] <- struct{}{}:
	default:
	}

	return &pb.UnblockDeveloperResp{
		Code:    http.StatusOK,
		Message: "UnblockDevelopered",
	}, nil
}

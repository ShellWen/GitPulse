package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/common/tasks"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnblockRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnblockRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnblockRelationLogic {
	return &UnblockRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnblockRelationLogic) UnblockRelation(in *pb.UnblockRelationReq) (*pb.UnblockRelationResp, error) {
	switch in.FetchType {
	case tasks.FetchCreatedRepo:
		l.syncCreatedRepo(in.Id)
	case tasks.FetchStarredRepo:
		l.syncStarredRepo(in.Id)
	case tasks.FetchFork:
		l.syncFork(in.Id)
	case tasks.FetchFollowing:
		l.syncFollowing(in.Id)
	case tasks.FetchFollower:
		l.syncFollower(in.Id)
	default:
		logx.Error("Invalid message type: ", in.FetchType)
		return &pb.UnblockRelationResp{
			Code:    http.StatusInternalServerError,
			Message: "Invalid message type",
		}, nil
	}

	return &pb.UnblockRelationResp{
		Code:    http.StatusOK,
		Message: "Unblocked",
	}, nil
}

func (l *UnblockRelationLogic) syncCreatedRepo(id int64) {
	if l.svcCtx.CreatedRepoUpdatedChan[id] == nil {
		l.svcCtx.CreatedRepoUpdatedChan[id] = make(chan struct{})
	}

	select {
	case l.svcCtx.CreatedRepoUpdatedChan[id] <- struct{}{}:
	default:
	}
}

func (l *UnblockRelationLogic) syncStarredRepo(id int64) {
	if l.svcCtx.StarredRepoUpdatedChan[id] == nil {
		l.svcCtx.StarredRepoUpdatedChan[id] = make(chan struct{})
	}

	select {
	case l.svcCtx.StarredRepoUpdatedChan[id] <- struct{}{}:
	default:
	}
}

func (l *UnblockRelationLogic) syncFork(id int64) {
	if l.svcCtx.ForkUpdatedChan[id] == nil {
		l.svcCtx.ForkUpdatedChan[id] = make(chan struct{})
	}

	select {
	case l.svcCtx.ForkUpdatedChan[id] <- struct{}{}:
	default:
	}
}

func (l *UnblockRelationLogic) syncFollowing(id int64) {
	if l.svcCtx.FollowingUpdatedChan[id] == nil {
		l.svcCtx.FollowingUpdatedChan[id] = make(chan struct{})
	}

	select {
	case l.svcCtx.FollowingUpdatedChan[id] <- struct{}{}:
	default:
	}
}

func (l *UnblockRelationLogic) syncFollower(id int64) {
	if l.svcCtx.FollowerUpdatedChan[id] == nil {
		l.svcCtx.FollowerUpdatedChan[id] = make(chan struct{})
	}

	select {
	case l.svcCtx.FollowerUpdatedChan[id] <- struct{}{}:
	default:
	}
}

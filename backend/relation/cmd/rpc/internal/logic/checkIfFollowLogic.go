package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/relation/model"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIfFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckIfFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIfFollowLogic {
	return &CheckIfFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckIfFollowLogic) CheckIfFollow(in *pb.CheckIfFollowReq) (resp *pb.CheckFollowResp, err error) {
	_, err = l.svcCtx.FollowModel.FindOneByFollowingIdFollowedId(l.ctx, in.FollowingId, in.FollowedId)

	if err == nil {
		resp = &pb.CheckFollowResp{
			IsFollow: true,
		}
	} else if errors.Is(err, model.ErrNotFound) {
		resp = &pb.CheckFollowResp{
			IsFollow: false,
		}
		err = nil
	}

	return
}

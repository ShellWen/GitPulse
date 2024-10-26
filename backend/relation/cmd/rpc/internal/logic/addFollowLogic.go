package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"time"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFollowLogic {
	return &AddFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------follow-----------------------
func (l *AddFollowLogic) AddFollow(in *pb.AddFollowReq) (resp *pb.AddFollowResp, err error) {
	follow := &model.Follow{
		DataCreateAt: time.Now(),
		DataUpdateAt: time.Now(),
		FollowingId:  in.FollowingId,
		FollowedId:   in.FollowedId,
	}

	_, err = l.svcCtx.FollowModel.Insert(l.ctx, follow)
	if err != nil {
		return nil, err
	}

	resp = &pb.AddFollowResp{}

	return
}

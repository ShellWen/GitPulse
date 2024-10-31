package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

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
		FollowingId: in.FollowingId,
		FollowedId:  in.FollowedId,
	}

	if _, err = l.svcCtx.FollowModel.Insert(l.ctx, follow); err != nil {
		resp = &pb.AddFollowResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.AddFollowResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

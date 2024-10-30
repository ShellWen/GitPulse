package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelFollowLogic {
	return &DelFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelFollowLogic) DelFollow(in *pb.DelFollowReq) (resp *pb.DelFollowResp, err error) {
	var follow *model.Follow
	if follow, err = l.svcCtx.FollowModel.FindOneByFollowingIdFollowedId(l.ctx, in.FollowingId, in.FollowedId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.DelFollowResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.DelFollowResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else if err = l.svcCtx.FollowModel.Delete(l.ctx, follow.DataId); err != nil {
		resp = &pb.DelFollowResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.DelFollowResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

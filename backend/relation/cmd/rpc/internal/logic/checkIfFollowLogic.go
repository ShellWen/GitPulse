package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"
	"github.com/ShellWen/GitPulse/relation/model"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
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
	if _, err = l.svcCtx.FollowModel.FindOneByFollowingIdFollowedId(l.ctx, in.FollowingId, in.FollowedId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.CheckFollowResp{
				Code:     http.StatusOK,
				Message:  http.StatusText(http.StatusOK),
				IsFollow: false,
			}
		default:
			resp = &pb.CheckFollowResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.CheckFollowResp{
			Code:     http.StatusOK,
			Message:  http.StatusText(http.StatusOK),
			IsFollow: true,
		}
	}

	err = nil
	return
}

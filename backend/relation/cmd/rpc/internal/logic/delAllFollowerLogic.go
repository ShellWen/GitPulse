package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllFollowerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllFollowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllFollowerLogic {
	return &DelAllFollowerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllFollowerLogic) DelAllFollower(in *pb.DelAllFollowerReq) (resp *pb.DelAllFollowerResp, err error) {
	var follower *[]*model.Follow
	if follower, err = l.svcCtx.FollowModel.SearchFollowerByDeveloperId(l.ctx, in.DeveloperId, 1, 9223372036854775807); err != nil {
		resp = &pb.DelAllFollowerResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		for _, follow := range *follower {
			if err = l.svcCtx.FollowModel.Delete(l.ctx, follow.DataId); err != nil {
				resp = &pb.DelAllFollowerResp{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
				break
			}
		}
		if err == nil {
			resp = &pb.DelAllFollowerResp{
				Code:    http.StatusOK,
				Message: http.StatusText(http.StatusOK),
			}
		}
	}

	err = nil
	return
}

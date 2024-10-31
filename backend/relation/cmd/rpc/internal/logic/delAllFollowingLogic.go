package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllFollowingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllFollowingLogic {
	return &DelAllFollowingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllFollowingLogic) DelAllFollowing(in *pb.DelAllFollowingReq) (resp *pb.DelAllFollowingResp, err error) {
	var following *[]*model.Follow
	if following, err = l.svcCtx.FollowModel.SearchFollowingByDeveloperId(l.ctx, in.DeveloperId, 1, 9223372036854775807); err != nil {
		resp = &pb.DelAllFollowingResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if len(*following) == 0 {
		resp = &pb.DelAllFollowingResp{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	} else {
		for _, follow := range *following {
			if err = l.svcCtx.FollowModel.Delete(l.ctx, follow.DataId); err != nil {
				resp = &pb.DelAllFollowingResp{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
				break
			}
		}
		if err == nil {
			resp = &pb.DelAllFollowingResp{
				Code:    http.StatusOK,
				Message: http.StatusText(http.StatusOK),
			}
		}
	}

	err = nil
	return
}

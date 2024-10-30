package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllFollowedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllFollowedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllFollowedLogic {
	return &DelAllFollowedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllFollowedLogic) DelAllFollowed(in *pb.DelAllFollowedReq) (resp *pb.DelAllFollowedResp, err error) {
	var followed *[]*model.Follow
	if followed, err = l.svcCtx.FollowModel.SearchFollowed(l.ctx, in.DeveloperId, 1, 9223372036854775807); err != nil {
		resp = &pb.DelAllFollowedResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if len(*followed) == 0 {
		resp = &pb.DelAllFollowedResp{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	} else {
		for _, follow := range *followed {
			if err = l.svcCtx.FollowModel.Delete(l.ctx, follow.DataId); err != nil {
				resp = &pb.DelAllFollowedResp{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
				break
			}
		}
		if err == nil {
			resp = &pb.DelAllFollowedResp{
				Code:    http.StatusOK,
				Message: http.StatusText(http.StatusOK),
			}
		}
	}

	err = nil
	return
}

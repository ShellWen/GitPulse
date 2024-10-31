package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFollowerByDeveloperIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchFollowerByDeveloperIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFollowerByDeveloperIdLogic {
	return &SearchFollowerByDeveloperIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchFollowerByDeveloperIdLogic) SearchFollowerByDeveloperId(in *pb.SearchFollowerByDeveloperIdReq) (resp *pb.SearchFollowerByDeveloperIdResp, err error) {
	var follows *[]*model.Follow
	var followerIds *[]int64

	if follows, err = l.svcCtx.FollowModel.SearchFollowerByDeveloperId(l.ctx, in.DeveloperId, in.Page, in.Limit); err != nil {
		resp = &pb.SearchFollowerByDeveloperIdResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if followerIds = l.buildFollowerIds(follows); len(*followerIds) == 0 {
		resp = &pb.SearchFollowerByDeveloperIdResp{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	} else {
		resp = &pb.SearchFollowerByDeveloperIdResp{
			Code:        http.StatusOK,
			Message:     http.StatusText(http.StatusOK),
			FollowerIds: *followerIds,
		}
	}

	err = nil
	return
}

func (l *SearchFollowerByDeveloperIdLogic) buildFollowerIds(follows *[]*model.Follow) (followerIds *[]int64) {
	followerIds = new([]int64)
	for _, follow := range *follows {
		*followerIds = append(*followerIds, follow.FollowerId)
	}

	return
}

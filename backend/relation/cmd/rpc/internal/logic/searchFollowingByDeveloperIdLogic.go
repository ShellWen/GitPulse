package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFollowingByDeveloperIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchFollowingByDeveloperIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFollowingByDeveloperIdLogic {
	return &SearchFollowingByDeveloperIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchFollowingByDeveloperIdLogic) SearchFollowingByDeveloperId(in *pb.SearchFollowingByDeveloperIdReq) (resp *pb.SearchFollowingByDeveloperIdResp, err error) {
	var follows *[]*model.Follow
	var followingIds *[]int64

	if follows, err = l.svcCtx.FollowModel.SearchFollowingByDeveloperId(l.ctx, in.DeveloperId, in.Page, in.Limit); err != nil {
		resp = &pb.SearchFollowingByDeveloperIdResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if followingIds = l.buildFollowingIds(follows); len(*followingIds) == 0 {
		resp = &pb.SearchFollowingByDeveloperIdResp{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	} else {
		resp = &pb.SearchFollowingByDeveloperIdResp{
			Code:         http.StatusOK,
			Message:      http.StatusText(http.StatusOK),
			FollowingIds: *followingIds,
		}
	}

	err = nil
	return
}

func (l *SearchFollowingByDeveloperIdLogic) buildFollowingIds(follows *[]*model.Follow) (followingIds *[]int64) {
	followingIds = new([]int64)
	for _, follow := range *follows {
		*followingIds = append(*followingIds, follow.FollowingId)
	}

	return
}

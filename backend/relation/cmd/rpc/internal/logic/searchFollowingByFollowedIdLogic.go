package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFollowingByFollowedIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchFollowingByFollowedIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFollowingByFollowedIdLogic {
	return &SearchFollowingByFollowedIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchFollowingByFollowedIdLogic) SearchFollowingByFollowedId(in *pb.SearchFollowingByFollowedIdReq) (resp *pb.SearchFollowByFollowedIdResp, err error) {
	var follows *[]*model.Follow
	var followingIds *[]int64

	if follows, err = l.svcCtx.FollowModel.SearchFollowing(l.ctx, in.FollowedId, in.Page, in.Limit); err != nil {
		resp = &pb.SearchFollowByFollowedIdResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if followingIds = l.buildFollowingIds(follows); len(*followingIds) == 0 {
		resp = &pb.SearchFollowByFollowedIdResp{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	} else {
		resp = &pb.SearchFollowByFollowedIdResp{
			Code:         http.StatusOK,
			Message:      http.StatusText(http.StatusOK),
			FollowingIds: *followingIds,
		}
	}

	err = nil
	return
}

func (l *SearchFollowingByFollowedIdLogic) buildFollowingIds(follows *[]*model.Follow) (followingIds *[]int64) {
	followingIds = new([]int64)
	for _, follow := range *follows {
		*followingIds = append(*followingIds, follow.FollowingId)
	}

	return
}

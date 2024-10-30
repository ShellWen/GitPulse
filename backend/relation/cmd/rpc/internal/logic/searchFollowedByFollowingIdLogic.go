package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFollowedByFollowingIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchFollowedByFollowingIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFollowedByFollowingIdLogic {
	return &SearchFollowedByFollowingIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchFollowedByFollowingIdLogic) SearchFollowedByFollowingId(in *pb.SearchFollowedByFollowingIdReq) (resp *pb.SearchFollowByFollowingIdResp, err error) {
	var follows *[]*model.Follow
	var followedIds *[]int64

	if follows, err = l.svcCtx.FollowModel.SearchFollowed(l.ctx, in.FollowingId, in.Page, in.Limit); err != nil {
		resp = &pb.SearchFollowByFollowingIdResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if followedIds = l.buildFollowedIds(follows); len(*followedIds) == 0 {
		resp = &pb.SearchFollowByFollowingIdResp{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	} else {
		resp = &pb.SearchFollowByFollowingIdResp{
			Code:        http.StatusOK,
			Message:     http.StatusText(http.StatusOK),
			FollowedIds: *followedIds,
		}
	}

	err = nil
	return
}

func (l *SearchFollowedByFollowingIdLogic) buildFollowedIds(follows *[]*model.Follow) (followedIds *[]int64) {
	followedIds = new([]int64)
	for _, follow := range *follows {
		*followedIds = append(*followedIds, follow.FollowedId)
	}

	return
}

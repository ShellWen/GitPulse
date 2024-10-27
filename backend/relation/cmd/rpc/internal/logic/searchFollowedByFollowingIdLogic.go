package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"

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
	if follows, err = l.svcCtx.FollowModel.SearchFollowed(l.ctx, in.FollowingId, in.Page, in.Limit); err != nil {
		return nil, err
	}

	var followedIds []int64
	for _, follow := range *follows {
		followedIds = append(followedIds, follow.FollowedId)
	}

	resp = &pb.SearchFollowByFollowingIdResp{
		FollowedIds: followedIds,
	}

	return
}

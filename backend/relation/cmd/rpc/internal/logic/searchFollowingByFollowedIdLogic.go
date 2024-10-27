package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"

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

	if follows, err = l.svcCtx.FollowModel.SearchFollowing(l.ctx, in.FollowedId, in.Page, in.Limit); err != nil {
		return nil, err
	}

	var followingIds []int64
	for _, follow := range *follows {
		followingIds = append(followingIds, follow.FollowingId)
	}

	resp = &pb.SearchFollowByFollowedIdResp{
		FollowingIds: followingIds,
	}

	return
}

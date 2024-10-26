package logic

import (
	"context"

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

func (l *SearchFollowedByFollowingIdLogic) SearchFollowedByFollowingId(in *pb.SearchFollowedByFollowingIdReq) (*pb.SearchFollowByFollowingIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchFollowByFollowingIdResp{}, nil
}

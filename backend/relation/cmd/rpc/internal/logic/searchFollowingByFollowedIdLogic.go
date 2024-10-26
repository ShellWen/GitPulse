package logic

import (
	"context"

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

func (l *SearchFollowingByFollowedIdLogic) SearchFollowingByFollowedId(in *pb.SearchFollowingByFollowedIdReq) (*pb.SearchFollowByFollowedIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchFollowByFollowedIdResp{}, nil
}

package logic

import (
	"context"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchForkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchForkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForkLogic {
	return &SearchForkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchForkLogic) SearchFork(in *pb.SearchForkReq) (*pb.SearchForkResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchForkResp{}, nil
}

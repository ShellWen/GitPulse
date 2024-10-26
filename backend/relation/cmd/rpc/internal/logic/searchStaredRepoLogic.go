package logic

import (
	"context"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchStaredRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchStaredRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchStaredRepoLogic {
	return &SearchStaredRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchStaredRepoLogic) SearchStaredRepo(in *pb.SearchStaredRepoReq) (*pb.SearchStaredRepoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchStaredRepoResp{}, nil
}

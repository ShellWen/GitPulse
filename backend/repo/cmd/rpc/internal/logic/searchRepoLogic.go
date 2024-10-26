package logic

import (
	"context"

	"github.com/ShellWen/GitPulse/repo/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchRepoLogic {
	return &SearchRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchRepoLogic) SearchRepo(in *pb.SearchRepoReq) (*pb.SearchRepoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchRepoResp{}, nil
}

package logic

import (
	"context"

	"github.com/ShellWen/GitPulse/developer/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchDeveloperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchDeveloperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchDeveloperLogic {
	return &SearchDeveloperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchDeveloperLogic) SearchDeveloper(in *pb.SearchDeveloperReq) (*pb.SearchDeveloperResp, error) {
	// TODO: a high-performance search logic

	return &pb.SearchDeveloperResp{}, nil
}

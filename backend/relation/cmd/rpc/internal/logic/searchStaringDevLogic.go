package logic

import (
	"context"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchStaringDevLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchStaringDevLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchStaringDevLogic {
	return &SearchStaringDevLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchStaringDevLogic) SearchStaringDev(in *pb.SearchStaringDevReq) (*pb.SearchStaringDevResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchStaringDevResp{}, nil
}

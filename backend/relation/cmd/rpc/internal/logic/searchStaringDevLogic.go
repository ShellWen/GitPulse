package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"

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

func (l *SearchStaringDevLogic) SearchStaringDev(in *pb.SearchStaringDevReq) (resp *pb.SearchStaringDevResp, err error) {
	var stars *[]*model.Star

	if stars, err = l.svcCtx.StarModel.SearchStaringDeveloper(l.ctx, in.RepoId, in.Page, in.Limit); err != nil {
		return nil, err
	}

	var staringDevIds []uint64
	for _, star := range *stars {
		staringDevIds = append(staringDevIds, star.DeveloperId)
	}

	resp = &pb.SearchStaringDevResp{
		DeveloperIds: staringDevIds,
	}

	return
}

package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"

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

func (l *SearchStaredRepoLogic) SearchStaredRepo(in *pb.SearchStaredRepoReq) (resp *pb.SearchStaredRepoResp, err error) {
	var stars *[]*model.Star

	if stars, err = l.svcCtx.StarModel.SearchStaredRepo(l.ctx, in.DeveloperId, in.Page, in.Limit); err != nil {
		return nil, err
	}

	var staredRepoIds []int64
	for _, star := range *stars {
		staredRepoIds = append(staredRepoIds, star.RepoId)
	}

	resp = &pb.SearchStaredRepoResp{
		RepoIds: staredRepoIds,
	}

	return
}

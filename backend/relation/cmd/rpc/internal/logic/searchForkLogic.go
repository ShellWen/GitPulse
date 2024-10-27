package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"

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

func (l *SearchForkLogic) SearchFork(in *pb.SearchForkReq) (resp *pb.SearchForkResp, err error) {
	var forks *[]*model.Fork

	if forks, err = l.svcCtx.ForkModel.SearchFork(l.ctx, in.OriginalRepoId, in.Page, in.Limit); err != nil {
		return nil, err
	}

	var forkRepoIds []int64
	for _, fork := range *forks {
		forkRepoIds = append(forkRepoIds, fork.ForkRepoId)
	}

	resp = &pb.SearchForkResp{
		ForkRepoIds: forkRepoIds,
	}

	return
}

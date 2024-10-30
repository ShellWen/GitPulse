package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

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
	var forkRepoIds *[]int64

	if forks, err = l.svcCtx.ForkModel.SearchFork(l.ctx, in.OriginalRepoId, in.Page, in.Limit); err != nil {
		resp = &pb.SearchForkResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if forkRepoIds = l.buildForkRepoIds(forks); len(*forkRepoIds) == 0 {
		resp = &pb.SearchForkResp{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	} else {
		resp = &pb.SearchForkResp{
			Code:        http.StatusOK,
			Message:     http.StatusText(http.StatusOK),
			ForkRepoIds: *forkRepoIds,
		}
	}

	err = nil
	return
}

func (l *SearchForkLogic) buildForkRepoIds(forks *[]*model.Fork) (forkRepoIds *[]int64) {
	forkRepoIds = new([]int64)
	for _, fork := range *forks {
		*forkRepoIds = append(*forkRepoIds, fork.ForkRepoId)
	}

	return
}

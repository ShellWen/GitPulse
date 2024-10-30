package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

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
	var staredRepoIds *[]int64

	if stars, err = l.svcCtx.StarModel.SearchStaredRepo(l.ctx, in.DeveloperId, in.Page, in.Limit); err != nil {
		resp = &pb.SearchStaredRepoResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if staredRepoIds = l.buildStaredRepoIds(stars); len(*staredRepoIds) == 0 {
		resp = &pb.SearchStaredRepoResp{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	} else {
		resp = &pb.SearchStaredRepoResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			RepoIds: *staredRepoIds,
		}
	}

	err = nil
	return
}

func (l *SearchStaredRepoLogic) buildStaredRepoIds(stars *[]*model.Star) (staredRepoIds *[]int64) {
	staredRepoIds = new([]int64)
	for _, star := range *stars {
		*staredRepoIds = append(*staredRepoIds, star.RepoId)
	}

	return
}

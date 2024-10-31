package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchStarredRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchStarredRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchStarredRepoLogic {
	return &SearchStarredRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchStarredRepoLogic) SearchStarredRepo(in *pb.SearchStarredRepoReq) (resp *pb.SearchStarredRepoResp, err error) {
	var stars *[]*model.Star
	var starredRepoIds *[]int64

	if stars, err = l.svcCtx.StarModel.SearchStarredRepo(l.ctx, in.DeveloperId, in.Page, in.Limit); err != nil {
		resp = &pb.SearchStarredRepoResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if starredRepoIds = l.buildStarredRepoIds(stars); len(*starredRepoIds) == 0 {
		resp = &pb.SearchStarredRepoResp{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	} else {
		resp = &pb.SearchStarredRepoResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			RepoIds: *starredRepoIds,
		}
	}

	err = nil
	return
}

func (l *SearchStarredRepoLogic) buildStarredRepoIds(stars *[]*model.Star) (starredRepoIds *[]int64) {
	starredRepoIds = new([]int64)
	for _, star := range *stars {
		*starredRepoIds = append(*starredRepoIds, star.RepoId)
	}

	return
}

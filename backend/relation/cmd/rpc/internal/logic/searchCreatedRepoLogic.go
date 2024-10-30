package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCreatedRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchCreatedRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCreatedRepoLogic {
	return &SearchCreatedRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchCreatedRepoLogic) SearchCreatedRepo(in *pb.SearchCreatedRepoReq) (resp *pb.SearchCreatedRepoResp, err error) {
	var CreateRepos *[]*model.CreateRepo
	var repoIds *[]int64

	if CreateRepos, err = l.svcCtx.CreateRepoModel.SearchCreatedRepo(l.ctx, in.DeveloperId, in.Page, in.Limit); err != nil {
		resp = &pb.SearchCreatedRepoResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if repoIds = l.buildRepoIds(CreateRepos); len(*repoIds) == 0 {
		resp = &pb.SearchCreatedRepoResp{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	} else {
		resp = &pb.SearchCreatedRepoResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			RepoIds: *repoIds,
		}
	}

	err = nil
	return
}

func (l *SearchCreatedRepoLogic) buildRepoIds(CreateRepos *[]*model.CreateRepo) (repoIds *[]int64) {
	repoIds = new([]int64)
	for _, CreateRepo := range *CreateRepos {
		*repoIds = append(*repoIds, CreateRepo.RepoId)
	}

	return
}

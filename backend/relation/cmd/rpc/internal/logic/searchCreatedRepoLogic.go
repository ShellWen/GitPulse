package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"

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
	CreateRepos, err = l.svcCtx.CreateRepoModel.SearchCreatedRepo(l.ctx, in.DeveloperId, in.Page, in.Limit)
	if err != nil {
		return nil, err
	}

	var repoIds []int64
	for _, CreateRepo := range *CreateRepos {
		repoIds = append(repoIds, CreateRepo.RepoId)
	}

	resp = &pb.SearchCreatedRepoResp{
		RepoIds: repoIds,
	}

	return
}

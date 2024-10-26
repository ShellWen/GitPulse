package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/pb"
	"github.com/ShellWen/GitPulse/repo/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRepoLogic {
	return &UpdateRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRepoLogic) UpdateRepo(in *pb.UpdateRepoReq) (resp *pb.UpdateRepoResp, err error) {
	var repo *model.Repo
	repo, err = l.svcCtx.RepoModel.FindOneById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	repo.Name = in.Name
	repo.Gist = in.Gist
	repo.StarCount = in.StarCount
	repo.ForkCount = in.ForkCount
	repo.IssueCount = in.IssueCount
	repo.CommitCount = in.CommitCount
	repo.PrCount = in.PrCount
	repo.Language = in.Language
	repo.Description = in.Description
	repo.Readme = in.Readme

	err = l.svcCtx.RepoModel.Update(l.ctx, repo)
	if err != nil {
		return nil, err
	}

	return
}

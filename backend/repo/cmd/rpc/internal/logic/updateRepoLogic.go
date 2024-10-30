package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/pb"
	"github.com/ShellWen/GitPulse/repo/model"
	"net/http"

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
	if repo, err = l.svcCtx.RepoModel.FindOneById(l.ctx, in.Id); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.UpdateRepoResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.UpdateRepoResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else if err = l.doUpdateRepo(repo, in); err != nil {
		resp = &pb.UpdateRepoResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.UpdateRepoResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

func (l *UpdateRepoLogic) doUpdateRepo(repo *model.Repo, in *pb.UpdateRepoReq) (err error) {
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
	return
}

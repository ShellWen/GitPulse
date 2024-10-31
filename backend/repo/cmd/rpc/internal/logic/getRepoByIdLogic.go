package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/repo/model"
	"net/http"

	"github.com/ShellWen/GitPulse/repo/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRepoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRepoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRepoByIdLogic {
	return &GetRepoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRepoByIdLogic) GetRepoById(in *pb.GetRepoByIdReq) (resp *pb.GetRepoByIdResp, err error) {
	var repo *model.Repo
	if repo, err = l.svcCtx.RepoModel.FindOneById(l.ctx, in.Id); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetRepoByIdResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetRepoByIdResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetRepoByIdResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Repo: &pb.Repo{
				DataId:                  repo.DataId,
				DataCreateAt:            repo.DataCreateAt.Unix(),
				DataUpdateAt:            repo.DataUpdateAt.Unix(),
				Id:                      repo.Id,
				Name:                    repo.Name,
				StarCount:               repo.StarCount,
				ForkCount:               repo.ForkCount,
				IssueCount:              repo.IssueCount,
				CommitCount:             repo.CommitCount,
				PrCount:                 repo.PrCount,
				Language:                repo.Language,
				Description:             repo.Description,
				LastFetchForkAt:         repo.LastFetchForkAt.Unix(),
				LastFetchContributionAt: repo.LastFetchContributionAt.Unix(),
			},
		}
	}

	err = nil
	return
}

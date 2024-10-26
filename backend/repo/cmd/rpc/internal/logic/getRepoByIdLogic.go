package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/repo/model"

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
	repo, err = l.svcCtx.RepoModel.FindOneById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	resp = &pb.GetRepoByIdResp{
		Repo: &pb.Repo{
			DataId:       repo.DataId,
			DataCreateAt: repo.DataCreateAt.Unix(),
			DataUpdateAt: repo.DataUpdateAt.Unix(),
			Id:           repo.Id,
			Name:         repo.Name,
			Gist:         repo.Gist,
			StarCount:    repo.StarCount,
			ForkCount:    repo.ForkCount,
			IssueCount:   repo.IssueCount,
			CommitCount:  repo.CommitCount,
			PrCount:      repo.PrCount,
			Language:     repo.Language,
			Description:  repo.Description,
			Readme:       repo.Readme,
		},
	}

	return
}

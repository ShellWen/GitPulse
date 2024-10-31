package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/repo/model"
	"net/http"
	"time"

	"github.com/ShellWen/GitPulse/repo/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRepoLogic {
	return &AddRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------repo-----------------------
func (l *AddRepoLogic) AddRepo(in *pb.AddRepoReq) (resp *pb.AddRepoResp, err error) {
	repo := &model.Repo{
		DataCreatedAt:           time.Now(),
		DataUpdatedAt:           time.Now(),
		Id:                      in.Id,
		Name:                    in.Name,
		StarCount:               in.StarCount,
		ForkCount:               in.ForkCount,
		IssueCount:              in.IssueCount,
		CommitCount:             in.CommitCount,
		PrCount:                 in.PrCount,
		Language:                in.Language,
		Description:             in.Description,
		LastFetchForkAt:         time.Unix(in.LastFetchForkAt, 0),
		LastFetchContributionAt: time.Unix(in.LastFetchContributionAt, 0),
	}

	if _, err = l.svcCtx.RepoModel.Insert(l.ctx, repo); err != nil {
		resp = &pb.AddRepoResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.AddRepoResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

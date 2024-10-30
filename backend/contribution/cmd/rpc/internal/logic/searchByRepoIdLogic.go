package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/contribution/model"
	"net/http"

	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchByRepoIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchByRepoIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchByRepoIdLogic {
	return &SearchByRepoIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchByRepoIdLogic) SearchByRepoId(in *pb.SearchByRepoIdReq) (resp *pb.SearchByRepoIdResp, err error) {
	var contributions *[]*model.Contribution
	var pbContributions *[]*pb.Contribution

	if contributions, err = l.svcCtx.ContributionModel.SearchByRepoId(l.ctx, in.RepoId, in.Page, in.Limit); err != nil {
		resp = &pb.SearchByRepoIdResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if pbContributions = l.buildPBContributions(contributions); len(*pbContributions) == 0 {
		resp = &pb.SearchByRepoIdResp{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	} else {
		resp = &pb.SearchByRepoIdResp{
			Code:          http.StatusOK,
			Message:       http.StatusText(http.StatusOK),
			Contributions: *pbContributions,
		}
	}

	err = nil
	return
}

func (l *SearchByRepoIdLogic) buildPBContributions(contributions *[]*model.Contribution) *[]*pb.Contribution {
	var pbContributions []*pb.Contribution
	for _, contribution := range *contributions {
		pbContributions = append(pbContributions, &pb.Contribution{
			DataId:         contribution.DataId,
			DataCreateAt:   contribution.DataCreateAt.Unix(),
			DataUpdateAt:   contribution.DataUpdateAt.Unix(),
			UserId:         contribution.UserId,
			RepoId:         contribution.RepoId,
			Category:       contribution.Category,
			Content:        contribution.Content,
			CreateAt:       contribution.CreateAt.Unix(),
			UpdateAt:       contribution.UpdateAt.Unix(),
			ContributionId: contribution.ContributionId,
		})
	}
	return &pbContributions
}

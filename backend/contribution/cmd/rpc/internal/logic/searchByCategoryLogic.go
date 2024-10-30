package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/contribution/model"
	"net/http"

	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchByCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchByCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchByCategoryLogic {
	return &SearchByCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchByCategoryLogic) SearchByCategory(in *pb.SearchByCategoryReq) (resp *pb.SearchByCategoryResp, err error) {
	var contributions *[]*model.Contribution
	var pbContributions *[]*pb.Contribution

	if contributions, err = l.svcCtx.ContributionModel.SearchByCategory(l.ctx, in.Category, in.Page, in.Limit); err != nil {
		resp = &pb.SearchByCategoryResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if pbContributions = l.buildPBContributions(contributions); len(*pbContributions) == 0 {
		resp = &pb.SearchByCategoryResp{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	} else {
		resp = &pb.SearchByCategoryResp{
			Code:          http.StatusOK,
			Message:       http.StatusText(http.StatusOK),
			Contributions: *pbContributions,
		}
	}

	err = nil
	return
}

func (l *SearchByCategoryLogic) buildPBContributions(contributions *[]*model.Contribution) *[]*pb.Contribution {
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

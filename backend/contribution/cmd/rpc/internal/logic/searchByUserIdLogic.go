package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/contribution/model"
	"net/http"

	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchByUserIdLogic {
	return &SearchByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchByUserIdLogic) SearchByUserId(in *pb.SearchByUserIdReq) (resp *pb.SearchByUserIdResp, err error) {
	var contributions *[]*model.Contribution
	var pbContributions *[]*pb.Contribution

	if contributions, err = l.svcCtx.ContributionModel.SearchByUserId(l.ctx, in.UserId, in.Page, in.Limit); err != nil {
		resp = &pb.SearchByUserIdResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if pbContributions = l.buildPBContributions(contributions); len(*pbContributions) == 0 {
		resp = &pb.SearchByUserIdResp{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	} else {
		resp = &pb.SearchByUserIdResp{
			Code:          http.StatusOK,
			Message:       http.StatusText(http.StatusOK),
			Contributions: *pbContributions,
		}
	}

	err = nil
	return
}

func (l *SearchByUserIdLogic) buildPBContributions(contributions *[]*model.Contribution) *[]*pb.Contribution {
	var pbContributions []*pb.Contribution
	for _, contribution := range *contributions {
		pbContributions = append(pbContributions, &pb.Contribution{
			DataId:         contribution.DataId,
			DataCreatedAt:  contribution.DataCreatedAt.Unix(),
			DataUpdatedAt:  contribution.DataUpdatedAt.Unix(),
			UserId:         contribution.UserId,
			RepoId:         contribution.RepoId,
			Category:       contribution.Category,
			Content:        contribution.Content,
			CreatedAt:      contribution.CreatedAt.Unix(),
			UpdatedAt:      contribution.UpdatedAt.Unix(),
			ContributionId: contribution.ContributionId,
		})
	}
	return &pbContributions
}

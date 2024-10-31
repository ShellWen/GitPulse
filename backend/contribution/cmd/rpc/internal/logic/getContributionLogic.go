package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/contribution/model"
	"net/http"

	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContributionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetContributionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContributionLogic {
	return &GetContributionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetContributionLogic) GetContribution(in *pb.GetContributionReq) (resp *pb.GetContributionResp, err error) {
	var contribution *model.Contribution

	if contribution, err = l.svcCtx.ContributionModel.FindOneByCategoryRepoIdContributionId(l.ctx, in.Category, in.RepoId, in.ContributionId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetContributionResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetContributionResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetContributionResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Contribution: &pb.Contribution{
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
			},
		}
	}

	err = nil
	return
}

package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/contribution/model"
	"net/http"
	"time"

	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateContributionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateContributionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateContributionLogic {
	return &UpdateContributionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateContributionLogic) UpdateContribution(in *pb.UpdateContributionReq) (resp *pb.UpdateContributionResp, err error) {
	var contribution *model.Contribution
	if contribution, err = l.svcCtx.ContributionModel.FindOneByCategoryRepoIdUserIdContributionId(l.ctx, in.Category, in.RepoId, in.UserId, in.ContributionId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.UpdateContributionResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.UpdateContributionResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else if err = l.doUpdateContribution(contribution, in); err != nil {
		resp = &pb.UpdateContributionResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.UpdateContributionResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

func (l *UpdateContributionLogic) doUpdateContribution(contribution *model.Contribution, in *pb.UpdateContributionReq) (err error) {
	contribution.Content = in.Content
	contribution.UpdateAt = time.Unix(in.UpdateAt, 0)
	contribution.CreateAt = time.Unix(in.CreateAt, 0)

	err = l.svcCtx.ContributionModel.Update(l.ctx, contribution)
	return
}

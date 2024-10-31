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

type DelContributionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelContributionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelContributionLogic {
	return &DelContributionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelContributionLogic) DelContribution(in *pb.DelContributionReq) (resp *pb.DelContributionResp, err error) {
	var contribution *model.Contribution

	if contribution, err = l.svcCtx.ContributionModel.FindOneByCategoryRepoIdContributionId(l.ctx, in.Category, in.RepoId, in.ContributionId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.DelContributionResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.DelContributionResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else if err = l.svcCtx.ContributionModel.Delete(l.ctx, contribution.DataId); err != nil {
		resp = &pb.DelContributionResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.DelContributionResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

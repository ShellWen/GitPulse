package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/contribution/model"
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
		return nil, err
	}

	contribution.Content = in.Content
	contribution.UpdateAt = time.Unix(in.UpdateAt, 0)
	contribution.CreateAt = time.Unix(in.CreateAt, 0)

	if err = l.svcCtx.ContributionModel.Update(l.ctx, contribution); err != nil {
		return nil, err
	}

	resp = &pb.UpdateContributionResp{}

	return
}

package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/contribution/model"

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

	if contribution, err = l.svcCtx.ContributionModel.FindOneByCategoryRepoIdUserIdContributionId(l.ctx, in.Category, in.RepoId, in.UserId, in.ContributionId); err != nil {
		return nil, err
	}

	if err = l.svcCtx.ContributionModel.Delete(l.ctx, contribution.DataId); err != nil {
		return nil, err
	}

	resp = &pb.DelContributionResp{}

	return
}

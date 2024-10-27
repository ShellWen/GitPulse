package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/contribution/model"

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

	if contribution, err = l.svcCtx.ContributionModel.FindOneByCategoryRepoIdUserIdContributionId(l.ctx, in.Category, in.RepoId, in.UserId, in.ContributionId); err != nil {
		return nil, err
	}

	resp = &pb.GetContributionResp{
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

	return
}

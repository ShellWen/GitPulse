package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/contribution/model"

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
	if contributions, err = l.svcCtx.ContributionModel.SearchByCategory(l.ctx, in.Category, in.Page, in.Limit); err != nil {
		return nil, err
	}

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

	resp = &pb.SearchByCategoryResp{
		Contributions: pbContributions,
	}

	return
}

package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/contribution/model"

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
	if contributions, err = l.svcCtx.ContributionModel.SearchByUserId(l.ctx, in.UserId, in.Page, in.Limit); err != nil {
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

	resp = &pb.SearchByUserIdResp{
		Contributions: pbContributions,
	}

	return
}

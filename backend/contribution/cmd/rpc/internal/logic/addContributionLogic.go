package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/contribution/model"
	"time"

	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddContributionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddContributionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddContributionLogic {
	return &AddContributionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------contribution-----------------------
func (l *AddContributionLogic) AddContribution(in *pb.AddContributionReq) (resp *pb.AddContributionResp, err error) {
	contribution := &model.Contribution{
		DataCreateAt:   time.Now(),
		DataUpdateAt:   time.Now(),
		UserId:         in.UserId,
		RepoId:         in.RepoId,
		Category:       in.Category,
		Content:        in.Content,
		CreateAt:       time.Unix(in.CreateAt, 0),
		UpdateAt:       time.Unix(in.UpdateAt, 0),
		ContributionId: in.ContributionId,
	}

	if _, err := l.svcCtx.ContributionModel.Insert(l.ctx, contribution); err != nil {
		return nil, err
	}

	resp = &pb.AddContributionResp{}

	return
}

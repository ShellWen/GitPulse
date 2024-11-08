package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/contribution/model"
	"net/http"
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
		DataCreatedAt:  time.Now(),
		DataUpdatedAt:  time.Now(),
		UserId:         in.UserId,
		RepoId:         in.RepoId,
		Category:       in.Category,
		Content:        in.Content,
		CreatedAt:      time.Unix(in.CreatedAt, 0),
		UpdatedAt:      time.Unix(in.UpdatedAt, 0),
		ContributionId: in.ContributionId,
	}

	if _, err := l.svcCtx.ContributionModel.Insert(l.ctx, contribution); err != nil {
		resp = &pb.AddContributionResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.AddContributionResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

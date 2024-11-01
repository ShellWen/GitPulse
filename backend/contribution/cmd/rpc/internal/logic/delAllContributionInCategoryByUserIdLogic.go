package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/contribution/model"
	"net/http"

	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllContributionInCategoryByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllContributionInCategoryByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllContributionInCategoryByUserIdLogic {
	return &DelAllContributionInCategoryByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllContributionInCategoryByUserIdLogic) DelAllContributionInCategoryByUserId(in *pb.DelAllContributionInCategoryByUserIdReq) (resp *pb.DelAllContributionInCategoryByUserIdResp, err error) {
	var contributions *[]*model.Contribution
	if contributions, err = l.svcCtx.ContributionModel.SearchByUserId(l.ctx, in.UserId, 1, 9223372036854775807); err != nil {
		resp = &pb.DelAllContributionInCategoryByUserIdResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		for _, contribution := range *contributions {
			if contribution.Category == in.Category {
				if err = l.svcCtx.ContributionModel.Delete(l.ctx, contribution.DataId); err != nil {
					resp = &pb.DelAllContributionInCategoryByUserIdResp{
						Code:    http.StatusInternalServerError,
						Message: err.Error(),
					}
					break
				}
			}
		}
		if err == nil {
			resp = &pb.DelAllContributionInCategoryByUserIdResp{
				Code:    http.StatusOK,
				Message: http.StatusText(http.StatusOK),
			}
		}
	}

	err = nil
	return
}

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

type GetIssuePROfUserUpdatedAtLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIssuePROfUserUpdatedAtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIssuePROfUserUpdatedAtLogic {
	return &GetIssuePROfUserUpdatedAtLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetIssuePROfUserUpdatedAtLogic) GetIssuePROfUserUpdatedAt(in *pb.GetIssuePROfUserUpdatedAtReq) (resp *pb.GetIssuePROfUserUpdatedAtResp, err error) {
	var createRepoUpdatedAt *model.IssuePrOfUserUpdatedAt
	if createRepoUpdatedAt, err = l.svcCtx.IssuePrOfUserUpdatedAtModel.FindOneByDeveloperId(l.ctx, in.UserId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetIssuePROfUserUpdatedAtResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetIssuePROfUserUpdatedAtResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetIssuePROfUserUpdatedAtResp{
			Code:      http.StatusOK,
			Message:   http.StatusText(http.StatusOK),
			UpdatedAt: createRepoUpdatedAt.UpdatedAt.Unix(),
		}
	}

	err = nil
	return
}

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

type GetReviewOfUserUpdatedAtLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetReviewOfUserUpdatedAtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReviewOfUserUpdatedAtLogic {
	return &GetReviewOfUserUpdatedAtLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetReviewOfUserUpdatedAtLogic) GetReviewOfUserUpdatedAt(in *pb.GetReviewOfUserUpdatedAtReq) (resp *pb.GetReviewOfUserUpdatedAtResp, err error) {
	var createRepoUpdatedAt *model.ReviewOfUserUpdatedAt
	if createRepoUpdatedAt, err = l.svcCtx.ReviewOfUserUpdatedAtModel.FindOneByDeveloperId(l.ctx, in.UserId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetReviewOfUserUpdatedAtResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetReviewOfUserUpdatedAtResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetReviewOfUserUpdatedAtResp{
			Code:      http.StatusOK,
			Message:   http.StatusText(http.StatusOK),
			UpdatedAt: createRepoUpdatedAt.UpdatedAt.Unix(),
		}
	}

	err = nil
	return
}

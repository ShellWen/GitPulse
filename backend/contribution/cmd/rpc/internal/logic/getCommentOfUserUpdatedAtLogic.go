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

type GetCommentOfUserUpdatedAtLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentOfUserUpdatedAtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentOfUserUpdatedAtLogic {
	return &GetCommentOfUserUpdatedAtLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentOfUserUpdatedAtLogic) GetCommentOfUserUpdatedAt(in *pb.GetCommentOfUserUpdatedAtReq) (resp *pb.GetCommentOfUserUpdatedAtResp, err error) {
	var createRepoUpdatedAt *model.CommentOfUserUpdatedAt
	if createRepoUpdatedAt, err = l.svcCtx.CommentOfUserUpdatedAtModel.FindOneByDeveloperId(l.ctx, in.UserId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetCommentOfUserUpdatedAtResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetCommentOfUserUpdatedAtResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetCommentOfUserUpdatedAtResp{
			Code:      http.StatusOK,
			Message:   http.StatusText(http.StatusOK),
			UpdatedAt: createRepoUpdatedAt.UpdatedAt.Unix(),
		}
	}

	err = nil
	return
}

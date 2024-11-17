package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCreatedRepoUpdatedAtLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCreatedRepoUpdatedAtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCreatedRepoUpdatedAtLogic {
	return &GetCreatedRepoUpdatedAtLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCreatedRepoUpdatedAtLogic) GetCreatedRepoUpdatedAt(in *pb.GetCreatedRepoUpdatedAtReq) (resp *pb.GetCreatedRepoUpdatedAtResp, err error) {
	var createRepoUpdatedAt *model.CreatedRepoUpdatedAt
	if createRepoUpdatedAt, err = l.svcCtx.CreatedRepoUpdatedAtModel.FindOneByDeveloperId(l.ctx, in.DeveloperId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetCreatedRepoUpdatedAtResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetCreatedRepoUpdatedAtResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetCreatedRepoUpdatedAtResp{
			Code:      http.StatusOK,
			Message:   http.StatusText(http.StatusOK),
			UpdatedAt: createRepoUpdatedAt.UpdatedAt.Unix(),
		}
	}

	err = nil
	return
}

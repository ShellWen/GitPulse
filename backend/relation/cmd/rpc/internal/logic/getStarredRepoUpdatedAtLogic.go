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

type GetStarredRepoUpdatedAtLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStarredRepoUpdatedAtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStarredRepoUpdatedAtLogic {
	return &GetStarredRepoUpdatedAtLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStarredRepoUpdatedAtLogic) GetStarredRepoUpdatedAt(in *pb.GetStarredRepoUpdatedAtReq) (resp *pb.GetStarredRepoUpdatedAtResp, err error) {
	var createRepoUpdatedAt *model.StarredRepoUpdatedAt
	if createRepoUpdatedAt, err = l.svcCtx.StarredRepoUpdatedAtModel.FindOneByDeveloperId(l.ctx, in.DeveloperId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetStarredRepoUpdatedAtResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetStarredRepoUpdatedAtResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetStarredRepoUpdatedAtResp{
			Code:      http.StatusOK,
			Message:   http.StatusText(http.StatusOK),
			UpdatedAt: createRepoUpdatedAt.UpdatedAt.Unix(),
		}
	}

	err = nil
	return
}

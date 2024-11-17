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

type GetFollowingUpdatedAtLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowingUpdatedAtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingUpdatedAtLogic {
	return &GetFollowingUpdatedAtLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowingUpdatedAtLogic) GetFollowingUpdatedAt(in *pb.GetFollowingUpdatedAtReq) (resp *pb.GetFollowingUpdatedAtResp, err error) {
	var createRepoUpdatedAt *model.FollowingUpdatedAt
	if createRepoUpdatedAt, err = l.svcCtx.FollowingUpdatedAtModel.FindOneByDeveloperId(l.ctx, in.DeveloperId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetFollowingUpdatedAtResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetFollowingUpdatedAtResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetFollowingUpdatedAtResp{
			Code:      http.StatusOK,
			Message:   http.StatusText(http.StatusOK),
			UpdatedAt: createRepoUpdatedAt.UpdatedAt.Unix(),
		}
	}

	err = nil
	return
}

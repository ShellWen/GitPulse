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

type GetFollowerUpdatedAtLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerUpdatedAtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerUpdatedAtLogic {
	return &GetFollowerUpdatedAtLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowerUpdatedAtLogic) GetFollowerUpdatedAt(in *pb.GetFollowerUpdatedAtReq) (resp *pb.GetFollowerUpdatedAtResp, err error) {
	var createRepoUpdatedAt *model.FollowerUpdatedAt
	if createRepoUpdatedAt, err = l.svcCtx.FollowerUpdatedAtModel.FindOneByDeveloperId(l.ctx, in.DeveloperId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetFollowerUpdatedAtResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetFollowerUpdatedAtResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetFollowerUpdatedAtResp{
			Code:      http.StatusOK,
			Message:   http.StatusText(http.StatusOK),
			UpdatedAt: createRepoUpdatedAt.UpdatedAt.Unix(),
		}
	}

	err = nil
	return
}

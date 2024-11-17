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

type GetForkUpdatedAtLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetForkUpdatedAtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetForkUpdatedAtLogic {
	return &GetForkUpdatedAtLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetForkUpdatedAtLogic) GetForkUpdatedAt(in *pb.GetForkUpdatedAtReq) (resp *pb.GetForkUpdatedAtResp, err error) {
	var createRepoUpdatedAt *model.ForkUpdatedAt
	if createRepoUpdatedAt, err = l.svcCtx.ForkUpdatedAtModel.FindOneByRepoId(l.ctx, in.OriginalRepoId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetForkUpdatedAtResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetForkUpdatedAtResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetForkUpdatedAtResp{
			Code:      http.StatusOK,
			Message:   http.StatusText(http.StatusOK),
			UpdatedAt: createRepoUpdatedAt.UpdatedAt.Unix(),
		}
	}

	err = nil
	return
}

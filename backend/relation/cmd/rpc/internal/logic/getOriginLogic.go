package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOriginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOriginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOriginLogic {
	return &GetOriginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOriginLogic) GetOrigin(in *pb.GetOriginReq) (resp *pb.GetOriginResp, err error) {
	var fork *model.Fork
	if fork, err = l.svcCtx.ForkModel.FindOneByForkRepoId(l.ctx, in.ForkRepoId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetOriginResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetOriginResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetOriginResp{
			Code:           http.StatusOK,
			Message:        http.StatusText(http.StatusOK),
			OriginalRepoId: fork.OriginalRepoId,
		}
	}

	err = nil
	return
}

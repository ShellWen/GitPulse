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

type DelForkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelForkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelForkLogic {
	return &DelForkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelForkLogic) DelFork(in *pb.DelForkReq) (resp *pb.DelForkResp, err error) {
	var fork *model.Fork
	if fork, err = l.svcCtx.ForkModel.FindOneByForkRepoId(l.ctx, in.ForkRepoId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.DelForkResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.DelForkResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else if err = l.svcCtx.ForkModel.Delete(l.ctx, fork.DataId); err != nil {
		resp = &pb.DelForkResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.DelForkResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

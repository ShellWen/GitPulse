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

type DelCreateRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCreateRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCreateRepoLogic {
	return &DelCreateRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCreateRepoLogic) DelCreateRepo(in *pb.DelCreateRepoReq) (resp *pb.DelCreateRepoResp, err error) {
	var createRepo *model.CreateRepo
	if createRepo, err = l.svcCtx.CreateRepoModel.FindOneByRepoId(l.ctx, in.RepoId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.DelCreateRepoResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.DelCreateRepoResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else if err = l.svcCtx.CreateRepoModel.Delete(l.ctx, createRepo.DataId); err != nil {
		resp = &pb.DelCreateRepoResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.DelCreateRepoResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

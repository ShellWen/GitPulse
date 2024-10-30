package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/pb"
	"github.com/ShellWen/GitPulse/repo/model"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type DelRepoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelRepoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelRepoByIdLogic {
	return &DelRepoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelRepoByIdLogic) DelRepoById(in *pb.DelRepoByIdReq) (resp *pb.DelRepoByIdResp, err error) {
	var repo *model.Repo

	id := in.Id
	if repo, err = l.svcCtx.RepoModel.FindOneById(l.ctx, id); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			resp = &pb.DelRepoByIdResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		} else {
			resp = &pb.DelRepoByIdResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else if err = l.svcCtx.RepoModel.Delete(l.ctx, repo.DataId); err != nil {
		resp = &pb.DelRepoByIdResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.DelRepoByIdResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

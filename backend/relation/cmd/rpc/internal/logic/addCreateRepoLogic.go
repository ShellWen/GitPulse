package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCreateRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCreateRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCreateRepoLogic {
	return &AddCreateRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------createRepo-----------------------
func (l *AddCreateRepoLogic) AddCreateRepo(in *pb.AddCreateRepoReq) (resp *pb.AddCreateRepoResp, err error) {
	createRepo := &model.CreateRepo{
		DeveloperId: in.DeveloperId,
		RepoId:      in.RepoId,
	}

	if _, err = l.svcCtx.CreateRepoModel.Insert(l.ctx, createRepo); err != nil {
		resp = &pb.AddCreateRepoResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.AddCreateRepoResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

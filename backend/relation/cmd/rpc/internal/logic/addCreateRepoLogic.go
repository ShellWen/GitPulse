package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"time"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

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
		DataCreateAt: time.Now(),
		DataUpdateAt: time.Now(),
		DeveloperId:  in.DeveloperId,
		RepoId:       in.RepoId,
	}

	_, err = l.svcCtx.CreateRepoModel.Insert(l.ctx, createRepo)
	if err != nil {
		return nil, err
	}

	resp = &pb.AddCreateRepoResp{}

	return
}

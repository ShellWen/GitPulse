package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCreatorIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCreatorIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCreatorIdLogic {
	return &GetCreatorIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCreatorIdLogic) GetCreatorId(in *pb.GetCreatorIdReq) (resp *pb.GetCreatorIdResp, err error) {
	var createRepo *model.CreateRepo
	createRepo, err = l.svcCtx.CreateRepoModel.FindOneByRepoId(l.ctx, in.RepoId)
	if err != nil {
		return nil, err
	}

	resp = &pb.GetCreatorIdResp{
		DeveloperId: createRepo.DeveloperId,
	}

	return
}

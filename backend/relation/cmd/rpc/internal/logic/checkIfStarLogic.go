package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/relation/model"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIfStarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckIfStarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIfStarLogic {
	return &CheckIfStarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckIfStarLogic) CheckIfStar(in *pb.CheckIfStarReq) (resp *pb.CheckIfStarResp, err error) {
	_, err = l.svcCtx.StarModel.FindOneByDeveloperIdRepoId(l.ctx, in.DeveloperId, in.RepoId)

	if err == nil {
		resp = &pb.CheckIfStarResp{
			IsStar: true,
		}
	} else if errors.Is(err, model.ErrNotFound) {
		resp = &pb.CheckIfStarResp{
			IsStar: false,
		}
		err = nil
	}

	return
}

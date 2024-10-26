package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"time"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddStarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStarLogic {
	return &AddStarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------star-----------------------
func (l *AddStarLogic) AddStar(in *pb.AddStarReq) (resp *pb.AddStarResp, err error) {
	star := &model.Star{
		DataCreateAt: time.Now(),
		DataUpdateAt: time.Now(),
		DeveloperId:  in.DeveloperId,
		RepoId:       in.RepoId,
	}

	_, err = l.svcCtx.StarModel.Insert(l.ctx, star)
	if err != nil {
		return nil, err
	}

	resp = &pb.AddStarResp{}

	return
}

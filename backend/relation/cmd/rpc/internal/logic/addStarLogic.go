package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

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
		DeveloperId: in.DeveloperId,
		RepoId:      in.RepoId,
	}

	if _, err = l.svcCtx.StarModel.Insert(l.ctx, star); err != nil {
		resp = &pb.AddStarResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.AddStarResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

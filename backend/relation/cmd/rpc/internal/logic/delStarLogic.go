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

type DelStarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelStarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelStarLogic {
	return &DelStarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelStarLogic) DelStar(in *pb.DelStarReq) (resp *pb.DelStarResp, err error) {
	var star *model.Star
	if star, err = l.svcCtx.StarModel.FindOneByDeveloperIdRepoId(l.ctx, in.DeveloperId, in.RepoId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.DelStarResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.DelStarResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else if err = l.svcCtx.StarModel.Delete(l.ctx, star.DataId); err != nil {
		resp = &pb.DelStarResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.DelStarResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

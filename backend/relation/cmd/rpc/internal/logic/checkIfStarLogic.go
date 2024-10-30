package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"
	"github.com/ShellWen/GitPulse/relation/model"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
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
	if _, err = l.svcCtx.StarModel.FindOneByDeveloperIdRepoId(l.ctx, in.DeveloperId, in.RepoId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.CheckIfStarResp{
				Code:    http.StatusOK,
				Message: http.StatusText(http.StatusOK),
				IsStar:  false,
			}
		default:
			resp = &pb.CheckIfStarResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.CheckIfStarResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			IsStar:  true,
		}
	}

	err = nil
	return
}

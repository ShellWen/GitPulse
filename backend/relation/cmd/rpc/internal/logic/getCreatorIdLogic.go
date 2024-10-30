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
	if createRepo, err = l.svcCtx.CreateRepoModel.FindOneByRepoId(l.ctx, in.RepoId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetCreatorIdResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetCreatorIdResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetCreatorIdResp{
			Code:        http.StatusOK,
			Message:     http.StatusText(http.StatusOK),
			DeveloperId: createRepo.DeveloperId,
		}
	}

	err = nil
	return
}

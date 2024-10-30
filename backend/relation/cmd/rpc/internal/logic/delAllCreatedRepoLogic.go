package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllCreatedRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllCreatedRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllCreatedRepoLogic {
	return &DelAllCreatedRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllCreatedRepoLogic) DelAllCreatedRepo(in *pb.DelAllCreatedRepoReq) (resp *pb.DelAllCreatedRepoResp, err error) {
	var createdRepos *[]*model.CreateRepo
	if createdRepos, err = l.svcCtx.CreateRepoModel.SearchCreatedRepo(l.ctx, in.DeveloperId, 1, 9223372036854775807); err != nil {
		resp = &pb.DelAllCreatedRepoResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if len(*createdRepos) == 0 {
		resp = &pb.DelAllCreatedRepoResp{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	} else {
		for _, createdRepo := range *createdRepos {
			if err = l.svcCtx.CreateRepoModel.Delete(l.ctx, createdRepo.DataId); err != nil {
				resp = &pb.DelAllCreatedRepoResp{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
				break
			}
		}
		if err == nil {
			resp = &pb.DelAllCreatedRepoResp{
				Code:    http.StatusOK,
				Message: http.StatusText(http.StatusOK),
			}
		}
	}

	err = nil
	return
}

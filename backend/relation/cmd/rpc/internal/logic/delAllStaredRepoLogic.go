package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllStaredRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllStaredRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllStaredRepoLogic {
	return &DelAllStaredRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllStaredRepoLogic) DelAllStaredRepo(in *pb.DelAllStaredRepoReq) (resp *pb.DelAllStaredRepoResp, err error) {
	var stars *[]*model.Star
	if stars, err = l.svcCtx.StarModel.SearchStaredRepo(l.ctx, in.DeveloperId, 1, 9223372036854775807); err != nil {
		resp = &pb.DelAllStaredRepoResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if len(*stars) == 0 {
		resp = &pb.DelAllStaredRepoResp{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	} else {
		for _, star := range *stars {
			if err = l.svcCtx.StarModel.Delete(l.ctx, star.DataId); err != nil {
				resp = &pb.DelAllStaredRepoResp{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
				break
			}
		}
		if err == nil {
			resp = &pb.DelAllStaredRepoResp{
				Code:    http.StatusOK,
				Message: http.StatusText(http.StatusOK),
			}
		}
	}

	err = nil
	return
}

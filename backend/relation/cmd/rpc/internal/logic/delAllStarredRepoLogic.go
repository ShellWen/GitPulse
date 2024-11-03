package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllStarredRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllStarredRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllStarredRepoLogic {
	return &DelAllStarredRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllStarredRepoLogic) DelAllStarredRepo(in *pb.DelAllStarredRepoReq) (resp *pb.DelAllStarredRepoResp, err error) {
	var stars *[]*model.Star
	if stars, err = l.svcCtx.StarModel.SearchStarredRepo(l.ctx, in.DeveloperId, 1, 9223372036854775807); err != nil {
		resp = &pb.DelAllStarredRepoResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		for _, star := range *stars {
			if err = l.svcCtx.StarModel.Delete(l.ctx, star.DataId); err != nil {
				resp = &pb.DelAllStarredRepoResp{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
				break
			}
		}
		if err == nil {
			resp = &pb.DelAllStarredRepoResp{
				Code:    http.StatusOK,
				Message: http.StatusText(http.StatusOK),
			}
		}
	}

	err = nil
	return
}

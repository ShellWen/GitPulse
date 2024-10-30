package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllForkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllForkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllForkLogic {
	return &DelAllForkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllForkLogic) DelAllFork(in *pb.DelAllForkReq) (resp *pb.DelAllForkResp, err error) {
	var forks *[]*model.Fork
	if forks, err = l.svcCtx.ForkModel.SearchFork(l.ctx, in.OriginalRepoId, 1, 9223372036854775807); err != nil {
		resp = &pb.DelAllForkResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if len(*forks) == 0 {
		resp = &pb.DelAllForkResp{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	} else {
		for _, fork := range *forks {
			if err = l.svcCtx.ForkModel.Delete(l.ctx, fork.DataId); err != nil {
				resp = &pb.DelAllForkResp{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
				break
			}
		}
		if err == nil {
			resp = &pb.DelAllForkResp{
				Code:    http.StatusOK,
				Message: http.StatusText(http.StatusOK),
			}
		}
	}

	err = nil
	return
}

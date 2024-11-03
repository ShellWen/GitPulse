package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllStaringDevLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllStaringDevLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllStaringDevLogic {
	return &DelAllStaringDevLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllStaringDevLogic) DelAllStaringDev(in *pb.DelAllStaringDevReq) (resp *pb.DelAllStaringDevResp, err error) {
	var star *[]*model.Star
	if star, err = l.svcCtx.StarModel.SearchStaringDeveloper(l.ctx, in.RepoId, 1, 9223372036854775807); err != nil {
		resp = &pb.DelAllStaringDevResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		for _, s := range *star {
			if err = l.svcCtx.StarModel.Delete(l.ctx, s.DataId); err != nil {
				resp = &pb.DelAllStaringDevResp{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
				break
			}
		}
		if err == nil {
			resp = &pb.DelAllStaringDevResp{
				Code:    http.StatusOK,
				Message: http.StatusText(http.StatusOK),
			}
		}
	}

	err = nil
	return
}

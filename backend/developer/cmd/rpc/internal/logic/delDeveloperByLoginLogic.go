package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/developer/model"
	"net/http"

	"github.com/ShellWen/GitPulse/developer/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelDeveloperByLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelDeveloperByLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelDeveloperByLoginLogic {
	return &DelDeveloperByLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelDeveloperByLoginLogic) DelDeveloperByLogin(in *pb.DelDeveloperByLoginReq) (resp *pb.DelDeveloperByLoginResp, err error) {
	var developer *model.Developer

	if developer, err = l.svcCtx.DeveloperModel.FindOneByLogin(l.ctx, in.Login); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			resp = &pb.DelDeveloperByLoginResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		} else {
			resp = &pb.DelDeveloperByLoginResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else if err = l.svcCtx.DeveloperModel.Delete(l.ctx, developer.DataId); err != nil {
		resp = &pb.DelDeveloperByLoginResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.DelDeveloperByLoginResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

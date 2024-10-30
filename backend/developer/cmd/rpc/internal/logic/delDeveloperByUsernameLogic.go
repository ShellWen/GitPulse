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

type DelDeveloperByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelDeveloperByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelDeveloperByUsernameLogic {
	return &DelDeveloperByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelDeveloperByUsernameLogic) DelDeveloperByUsername(in *pb.DelDeveloperByUsernameReq) (resp *pb.DelDeveloperByUsernameResp, err error) {
	var developer *model.Developer

	username := in.Username
	if developer, err = l.svcCtx.DeveloperModel.FindOneByUsername(l.ctx, username); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			resp = &pb.DelDeveloperByUsernameResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		} else {
			resp = &pb.DelDeveloperByUsernameResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else if err = l.svcCtx.DeveloperModel.Delete(l.ctx, developer.DataId); err != nil {
		resp = &pb.DelDeveloperByUsernameResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.DelDeveloperByUsernameResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

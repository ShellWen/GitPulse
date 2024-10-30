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

type DelDeveloperByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelDeveloperByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelDeveloperByIdLogic {
	return &DelDeveloperByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelDeveloperByIdLogic) DelDeveloperById(in *pb.DelDeveloperByIdReq) (resp *pb.DelDeveloperByIdResp, err error) {
	var developer *model.Developer

	id := in.Id
	if developer, err = l.svcCtx.DeveloperModel.FindOneById(l.ctx, id); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			resp = &pb.DelDeveloperByIdResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		} else {
			resp = &pb.DelDeveloperByIdResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else if err = l.svcCtx.DeveloperModel.Delete(l.ctx, developer.DataId); err != nil {
		resp = &pb.DelDeveloperByIdResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.DelDeveloperByIdResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

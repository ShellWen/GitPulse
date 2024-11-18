package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/id_generator/internal/svc"
	"github.com/ShellWen/GitPulse/id_generator/pb"
	gonanoid "github.com/matoous/go-nanoid"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIdLogic {
	return &GetIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetIdLogic) GetId(_ *pb.GetIdReq) (*pb.GetIdResp, error) {
	id, err := l.doGetId()
	if err != nil {
		return &pb.GetIdResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	} else {
		return &pb.GetIdResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Id:      id,
		}, nil
	}
}

func (l *GetIdLogic) doGetId() (string, error) {
	return gonanoid.Nanoid()
}

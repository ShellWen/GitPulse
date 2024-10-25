package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/developer/model"

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
	developer, err = l.svcCtx.DeveloperModel.FindOneById(l.ctx, id)
	if err != nil {
		return nil, err
	}

	dataId := developer.DataId
	err = l.svcCtx.DeveloperModel.Delete(l.ctx, dataId)
	if err != nil {
		return nil, err
	}

	resp = &pb.DelDeveloperByIdResp{}

	return
}

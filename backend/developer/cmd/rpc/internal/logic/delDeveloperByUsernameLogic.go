package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/developer/model"

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
	developer, err = l.svcCtx.DeveloperModel.FindOneByUsername(l.ctx, username)
	if err != nil {
		return nil, err
	}

	dataId := developer.DataId
	err = l.svcCtx.DeveloperModel.Delete(l.ctx, dataId)
	if err != nil {
		return nil, err
	}

	resp = &pb.DelDeveloperByUsernameResp{}

	return
}

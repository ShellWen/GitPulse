package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/developer/model"

	"github.com/ShellWen/GitPulse/developer/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeveloperByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeveloperByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeveloperByUsernameLogic {
	return &GetDeveloperByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDeveloperByUsernameLogic) GetDeveloperByUsername(in *pb.GetDeveloperByUsernameReq) (resp *pb.GetDeveloperByUsernameResp, err error) {
	var developer *model.Developer
	developer, err = l.svcCtx.DeveloperModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}

	resp = &pb.GetDeveloperByUsernameResp{
		Developer: &pb.Developer{
			DataId:       developer.DataId,
			DataCreateAt: developer.DataCreateAt.Unix(),
			DataUpdateAt: developer.DataUpdateAt.Unix(),
			Id:           developer.Id,
			Name:         developer.Name,
			Username:     developer.Username,
			AvatarUrl:    developer.AvatarUrl,
			Company:      developer.Company,
			Location:     developer.Location,
			Bio:          developer.Bio,
			Blog:         developer.Blog,
			Email:        developer.Email,
			CreateAt:     developer.CreateAt.Unix(),
			UpdateAt:     developer.UpdateAt.Unix(),
		},
	}

	return
}

package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/developer/model"
	"time"

	"github.com/ShellWen/GitPulse/developer/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeveloperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDeveloperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeveloperLogic {
	return &UpdateDeveloperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDeveloperLogic) UpdateDeveloper(in *pb.UpdateDeveloperReq) (resp *pb.UpdateDeveloperResp, err error) {
	var developer *model.Developer
	developer, err = l.svcCtx.DeveloperModel.FindOneById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	developer.Name = in.Name
	developer.Username = in.Username
	developer.AvatarUrl = in.AvatarUrl
	developer.Company = in.Company
	developer.Location = in.Location
	developer.Bio = in.Bio
	developer.Blog = in.Blog
	developer.Email = in.Email
	developer.CreateAt = time.Unix(in.CreateAt, 0)
	developer.UpdateAt = time.Unix(in.UpdateAt, 0)

	err = l.svcCtx.DeveloperModel.Update(l.ctx, developer)
	if err != nil {
		return nil, err
	}

	resp = &pb.UpdateDeveloperResp{}

	return
}

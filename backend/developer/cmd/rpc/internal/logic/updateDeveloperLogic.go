package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/developer/model"
	"net/http"
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
	if developer, err = l.svcCtx.DeveloperModel.FindOneById(l.ctx, in.Id); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			resp = &pb.UpdateDeveloperResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		} else {
			resp = &pb.UpdateDeveloperResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else if err = l.doUpdateDeveloper(developer, in); err != nil {
		resp = &pb.UpdateDeveloperResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.UpdateDeveloperResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

func (l *UpdateDeveloperLogic) doUpdateDeveloper(developer *model.Developer, in *pb.UpdateDeveloperReq) (err error) {
	developer.Name = in.Name
	developer.Username = in.Username
	developer.AvatarUrl = in.AvatarUrl
	developer.Company = in.Company
	developer.Location = in.Location
	developer.Bio = in.Bio
	developer.Blog = in.Blog
	developer.Email = in.Email
	developer.TwitterUsername = in.TwitterUsername
	developer.Followers = in.Followers
	developer.Following = in.Following
	developer.Repos = in.Repos
	developer.Stars = in.Stars
	developer.Gists = in.Gists
	developer.CreateAt = time.Unix(in.CreateAt, 0)
	developer.UpdateAt = time.Unix(in.UpdateAt, 0)

	err = l.svcCtx.DeveloperModel.Update(l.ctx, developer)
	return
}

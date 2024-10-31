package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/developer/model"
	"net/http"
	"time"

	"github.com/ShellWen/GitPulse/developer/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDeveloperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDeveloperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeveloperLogic {
	return &AddDeveloperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------developer-----------------------
func (l *AddDeveloperLogic) AddDeveloper(in *pb.AddDeveloperReq) (resp *pb.AddDeveloperResp, err error) {
	developer := &model.Developer{
		DataCreateAt:    time.Now(),
		DataUpdateAt:    time.Now(),
		Id:              in.Id,
		Name:            in.Name,
		Username:        in.Username,
		AvatarUrl:       in.AvatarUrl,
		Company:         in.Company,
		Location:        in.Location,
		Bio:             in.Bio,
		Blog:            in.Blog,
		Email:           in.Email,
		TwitterUsername: in.TwitterUsername,
		Followers:       in.Followers,
		Following:       in.Following,
		Repos:           in.Repos,
		Stars:           in.Stars,
		Gists:           in.Gists,
		CreateAt:        time.Unix(in.CreateAt, 0),
		UpdateAt:        time.Unix(in.UpdateAt, 0),
	}

	if _, err = l.svcCtx.DeveloperModel.Insert(l.ctx, developer); err != nil {
		resp = &pb.AddDeveloperResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.AddDeveloperResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

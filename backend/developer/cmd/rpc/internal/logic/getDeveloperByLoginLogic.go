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

type GetDeveloperByLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeveloperByLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeveloperByLoginLogic {
	return &GetDeveloperByLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDeveloperByLoginLogic) GetDeveloperByLogin(in *pb.GetDeveloperByLoginReq) (resp *pb.GetDeveloperByLoginResp, err error) {
	var developer *model.Developer
	if developer, err = l.svcCtx.DeveloperModel.FindOneByLogin(l.ctx, in.Login); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetDeveloperByLoginResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetDeveloperByLoginResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetDeveloperByLoginResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Developer: &pb.Developer{
				DataId:          developer.DataId,
				DataCreatedAt:   developer.DataCreatedAt.Unix(),
				DataUpdatedAt:   developer.DataUpdatedAt.Unix(),
				Id:              developer.Id,
				Name:            developer.Name,
				Login:           developer.Login,
				AvatarUrl:       developer.AvatarUrl,
				Company:         developer.Company,
				Location:        developer.Location,
				Bio:             developer.Bio,
				Blog:            developer.Blog,
				Email:           developer.Email,
				TwitterUsername: developer.TwitterUsername,
				Followers:       developer.Followers,
				Following:       developer.Following,
				Repos:           developer.Repos,
				Stars:           developer.Stars,
				Gists:           developer.Gists,
				CreatedAt:       developer.CreatedAt.Unix(),
				UpdatedAt:       developer.UpdatedAt.Unix(),
			},
		}
	}

	err = nil
	return
}

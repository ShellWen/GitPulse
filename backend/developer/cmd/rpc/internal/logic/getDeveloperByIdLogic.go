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

type GetDeveloperByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeveloperByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeveloperByIdLogic {
	return &GetDeveloperByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDeveloperByIdLogic) GetDeveloperById(in *pb.GetDeveloperByIdReq) (resp *pb.GetDeveloperByIdResp, err error) {
	var developer *model.Developer
	if developer, err = l.svcCtx.DeveloperModel.FindOneById(l.ctx, in.Id); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetDeveloperByIdResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetDeveloperByIdResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetDeveloperByIdResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
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
	}

	err = nil
	return
}

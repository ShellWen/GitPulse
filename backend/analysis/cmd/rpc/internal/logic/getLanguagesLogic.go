package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/model"
	"net/http"

	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLanguagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLanguagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLanguagesLogic {
	return &GetLanguagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLanguagesLogic) GetLanguages(in *pb.GetAnalysisReq) (resp *pb.GetLanguagesResp, err error) {
	var languages *model.Languages
	if languages, err = l.svcCtx.LanguagesModel.FindOneByDeveloperId(l.ctx, in.DeveloperId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.GetLanguagesResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.GetLanguagesResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		resp = &pb.GetLanguagesResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Languages: &pb.Languages{
				DataId:        languages.DataId,
				DataCreatedAt: languages.DataCreatedAt.Unix(),
				DataUpdatedAt: languages.DataUpdatedAt.Unix(),
				Languages:     languages.Languages,
			},
		}
	}

	err = nil
	return
}

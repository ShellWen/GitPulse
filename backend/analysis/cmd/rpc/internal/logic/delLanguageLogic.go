package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/pb"
	"github.com/ShellWen/GitPulse/analysis/model"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLanguageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLanguageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLanguageLogic {
	return &DelLanguageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------languages-----------------------
func (l *DelLanguageLogic) DelLanguage(in *pb.DelAnalysisReq) (resp *pb.DelAnalysisResp, err error) {
	var languages *model.Languages
	if languages, err = l.svcCtx.LanguagesModel.FindOneByDeveloperId(l.ctx, in.DeveloperId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			resp = &pb.DelAnalysisResp{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			resp = &pb.DelAnalysisResp{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else if err = l.svcCtx.LanguagesModel.Delete(l.ctx, languages.DataId); err != nil {
		resp = &pb.DelAnalysisResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.DelAnalysisResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}

	err = nil
	return
}

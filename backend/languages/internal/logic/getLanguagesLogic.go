package logic

import (
	"context"
	githublangsgo "github.com/NDoolan360/github-langs-go"
	"github.com/ShellWen/GitPulse/languages/internal/svc"
	"github.com/ShellWen/GitPulse/languages/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

type GetLanguagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLanguagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLanguagesLogic {
	return &GetLanguagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLanguagesLogic) GetLanguages(req *types.GetLanguages) (resp *[]types.Language, err error) {
	resp = &[]types.Language{}

	var allLang map[string]githublangsgo.Language

	if allLang, err = githublangsgo.GetAllLanguages(); err != nil {
		logx.Error("GetAllLanguages: ", err)
		return
	}

	for langName, lang := range allLang {
		*resp = append(*resp, types.Language{
			Id:    strings.Replace(strings.ToLower(langName), " ", "-", -1),
			Name:  langName,
			Color: lang.Color,
		})
	}

	return
}

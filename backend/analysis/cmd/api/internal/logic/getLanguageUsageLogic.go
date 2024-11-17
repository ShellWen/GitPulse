package logic

import (
	"context"
	githublangsgo "github.com/NDoolan360/github-langs-go"
	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/types"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/analysis"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strings"
	"time"
)

type GetLanguageUsageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLanguageUsageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLanguageUsageLogic {
	return &GetLanguageUsageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLanguageUsageLogic) GetLanguageUsage(req *types.GetLanguageUsageReq) (resp *types.GetLanguageUsageResp, err error) {
	if resp, err = l.doGetLanguageUsage(req); err != nil {
		logx.Error(err)
		return
	}
	return
}

func (l *GetLanguageUsageLogic) doGetLanguageUsage(req *types.GetLanguageUsageReq) (resp *types.GetLanguageUsageResp, err error) {
	var (
		id                     int64
		usage                  map[string]float64
		updatedAt              time.Time
		languageWithPercentage = make([]types.LanguageWithPercentage, 0)
	)

	if id, err = customGithub.GetIdByLogin(l.ctx, req.Login); err != nil {
		return
	}

	if usage, updatedAt, err = l.getLanguageUsageFromRpc(id); err != nil {
		return
	}

	for name, percentage := range usage {
		var color string

		if color, err = l.getLanguageColor(name); err != nil {
			return
		}

		languageWithPercentage = append(languageWithPercentage, types.LanguageWithPercentage{
			Language: types.Language{
				Id:    strings.ToLower(name),
				Name:  name,
				Color: color,
			},
			Percentage: percentage,
		})
	}

	resp = &types.GetLanguageUsageResp{
		LanguageUsage: types.LanguageUsage{
			Id:        id,
			Languages: languageWithPercentage,
			UpdatedAt: updatedAt.Format(time.RFC3339),
		},
	}

	return
}

func (l *GetLanguageUsageLogic) getLanguageUsageFromRpc(id int64) (usage map[string]float64, updatedAt time.Time, err error) {
	var (
		analysisRpcClient = l.svcCtx.AnalysisRpcClient
		rpcUpdateResp     *analysis.UpdateAnalysisResp
		rpcGetResp        *analysis.GetLanguagesResp
	)

	if rpcUpdateResp, err = analysisRpcClient.UpdateLanguage(l.ctx, &analysis.UpdateAnalysisReq{
		DeveloperId: id,
	}); err != nil || rpcUpdateResp.Code != http.StatusOK {
		return
	}

	if rpcGetResp, err = analysisRpcClient.GetLanguages(l.ctx, &analysis.GetAnalysisReq{
		DeveloperId: id,
	}); err != nil || rpcUpdateResp.Code != http.StatusOK {
		return
	}

	if err = jsonx.UnmarshalFromString(rpcGetResp.Languages.Languages, &usage); err != nil {
		return
	}

	updatedAt = time.Unix(rpcGetResp.Languages.DataUpdatedAt, 0)

	return
}

func (l *GetLanguageUsageLogic) getLanguageColor(name string) (color string, err error) {
	var language githublangsgo.Language

	if language, err = githublangsgo.GetLanguage(name); err != nil {
		return
	}

	color = language.Color
	return
}

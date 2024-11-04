package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/model"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/relation"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/repo"
	"math"
	"net/http"
	"time"

	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLanguageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLanguageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLanguageLogic {
	return &UpdateLanguageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLanguageLogic) UpdateLanguage(in *pb.UpdateAnalysisReq) (resp *pb.UpdateAnalysisResp, err error) {
	if err = l.doUpdateLanguage(in.DeveloperId); err != nil {
		resp = &pb.UpdateAnalysisResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.UpdateAnalysisResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}
	return
}

func (l *UpdateLanguageLogic) doUpdateLanguage(id int64) (err error) {
	var (
		relationZrpcClient   = relation.NewRelation(l.svcCtx.RpcClient)
		repoZrpcClient       = repo.NewRepoZrpcClient(l.svcCtx.RpcClient)
		createRepoResp       *relation.SearchCreatedRepoResp
		createRepoIds        []int64
		allLanguageBytes     map[string]int64
		allLanguageRepoCount map[string]int64
		allMetrics           map[string]float64
		totalMetric          float64
		languagesItem        *model.Languages
		jsonBytes            []byte
	)

	if createRepoResp, err = relationZrpcClient.SearchCreatedRepo(l.ctx, &relation.SearchCreatedRepoReq{
		DeveloperId: id,
	}); err != nil {
		return
	}

	createRepoIds = createRepoResp.RepoIds

	for _, repoId := range createRepoIds {
		var (
			repoResp      *repo.GetRepoByIdResp
			languageBytes map[string]int64
		)

		if repoResp, err = repoZrpcClient.GetRepoById(l.ctx, &repo.GetRepoByIdReq{
			Id: repoId,
		}); err != nil {
			return
		}

		if err = json.Unmarshal([]byte(repoResp.Repo.Language), &languageBytes); err != nil {
			return
		}

		for language, bytes := range languageBytes {
			allLanguageBytes[language] += bytes
			allLanguageRepoCount[language]++
		}
	}

	for language, bytes := range allLanguageBytes {
		allMetrics[language] = math.Sqrt(float64(bytes) * float64(allLanguageRepoCount[language]))
		totalMetric += allMetrics[language]
	}

	for language, metric := range allMetrics {
		allMetrics[language] = metric / totalMetric
	}

	if jsonBytes, err = json.Marshal(allMetrics); err != nil {
		return
	}

	if languagesItem, err = l.svcCtx.LanguagesModel.FindOneByDeveloperId(l.ctx, id); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			if _, err = l.svcCtx.LanguagesModel.Insert(l.ctx, &model.Languages{
				DataCreatedAt: time.Now(),
				DataUpdatedAt: time.Now(),
				DeveloperId:   id,
				Languages:     "{}",
			}); err != nil {
				return
			}
		} else {
			return
		}
	}
	languagesItem.Languages = string(jsonBytes)
	if err = l.svcCtx.LanguagesModel.Update(l.ctx, languagesItem); err != nil {
		return
	}

	return
}

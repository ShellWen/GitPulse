package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/model"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	locks "github.com/ShellWen/GitPulse/common/lock"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/relation"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/repo"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
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
		relationZrpcClient   = l.svcCtx.RelationRpcClient
		repoZrpcClient       = l.svcCtx.RepoRpcClient
		updateCreateRepoResp *relation.UpdateCreateRepoResp
		createRepoResp       *relation.SearchCreatedRepoResp
		createRepoIds        []int64
		allLanguageBytes     = make(map[string]int64)
		allLanguageRepoCount = make(map[string]int64)
		allMetrics           = make(map[string]float64)
		totalMetric          float64
		languagesItem        *model.Languages
		jsonBytes            []byte
	)

	lock, err := l.acquireUpdateLanguageLock(id)
	if err != nil {
		return
	}
	defer lock.Unlock()

	needUpdate, err := l.checkIfNeedUpdateLanguages(id)
	if err != nil {
		return
	}

	if !needUpdate {
		return
	}

	if updateCreateRepoResp, err = relationZrpcClient.UpdateCreateRepo(l.ctx, &relation.UpdateCreateRepoReq{
		DeveloperId: id,
	}); err != nil || updateCreateRepoResp.Code != http.StatusOK {
		return
	}

	if createRepoResp, err = relationZrpcClient.SearchCreatedRepo(l.ctx, &relation.SearchCreatedRepoReq{
		DeveloperId: id,
		Limit:       1000,
		Page:        1,
	}); err != nil {
		return
	} else if createRepoResp.Code == http.StatusInternalServerError {
		err = errors.New(createRepoResp.Message)
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
			logx.Error(err)
			err = nil
			continue
		}

		if err = json.Unmarshal([]byte(repoResp.GetRepo().GetLanguage()), &languageBytes); err != nil {
			logx.Error(err)
			err = nil
			continue
		}

		for language, bytes := range languageBytes {
			allLanguageBytes[language] += bytes
			allLanguageRepoCount[language]++
		}
	}

	for language, bytes := range allLanguageBytes {
		allMetrics[language] = math.Sqrt(float64(bytes)) * math.Sqrt(float64(allLanguageRepoCount[language]))
		totalMetric += allMetrics[language]
	}

	for language, metric := range allMetrics {
		allMetrics[language] = (metric / totalMetric) * 100
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
			if languagesItem, err = l.svcCtx.LanguagesModel.FindOneByDeveloperId(l.ctx, id); err != nil {
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

func (l *UpdateLanguageLogic) acquireUpdateLanguageLock(id int64) (*api.Lock, error) {
	lock, err := l.svcCtx.ConsulClient.LockOpts(&api.LockOptions{
		Key:         locks.GetNewLockKey(locks.UpdateLanguages, id),
		Value:       []byte("locked"),
		SessionTTL:  "10s",
		SessionName: uuid.Must(uuid.NewV7()).String(),
	})

	if err != nil {
		logx.Error("Failed to create lock: ", err)
		return nil, err
	}

	_, err = lock.Lock(nil)

	if err != nil {
		logx.Error("Failed to acquire lock: ", err)
		return nil, err
	}

	return lock, nil
}

func (l *UpdateLanguageLogic) checkIfNeedUpdateLanguages(id int64) (bool, error) {
	if languagesUpdatedAt, err := l.svcCtx.LanguagesModel.FindOneByDeveloperId(l.ctx, id); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			return true, nil
		default:
			return false, err
		}
	} else {
		if customGithub.CheckIfDataExpired(languagesUpdatedAt.DataUpdatedAt) {
			return true, nil
		} else {
			return false, nil
		}
	}
}

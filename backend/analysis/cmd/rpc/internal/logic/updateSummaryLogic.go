package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/model"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	"github.com/ShellWen/GitPulse/common/llm"
	locks "github.com/ShellWen/GitPulse/common/lock"
	"github.com/hashicorp/consul/api"
	"github.com/zeromicro/go-zero/core/jsonx"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSummaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSummaryLogic {
	return &UpdateSummaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSummaryLogic) UpdateSummary(in *pb.UpdateAnalysisReq) (*pb.UpdateAnalysisResp, error) {
	err := l.doUpdateSummary(in.DeveloperId)
	if err != nil {
		return &pb.UpdateAnalysisResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	} else {
		return &pb.UpdateAnalysisResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}, nil
	}
}

func (l *UpdateSummaryLogic) doUpdateSummary(id int64) error {
	lock, err := l.acquireUpdateSummaryLock(id)
	if err != nil {
		return err
	}
	defer lock.Unlock()

	needUpdate, err := l.checkIfNeedUpdateSummary(id)
	if err != nil {
		return err
	}

	if !needUpdate {
		return nil
	}

	summary, err := l.getSummaryByLLModel(id)
	if err != nil {
		return err
	}

	summaryItem, err := l.svcCtx.SummaryModel.FindOneByDeveloperId(l.ctx, id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			if _, err = l.svcCtx.SummaryModel.Insert(l.ctx, &model.Summary{
				DataCreatedAt: time.Now(),
				DataUpdatedAt: time.Now(),
				DeveloperId:   id,
				Summary:       summary,
			}); err != nil {
				return err
			}
			summaryItem, err = l.svcCtx.SummaryModel.FindOneByDeveloperId(l.ctx, id)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	summaryItem.Summary = summary
	if err = l.svcCtx.SummaryModel.Update(l.ctx, summaryItem); err != nil {
		return err
	}

	return nil
}

func (l *UpdateSummaryLogic) getSummaryByLLModel(id int64) (string, error) {
	var (
		httpClient     = &http.Client{}
		sparkModelResp llm.SparkModelResp
		respStr        string
		role           = "你是一名专业的 GitHub 数据分析师。" +
			"\n任务：我将提供某 GitHub 用户的简介、贡献内容、使用语言及百分比。" +
			"请分析其内容，为该用户做一个总结，建议包括：擅长的编程领域、" +
			"文化/地区信息、能力水平、个人风格、个人性格等，不要求包含全部方面，可以自由发挥，总结更多方面。" +
			"\n回复格式：请以中文回复。请直接回复总结内容，不要换行，不要分段。纯文本即可，不需要包含其他信息。字数控制在 1000 字以内。"
		allText = ""
		jsonStr string
	)

	if text, err := getTextFromProfile(l.ctx, l.svcCtx, id); err != nil {
		return "", err
	} else {
		allText += text
	}

	if text, err := getTextFromContribution(l.ctx, l.svcCtx, id, 1000); err != nil {
		return "", err
	} else {
		allText += text
	}

	if text, err := getLanguageAsText(l.ctx, l.svcCtx, id); err != nil {
		return "", err
	} else {
		allText += text
	}

	sparkModelData := &llm.SparkModelData{
		MaxTokens:   l.svcCtx.Config.SparkModelConf.MaxTokens,
		TopK:        l.svcCtx.Config.SparkModelConf.TopK,
		Temperature: l.svcCtx.Config.SparkModelConf.Temperature,
		Messages: [2]llm.SparkModelMessage{
			{
				Role:    "system",
				Content: role,
			},
			{
				Role:    "user",
				Content: allText,
			},
		},
		Model: l.svcCtx.Config.SparkModelConf.Model,
	}

	jsonStr, err := jsonx.MarshalToString(sparkModelData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", l.svcCtx.Config.SparkModelConf.Url, strings.NewReader(jsonStr))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+l.svcCtx.Config.SparkModelConf.APIPassword)
	req.Close = true

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &sparkModelResp)
	if err != nil {
		return "", err
	}

	if len(sparkModelResp.Choices) == 0 {
		return "", nil
	}

	respStr = sparkModelResp.Choices[0].Message.Content

	return respStr, nil
}

func (l *UpdateSummaryLogic) acquireUpdateSummaryLock(id int64) (*api.Lock, error) {
	lock, err := l.svcCtx.ConsulClient.LockOpts(&api.LockOptions{
		Key:         locks.GetNewLockKey(locks.UpdateSummary, id),
		Value:       []byte("locked"),
		SessionTTL:  "10s",
		SessionName: "update_summary",
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

func (l *UpdateSummaryLogic) checkIfNeedUpdateSummary(id int64) (bool, error) {
	if summary, err := l.svcCtx.SummaryModel.FindOneByDeveloperId(l.ctx, id); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			return true, nil
		default:
			return false, err
		}
	} else {
		if customGithub.CheckIfDataExpired(summary.DataUpdatedAt) {
			return true, nil
		} else {
			return false, nil
		}
	}
}

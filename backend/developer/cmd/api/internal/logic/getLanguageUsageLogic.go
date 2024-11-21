package logic

import (
	"context"
	"encoding/json"
	"errors"
	githublangsgo "github.com/NDoolan360/github-langs-go"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/analysis"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	"github.com/ShellWen/GitPulse/common/tasks"
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/types"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	zeroErrors "github.com/zeromicro/x/errors"
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

func (l *GetLanguageUsageLogic) GetLanguageUsage(req *types.GetLanguageUsageReq) (*types.LanguageUsage, *types.TaskState, error) {
	reqId := req.TaskId
	id, err := customGithub.GetIdByLogin(l.ctx, req.Login)
	if err != nil {
		logx.Error("Failed to get id by login ", req.Login, err)
		return nil, nil, err
	}
	taskId := tasks.GetNewAPITaskKey(tasks.APIGetLanguage, id, reqId)

	taskInfo, err := l.svcCtx.AsynqInspector.GetTaskInfo(tasks.APITaskQueue, taskId)
	if err != nil {
		switch {
		case errors.Is(err, asynq.ErrTaskNotFound):
			return nil, nil, zeroErrors.New(http.StatusNotFound, "Task not found")
		default:
			return nil, nil, zeroErrors.New(http.StatusInternalServerError, "Failed to get task info: "+err.Error())
		}
	}

	switch taskInfo.State {
	case asynq.TaskStatePending, asynq.TaskStateActive:
		return nil, &types.TaskState{
			State: taskInfo.State.String(),
		}, nil
	case asynq.TaskStateRetry:
		return nil, &types.TaskState{
			State:  taskInfo.State.String(),
			Reason: taskInfo.LastErr,
		}, nil
	case asynq.TaskStateArchived:
		return nil, &types.TaskState{
			State:  "fail",
			Reason: taskInfo.LastErr,
		}, nil
	case asynq.TaskStateCompleted:
		var languages = analysis.Languages{}
		err = json.Unmarshal(taskInfo.Result, &languages)
		if err != nil {
			return nil, nil, zeroErrors.New(http.StatusInternalServerError, "Failed to unmarshal task result: "+err.Error())
		}
		languageUsage, err := l.buildResp(id, &languages)
		if err != nil {
			return nil, nil, zeroErrors.New(http.StatusInternalServerError, "Failed to build response: "+err.Error())
		}
		return languageUsage, nil, nil
	default:
		return nil, nil, zeroErrors.New(http.StatusInternalServerError, "Unexpected task state: "+taskInfo.State.String())
	}
}

func (l *GetLanguageUsageLogic) buildResp(developerId int64, languageUsage *analysis.Languages) (*types.LanguageUsage, error) {
	var (
		usageMap = make(map[string]float64)
		usageArr []types.LanguageWithPercentage
	)

	err := json.Unmarshal([]byte(languageUsage.Languages), &usageMap)
	if err != nil {
		logx.Error("Failed to unmarshal languages ", languageUsage.Languages, err)
		return nil, err
	}

	for name, percentage := range usageMap {
		color, err := l.getLanguageColor(name)
		if err != nil {
			return nil, err
		}

		usageArr = append(usageArr, types.LanguageWithPercentage{
			Language: types.Language{
				Id:    strings.Replace(strings.ToLower(name), " ", "-", -1),
				Name:  name,
				Color: color,
			},
			Percentage: percentage,
		})
	}

	return &types.LanguageUsage{
		Id:        developerId,
		Languages: usageArr,
		UpdatedAt: time.Unix(languageUsage.DataUpdatedAt, 0).Format(time.RFC3339),
	}, nil
}

func (l *GetLanguageUsageLogic) getLanguageColor(name string) (color string, err error) {
	var language githublangsgo.Language

	if language, err = githublangsgo.GetLanguage(name); err != nil {
		return
	}

	color = language.Color
	return
}

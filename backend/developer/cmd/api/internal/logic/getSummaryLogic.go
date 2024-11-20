package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/analysis"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	"github.com/ShellWen/GitPulse/common/tasks"
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/types"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	zeroErrors "github.com/zeromicro/x/errors"
	"net/http"
)

type GetSummaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSummaryLogic {
	return &GetSummaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSummaryLogic) GetSummary(req *types.GetSummaryReq) (*types.Summary, *types.TaskState, error) {
	reqId := req.TaskId
	id, err := customGithub.GetIdByLogin(l.ctx, req.Login)
	if err != nil {
		logx.Error("Failed to get id by login ", req.Login, err)
		return nil, nil, err
	}
	taskId := tasks.GetNewAPITaskKey(tasks.APIGetSummary, id, reqId)

	taskInfo, err := l.svcCtx.AsynqInspector.GetTaskInfo("default", taskId)
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
		var summary = analysis.Summary{}
		err = json.Unmarshal(taskInfo.Result, &summary)
		if err != nil {
			return nil, nil, zeroErrors.New(http.StatusInternalServerError, "Failed to unmarshal task result: "+err.Error())
		}
		return l.buildResp(id, &summary), nil, nil
	default:
		return nil, nil, zeroErrors.New(http.StatusInternalServerError, "Unexpected task state: "+taskInfo.State.String())
	}
}

func (l *GetSummaryLogic) buildResp(developerId int64, summary *analysis.Summary) *types.Summary {
	return &types.Summary{
		Id:      developerId,
		Summary: summary.Summary,
	}
}

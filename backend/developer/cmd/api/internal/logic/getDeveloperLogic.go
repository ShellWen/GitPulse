package logic

import (
	"context"
	"encoding/json"
	"errors"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	"github.com/ShellWen/GitPulse/common/tasks"
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/types"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"github.com/hibiken/asynq"
	zeroErrors "github.com/zeromicro/x/errors"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeveloperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeveloperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeveloperLogic {
	return &GetDeveloperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDeveloperLogic) GetDeveloper(req *types.GetDeveloperReq) (*types.Developer, *types.TaskState, error) {
	reqId := req.TaskId
	id, err := customGithub.GetIdByLogin(l.ctx, req.Login)
	if err != nil {
		logx.Error("Failed to get id by login ", req.Login, err)
		return nil, nil, err
	}
	taskId := tasks.GetNewAPITaskKey(tasks.APIGetDeveloper, id, reqId)

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
		var dev = developer.Developer{}
		err = json.Unmarshal(taskInfo.Result, &dev)
		if err != nil {
			return nil, nil, zeroErrors.New(http.StatusInternalServerError, "Failed to unmarshal task result: "+err.Error())
		}
		return l.buildResp(&dev), nil, nil
	default:
		return nil, nil, zeroErrors.New(http.StatusInternalServerError, "Unexpected task state: "+taskInfo.State.String())
	}
}

func (l *GetDeveloperLogic) buildResp(dev *developer.Developer) (resp *types.Developer) {
	resp = &types.Developer{
		Id:        dev.Id,
		Name:      dev.Name,
		Login:     dev.Login,
		AvatarUrl: dev.AvatarUrl,
		Company:   dev.Company,
		Location:  dev.Location,
		Bio:       dev.Bio,
		Blog:      dev.Blog,
		Email:     dev.Email,
		Followers: dev.Followers,
		Following: dev.Following,
		Stars:     dev.Stars,
		Repos:     dev.Repos,
		Gists:     dev.Gists,
		CreatedAt: time.Unix(dev.CreatedAt, 0).Format(time.RFC3339),
		UpdatedAt: time.Unix(dev.UpdatedAt, 0).Format(time.RFC3339),
	}
	return
}

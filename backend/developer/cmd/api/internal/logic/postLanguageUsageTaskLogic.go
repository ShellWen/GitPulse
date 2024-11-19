package logic

import (
	"context"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	"github.com/ShellWen/GitPulse/common/tasks"
	"github.com/ShellWen/GitPulse/id_generator/idgenerator"
	"github.com/hibiken/asynq"
	zeroErrors "github.com/zeromicro/x/errors"
	"net/http"

	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostLanguageUsageTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostLanguageUsageTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostLanguageUsageTaskLogic {
	return &PostLanguageUsageTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostLanguageUsageTaskLogic) PostLanguageUsageTask(rayId string, req *types.PostTaskReq) (*types.TaskId, error) {
	id, err := customGithub.GetIdByLogin(l.ctx, req.Login)
	if err != nil {
		logx.Error("Failed to get id by login ", req.Login, err)
		return nil, err
	}

	var reqId string
	if rayId != "" {
		reqId = rayId
	} else {
		resp, err := l.svcCtx.IdGeneratorRpcClient.GetId(l.ctx, &idgenerator.GetIdReq{})
		if err != nil {
			logx.Error("Failed to get id ", err)
			return nil, zeroErrors.New(http.StatusInternalServerError, "Failed to get id")
		}
		reqId = resp.Id
	}

	task, taskId, err := tasks.NewAPITask(tasks.APIGetLanguage, id, reqId)
	if err != nil {
		logx.Error("Failed to create task ", err)
		return nil, zeroErrors.New(http.StatusInternalServerError, "Failed to create task")
	}

	_, err = l.svcCtx.AsynqClient.Enqueue(task, asynq.TaskID(taskId), asynq.Retention(tasks.APITaskExpireTime))
	if err != nil {
		logx.Error("Failed to enqueue task ", err)
		return nil, zeroErrors.New(http.StatusInternalServerError, "Failed to enqueue task")
	}

	return &types.TaskId{
		TaskId: reqId,
	}, nil
}
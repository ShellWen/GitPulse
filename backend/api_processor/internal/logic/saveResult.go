package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/api_processor/internal/svc"
	"github.com/hibiken/asynq"
	"time"
)

const expireTime = time.Minute * 10

func SaveResult(_ context.Context, _ *svc.ServiceContext, task *asynq.Task, data []byte) error {
	_, err := task.ResultWriter().Write(data)
	return err
}

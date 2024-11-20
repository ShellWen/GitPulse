package consumer

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/api_processor/internal/logic"
	"github.com/ShellWen/GitPulse/api_processor/internal/svc"
	"github.com/ShellWen/GitPulse/common/tasks"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type APITaskConsumer struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewAPITaskConsumer(ctx context.Context, svc *svc.ServiceContext) *APITaskConsumer {
	return &APITaskConsumer{
		ctx: ctx,
		svc: svc,
	}
}

func (c *APITaskConsumer) Register() *asynq.ServeMux {
	mux := asynq.NewServeMux()

	mux.HandleFunc(tasks.APITaskName, c.Consume)

	return mux
}

func (c *APITaskConsumer) Consume(ctx context.Context, task *asynq.Task) error {
	logx.Info("consume message: ", task.Type(), task.Payload())

	var (
		err  error
		data []byte
		msg  = tasks.APIPayload{}
	)

	err = jsonx.Unmarshal(task.Payload(), &msg)
	if err != nil {
		return err
	}

	switch msg.Type {
	case tasks.APIGetDeveloper:
		logx.Info("consume message: APIGetDeveloper")
		data, err = logic.GetDeveloper(c.ctx, c.svc, msg.Id)
	case tasks.APIGetLanguage:
		logx.Info("consume message: APIGetLanguage")
		data, err = logic.GetLanguage(c.ctx, c.svc, msg.Id)
	case tasks.APIGetPulsePoint:
		logx.Info("consume message: APIGetPulsePoint")
		data, err = logic.GetPulsePoint(c.ctx, c.svc, msg.Id)
	case tasks.APIGetRegion:
		logx.Info("consume message: APIGetRegion")
		data, err = logic.GetRegion(c.ctx, c.svc, msg.Id)
	case tasks.APIGetSummary:
		logx.Info("consume message: APIGetSummary")
		data, err = logic.GetSummary(c.ctx, c.svc, msg.Id)
	default:
		err = errors.New("unexpected message type: " + strconv.FormatInt(int64(msg.Type), 10))
	}
	if err != nil {
		return err
	}

	err = logic.SaveResult(c.ctx, c.svc, task, data)

	return err
}

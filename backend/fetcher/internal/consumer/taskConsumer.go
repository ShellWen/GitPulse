package consumer

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/common/tasks"
	"github.com/ShellWen/GitPulse/fetcher/internal/logic"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type FetcherTaskConsumer struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewFetcherTaskConsumer(ctx context.Context, svc *svc.ServiceContext) *FetcherTaskConsumer {
	return &FetcherTaskConsumer{
		ctx: ctx,
		svc: svc,
	}
}

func (c *FetcherTaskConsumer) Register() *asynq.ServeMux {
	mux := asynq.NewServeMux()

	mux.HandleFunc(tasks.FetcherTaskName, c.Consume)

	return mux
}

func (c *FetcherTaskConsumer) Consume(ctx context.Context, task *asynq.Task) (err error) {
	logx.Info("consume message: ", task.Type(), task.Payload())
	msg := tasks.FetchPayload{}
	if err = jsonx.Unmarshal(task.Payload(), &msg); err != nil {
		return
	}

	switch msg.Type {
	case tasks.FetchDeveloper:
		err = logic.FetchDeveloper(c.ctx, c.svc, msg.Id)
	case tasks.FetchRepo:
		err = logic.FetchRepo(c.ctx, c.svc, msg.Id)
	case tasks.FetchCreatedRepo:
		err = logic.FetchCreatedRepo(c.ctx, c.svc, msg.Id)
	case tasks.FetchStarredRepo:
		err = logic.FetchStarredRepo(c.ctx, c.svc, msg.Id)
	case tasks.FetchFollower:
		err = logic.FetchFollower(c.ctx, c.svc, msg.Id)
	case tasks.FetchFollowing:
		err = logic.FetchFollowing(c.ctx, c.svc, msg.Id)
	case tasks.FetchFork:
		err = logic.FetchFork(c.ctx, c.svc, msg.Id)
	case tasks.FetchIssuePROfUser:
		_, err = logic.FetchIssuePROfUser(c.ctx, c.svc, msg.Id, msg.UpdateAfter, msg.SearchLimit)
	case tasks.FetchCommentOfUser:
		_, err = logic.FetchCommentOfUser(c.ctx, c.svc, msg.Id, msg.UpdateAfter, msg.SearchLimit)
	case tasks.FetchReviewOfUser:
		_, err = logic.FetchReviewOfUser(c.ctx, c.svc, msg.Id, msg.UpdateAfter, msg.SearchLimit)
	default:
		err = errors.New("unexpected message type: " + strconv.FormatInt(int64(msg.Type), 10))
	}

	return
}

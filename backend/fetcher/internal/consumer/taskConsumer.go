package consumer

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/common/message"
	"github.com/ShellWen/GitPulse/fetcher/internal/logic"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
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

func (c *FetcherTaskConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	logx.Info("consume message: ", value)

	msg := message.FetcherTask{}
	if err = jsonx.UnmarshalFromString(value, &msg); err != nil {
		return
	}

	switch msg.Type {
	case message.FetchDeveloper:
		err = logic.FetchDeveloper(c.ctx, c.svc, msg.Id)
	case message.FetchRepo:
		err = logic.FetchRepo(c.ctx, c.svc, msg.Id)
	case message.FetchCreatedRepo:
		err = logic.FetchCreatedRepo(c.ctx, c.svc, msg.Id)
	case message.FetchStarredRepo:
		err = logic.FetchStarredRepo(c.ctx, c.svc, msg.Id)
	case message.FetchFollow:
		err = logic.FetchFollow(c.ctx, c.svc, msg.Id)
	case message.FetchFollower:
		err = logic.FetchFollower(c.ctx, c.svc, msg.Id)
	case message.FetchFollowing:
		err = logic.FetchFollowing(c.ctx, c.svc, msg.Id)
	case message.FetchFork:
		err = logic.FetchFork(c.ctx, c.svc, msg.Id)
	case message.FetchContributionOfUser:
		err = logic.FetchContributionOfUser(c.ctx, c.svc, msg.Id)
	case message.FetchIssuePROfUser:
		err = logic.FetchIssuePROfUser(c.ctx, c.svc, msg.Id)
	case message.FetchCommentOfUser:
		err = logic.FetchCommentOfUser(c.ctx, c.svc, msg.Id)
	default:
		err = errors.New("unexpected message type: " + strconv.FormatInt(int64(msg.Type), 10))
	}

	return
}

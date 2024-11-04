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
		if err = logic.FetchDeveloper(c.ctx, c.svc, msg.Id); err != nil {
			_ = c.svc.KqDeveloperUpdateCompletePusher.Push(c.ctx, value)
			return
		}
		err = c.svc.KqDeveloperUpdateCompletePusher.Push(c.ctx, value)
	case message.FetchRepo:
		if err = logic.FetchRepo(c.ctx, c.svc, msg.Id); err != nil {
			_ = c.svc.KqRepoUpdateCompletePusher.Push(c.ctx, value)
			return
		}
		err = c.svc.KqRepoUpdateCompletePusher.Push(c.ctx, value)
	case message.FetchCreatedRepo:
		if err = logic.FetchCreatedRepo(c.ctx, c.svc, msg.Id); err != nil {
			_ = c.svc.KqRelationUpdateCompletePusher.Push(c.ctx, value)
			return
		}
		err = c.svc.KqRelationUpdateCompletePusher.Push(c.ctx, value)
	case message.FetchStarredRepo:
		if err = logic.FetchStarredRepo(c.ctx, c.svc, msg.Id); err != nil {
			_ = c.svc.KqRelationUpdateCompletePusher.Push(c.ctx, value)
			return
		}
		err = c.svc.KqRelationUpdateCompletePusher.Push(c.ctx, value)
	case message.FetchFollow:
		if err = logic.FetchFollow(c.ctx, c.svc, msg.Id); err != nil {
			_ = c.svc.KqRelationUpdateCompletePusher.Push(c.ctx, value)
			return
		}
		err = c.svc.KqRelationUpdateCompletePusher.Push(c.ctx, value)
	case message.FetchFollower:
		if err = logic.FetchFollower(c.ctx, c.svc, msg.Id); err != nil {
			_ = c.svc.KqRelationUpdateCompletePusher.Push(c.ctx, value)
			return
		}
		err = c.svc.KqRelationUpdateCompletePusher.Push(c.ctx, value)
	case message.FetchFollowing:
		if err = logic.FetchFollowing(c.ctx, c.svc, msg.Id); err != nil {
			_ = c.svc.KqRelationUpdateCompletePusher.Push(c.ctx, value)
			return
		}
		err = c.svc.KqRelationUpdateCompletePusher.Push(c.ctx, value)
	case message.FetchFork:
		if err = logic.FetchFork(c.ctx, c.svc, msg.Id); err != nil {
			_ = c.svc.KqRelationUpdateCompletePusher.Push(c.ctx, value)
			return
		}
		err = c.svc.KqRelationUpdateCompletePusher.Push(c.ctx, value)
	case message.FetchContributionOfUser:
		if err = logic.FetchContributionOfUser(c.ctx, c.svc, msg.Id); err != nil {
			_ = c.svc.KqContributionUpdateCompletePusher.Push(c.ctx, value)
			return
		}
		err = c.svc.KqContributionUpdateCompletePusher.Push(c.ctx, value)
	case message.FetchIssuePROfUser:
		if err = logic.FetchIssuePROfUser(c.ctx, c.svc, msg.Id); err != nil {
			_ = c.svc.KqContributionUpdateCompletePusher.Push(c.ctx, value)
			return
		}
		err = c.svc.KqContributionUpdateCompletePusher.Push(c.ctx, value)
	case message.FetchCommentOfUser:
		if err = logic.FetchCommentOfUser(c.ctx, c.svc, msg.Id); err != nil {
			_ = c.svc.KqContributionUpdateCompletePusher.Push(c.ctx, value)
			return
		}
		err = c.svc.KqContributionUpdateCompletePusher.Push(c.ctx, value)
	default:
		err = errors.New("unexpected message type: " + strconv.FormatInt(int64(msg.Type), 10))
	}

	return
}

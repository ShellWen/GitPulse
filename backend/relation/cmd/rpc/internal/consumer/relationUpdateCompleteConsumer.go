package consumer

import (
	"context"
	"github.com/ShellWen/GitPulse/common/message"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type ContributionUpdateCompleteConsumer struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewContributionUpdateCompleteConsumer(ctx context.Context, svc *svc.ServiceContext) *ContributionUpdateCompleteConsumer {
	return &ContributionUpdateCompleteConsumer{
		ctx: ctx,
		svc: svc,
	}
}

func (c *ContributionUpdateCompleteConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	logx.Info("consume message: ", value)

	var (
		completedFetchTask *message.FetcherTask
	)

	if err = jsonx.UnmarshalFromString(value, &completedFetchTask); err != nil {
		logx.Error("UnmarshalFromString error: ", err)
		return
	}

	switch completedFetchTask.Type {
	case message.FetchCreatedRepo:
		c.syncCreatedRepo(ctx, completedFetchTask)
	case message.FetchStarredRepo:
		c.syncStarredRepo(ctx, completedFetchTask)
	case message.FetchFork:
		c.syncFork(ctx, completedFetchTask)
	case message.FetchFollowing:
		c.syncFollowing(ctx, completedFetchTask)
	case message.FetchFollower:
		c.syncFollower(ctx, completedFetchTask)
	default:
		logx.Error("Invalid message type: ", completedFetchTask.Type)
		return
	}

	return
}

func (c *ContributionUpdateCompleteConsumer) syncCreatedRepo(ctx context.Context, completedFetchTask *message.FetcherTask) {
	if c.svc.CreatedRepoUpdatedChan[completedFetchTask.Id] == nil {
		c.svc.CreatedRepoUpdatedChan[completedFetchTask.Id] = make(chan struct{})
	}

	select {
	case c.svc.CreatedRepoUpdatedChan[completedFetchTask.Id] <- struct{}{}:
	default:
	}
}

func (c *ContributionUpdateCompleteConsumer) syncStarredRepo(ctx context.Context, completedFetchTask *message.FetcherTask) {
	if c.svc.StarredRepoUpdatedChan[completedFetchTask.Id] == nil {
		c.svc.StarredRepoUpdatedChan[completedFetchTask.Id] = make(chan struct{})
	}

	select {
	case c.svc.StarredRepoUpdatedChan[completedFetchTask.Id] <- struct{}{}:
	default:
	}
}

func (c *ContributionUpdateCompleteConsumer) syncFork(ctx context.Context, completedFetchTask *message.FetcherTask) {
	if c.svc.ForkUpdatedChan[completedFetchTask.Id] == nil {
		c.svc.ForkUpdatedChan[completedFetchTask.Id] = make(chan struct{})
	}

	select {
	case c.svc.ForkUpdatedChan[completedFetchTask.Id] <- struct{}{}:
	default:
	}
}

func (c *ContributionUpdateCompleteConsumer) syncFollowing(ctx context.Context, completedFetchTask *message.FetcherTask) {
	if c.svc.FollowingUpdatedChan[completedFetchTask.Id] == nil {
		c.svc.FollowingUpdatedChan[completedFetchTask.Id] = make(chan struct{})
	}

	select {
	case c.svc.FollowingUpdatedChan[completedFetchTask.Id] <- struct{}{}:
	default:
	}
}

func (c *ContributionUpdateCompleteConsumer) syncFollower(ctx context.Context, completedFetchTask *message.FetcherTask) {
	if c.svc.FollowerUpdatedChan[completedFetchTask.Id] == nil {
		c.svc.FollowerUpdatedChan[completedFetchTask.Id] = make(chan struct{})
	}

	select {
	case c.svc.FollowerUpdatedChan[completedFetchTask.Id] <- struct{}{}:
	default:
	}
}

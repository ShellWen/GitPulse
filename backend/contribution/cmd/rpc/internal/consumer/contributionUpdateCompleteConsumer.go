package consumer

import (
	"context"
	"github.com/ShellWen/GitPulse/common/message"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/svc"
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
	case message.FetchContributionOfUser:
		c.syncAllContribution(ctx, completedFetchTask)
	case message.FetchIssuePROfUser:
		c.syncIssuePr(ctx, completedFetchTask)
	case message.FetchCommentOfUser:
		c.syncCommentReview(ctx, completedFetchTask)
	default:
		logx.Error("Invalid message type: ", completedFetchTask.Type)
		return
	}

	return
}

func (c *ContributionUpdateCompleteConsumer) syncIssuePr(ctx context.Context, completedFetchTask *message.FetcherTask) {
	if c.svc.IssuePrUpdatedChan[completedFetchTask.Id] == nil {
		c.svc.IssuePrUpdatedChan[completedFetchTask.Id] = make(chan struct{})
	}

	select {
	case c.svc.IssuePrUpdatedChan[completedFetchTask.Id] <- struct{}{}:
	default:
	}
}

func (c *ContributionUpdateCompleteConsumer) syncCommentReview(ctx context.Context, completedFetchTask *message.FetcherTask) {
	if c.svc.CommentReviewUpdatedChan[completedFetchTask.Id] == nil {
		c.svc.CommentReviewUpdatedChan[completedFetchTask.Id] = make(chan struct{})
	}

	select {
	case c.svc.CommentReviewUpdatedChan[completedFetchTask.Id] <- struct{}{}:
	default:
	}
}

func (c *ContributionUpdateCompleteConsumer) syncAllContribution(ctx context.Context, completedFetchTask *message.FetcherTask) {
	if c.svc.AllUpdatedChan[completedFetchTask.Id] == nil {
		c.svc.AllUpdatedChan[completedFetchTask.Id] = make(chan struct{})
	}

	select {
	case c.svc.AllUpdatedChan[completedFetchTask.Id] <- struct{}{}:
	default:
	}
	return
}

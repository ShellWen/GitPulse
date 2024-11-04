package consumer

import (
	"context"
	"github.com/ShellWen/GitPulse/common/message"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type RepoUpdateCompleteConsumer struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewRepoUpdateCompleteConsumer(ctx context.Context, svc *svc.ServiceContext) *RepoUpdateCompleteConsumer {
	return &RepoUpdateCompleteConsumer{
		ctx: ctx,
		svc: svc,
	}
}

func (c *RepoUpdateCompleteConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	logx.Info("consume message: ", value)

	var (
		completedFetchTask *message.FetcherTask
	)

	if err = jsonx.UnmarshalFromString(value, &completedFetchTask); err != nil {
		logx.Error("UnmarshalFromString error: ", err)
		return
	}

	if completedFetchTask.Type != message.FetchRepo {
		logx.Error("Invalid message type: ", completedFetchTask.Type)
		return
	}

	if c.svc.RepoUpdatedChan[completedFetchTask.Id] == nil {
		c.svc.RepoUpdatedChan[completedFetchTask.Id] = make(chan struct{})
	}

	select {
	case c.svc.RepoUpdatedChan[completedFetchTask.Id] <- struct{}{}:
	default:
	}

	return
}

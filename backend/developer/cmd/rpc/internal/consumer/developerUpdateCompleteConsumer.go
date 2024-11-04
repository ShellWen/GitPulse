package consumer

import (
	"context"
	"github.com/ShellWen/GitPulse/common/message"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeveloperUpdateCompleteConsumer struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewDeveloperUpdateCompleteConsumer(ctx context.Context, svc *svc.ServiceContext) *DeveloperUpdateCompleteConsumer {
	return &DeveloperUpdateCompleteConsumer{
		ctx: ctx,
		svc: svc,
	}
}

func (c *DeveloperUpdateCompleteConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	logx.Info("consume message: ", value)

	var (
		completedFetchTask *message.FetcherTask
	)

	if err = jsonx.UnmarshalFromString(value, &completedFetchTask); err != nil {
		logx.Error("UnmarshalFromString error: ", err)
		return
	}

	if completedFetchTask.Type != message.FetchDeveloper {
		logx.Error("Invalid message type: ", completedFetchTask.Type)
		return
	}

	if c.svc.DeveloperUpdatedChan[completedFetchTask.Id] == nil {
		c.svc.DeveloperUpdatedChan[completedFetchTask.Id] = make(chan struct{})
	}

	select {
	case c.svc.DeveloperUpdatedChan[completedFetchTask.Id] <- struct{}{}:
	default:
	}

	return
}

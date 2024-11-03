package consumer

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/model"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type ForkUpdateConsumer struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewForkUpdateConsumer(ctx context.Context, svc *svc.ServiceContext) *ForkUpdateConsumer {
	return &ForkUpdateConsumer{
		ctx: ctx,
		svc: svc,
	}
}

func (c *ForkUpdateConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	logx.Info("consume message: ", value)

	var newFork *model.Fork

	if err = jsonx.UnmarshalFromString(value, &newFork); err != nil {
		logx.Error("UnmarshalFromString error: ", err)
		return
	}

	err = insertNewFork(c, newFork)

	return
}

func insertNewFork(c *ForkUpdateConsumer, newFork *model.Fork) error {
	if _, err := c.svc.ForkModel.Insert(c.ctx, newFork); err != nil {
		logx.Error("Insert error: ", err)
		return err
	}

	return nil
}

package consumer

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/model"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type StarUpdateConsumer struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewStarUpdateConsumer(ctx context.Context, svc *svc.ServiceContext) *StarUpdateConsumer {
	return &StarUpdateConsumer{
		ctx: ctx,
		svc: svc,
	}
}

func (c *StarUpdateConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	logx.Info("consume message: ", value)

	var newStar *model.Star

	if err = jsonx.UnmarshalFromString(value, &newStar); err != nil {
		logx.Error("UnmarshalFromString error: ", err)
		return
	}

	err = insertNewStar(c, newStar)

	return
}

func insertNewStar(c *StarUpdateConsumer, newStar *model.Star) error {
	if _, err := c.svc.StarModel.Insert(c.ctx, newStar); err != nil {
		logx.Error("Insert error: ", err)
		return err
	}

	return nil
}

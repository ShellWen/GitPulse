package consumer

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/model"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type FollowUpdateConsumer struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewFollowUpdateConsumer(ctx context.Context, svc *svc.ServiceContext) *FollowUpdateConsumer {
	return &FollowUpdateConsumer{
		ctx: ctx,
		svc: svc,
	}
}

func (c *FollowUpdateConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	logx.Info("consume message: ", value)

	var newFollow *model.Follow

	if err = jsonx.UnmarshalFromString(value, &newFollow); err != nil {
		logx.Error("UnmarshalFromString error: ", err)
		return
	}

	err = insertNewFollow(c, newFollow)

	return
}

func insertNewFollow(c *FollowUpdateConsumer, newFollow *model.Follow) error {
	if _, err := c.svc.FollowModel.Insert(c.ctx, newFollow); err != nil {
		logx.Error("Insert error: ", err)
		return err
	}

	return nil
}

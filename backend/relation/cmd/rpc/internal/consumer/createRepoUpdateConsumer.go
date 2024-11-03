package consumer

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/model"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRepoUpdateConsumer struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewCreateRepoUpdateConsumer(ctx context.Context, svc *svc.ServiceContext) *CreateRepoUpdateConsumer {
	return &CreateRepoUpdateConsumer{
		ctx: ctx,
		svc: svc,
	}
}

func (c *CreateRepoUpdateConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	logx.Info("consume message: ", value)

	var newCreateRepo *model.CreateRepo

	if err = jsonx.UnmarshalFromString(value, &newCreateRepo); err != nil {
		logx.Error("UnmarshalFromString error: ", err)
		return
	}

	err = insertNewCreateRepo(c, newCreateRepo)

	return
}

func insertNewCreateRepo(c *CreateRepoUpdateConsumer, newCreateRepo *model.CreateRepo) error {
	if _, err := c.svc.CreateRepoModel.Insert(c.ctx, newCreateRepo); err != nil {
		logx.Error("Insert error: ", err)
		return err
	}

	return nil
}

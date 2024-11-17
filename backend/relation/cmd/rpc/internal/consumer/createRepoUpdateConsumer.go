package consumer

import (
	"context"
	"errors"
	locks "github.com/ShellWen/GitPulse/common/lock"
	"github.com/ShellWen/GitPulse/common/tasks"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/model"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
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

	if newCreateRepo.DataId == tasks.FetchCreatedRepoCompletedDataId {
		if err = updateCreatedRepoUpdatedAt(c, newCreateRepo.DeveloperId); err != nil {
			return
		}
		if err = unblockCreatedRepoUpdateLock(c, newCreateRepo.DeveloperId); err != nil {
			return
		}
	} else {
		err = insertNewCreateRepo(c, newCreateRepo)
		return
	}

	return
}

func insertNewCreateRepo(c *CreateRepoUpdateConsumer, newCreateRepo *model.CreateRepo) error {
	if _, err := c.svc.CreateRepoModel.Insert(c.ctx, newCreateRepo); err != nil {
		logx.Error("Insert error: ", err)
		return err
	}

	return nil
}

func unblockCreatedRepoUpdateLock(c *CreateRepoUpdateConsumer, developerId int64) error {
	if _, err := c.svc.RedisClient.LpushCtx(c.ctx, locks.GetNewLockKey(locks.UpdateCreatedRepo, developerId), ""); err != nil {
		logx.Error("LpushCtx error: ", err)
		return err
	}

	logx.Info("Unblock createRepo update lock success")
	return nil
}

func updateCreatedRepoUpdatedAt(c *CreateRepoUpdateConsumer, developerId int64) error {
	createRepoUpdatedAt, err := c.svc.CreatedRepoUpdatedAtModel.FindOneByDeveloperId(c.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			createRepoUpdatedAt = &model.CreatedRepoUpdatedAt{
				DeveloperId: developerId,
				UpdatedAt:   time.Now(),
			}
			if _, err = c.svc.CreatedRepoUpdatedAtModel.Insert(c.ctx, createRepoUpdatedAt); err != nil {
				logx.Error("Insert error: ", err)
				return err
			}
			return nil
		default:
			logx.Error("FindOneByDeveloperId error: ", err)
			return err
		}
	}

	createRepoUpdatedAt.UpdatedAt = time.Now()

	if err = c.svc.CreatedRepoUpdatedAtModel.Update(c.ctx, createRepoUpdatedAt); err != nil {
		logx.Error("Update error: ", err)
		return err
	}

	logx.Info("Update createRepo updatedAt success")
	return nil
}

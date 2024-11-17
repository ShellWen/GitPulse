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

	if newFork.DataId == tasks.FetchForkCompletedDataId {
		if err = updateForkUpdatedAt(c, newFork.OriginalRepoId); err != nil {
			return
		}
		if err = unblockForkUpdateLock(c, newFork.OriginalRepoId); err != nil {
			return
		}
	} else {
		err = insertNewFork(c, newFork)
		return
	}

	return
}

func insertNewFork(c *ForkUpdateConsumer, newFork *model.Fork) error {
	if _, err := c.svc.ForkModel.Insert(c.ctx, newFork); err != nil {
		logx.Error("Insert error: ", err)
		return err
	}

	return nil
}

func unblockForkUpdateLock(c *ForkUpdateConsumer, developerId int64) error {
	if _, err := c.svc.RedisClient.LpushCtx(c.ctx, locks.GetNewLockKey(locks.UpdateFork, developerId), ""); err != nil {
		logx.Error("LpushCtx error: ", err)
		return err
	}

	logx.Info("Unblock fork update lock success")
	return nil
}

func updateForkUpdatedAt(c *ForkUpdateConsumer, repoId int64) error {
	forkUpdatedAt, err := c.svc.ForkUpdatedAtModel.FindOneByRepoId(c.ctx, repoId)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			forkUpdatedAt = &model.ForkUpdatedAt{
				RepoId:    repoId,
				UpdatedAt: time.Now(),
			}
			if _, err = c.svc.ForkUpdatedAtModel.Insert(c.ctx, forkUpdatedAt); err != nil {
				logx.Error("Insert error: ", err)
				return err
			}
			return nil
		default:
			logx.Error("FindOneByRepoId error: ", err)
			return err
		}
	}

	forkUpdatedAt.UpdatedAt = time.Now()

	if err = c.svc.ForkUpdatedAtModel.Update(c.ctx, forkUpdatedAt); err != nil {
		logx.Error("Update error: ", err)
		return err
	}

	return nil
}

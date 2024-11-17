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

	switch newStar.DataId {
	case tasks.FetchStarredRepoCompletedDataId:
		if err = updateStarredRepoUpdatedAt(c, newStar.DeveloperId); err != nil {
			return
		}
		if err = unblockStarredRepoUpdateLock(c, newStar.DeveloperId); err != nil {
			return
		}
	default:
		err = insertNewStar(c, newStar)
		return
	}

	return
}

func insertNewStar(c *StarUpdateConsumer, newStar *model.Star) error {
	if _, err := c.svc.StarModel.Insert(c.ctx, newStar); err != nil {
		logx.Error("Insert error: ", err)
		return err
	}

	return nil
}

func unblockStarredRepoUpdateLock(c *StarUpdateConsumer, developerId int64) error {
	if _, err := c.svc.RedisClient.LpushCtx(c.ctx, locks.GetNewLockKey(locks.UpdateStarredRepo, developerId), ""); err != nil {
		logx.Error("LpushCtx error: ", err)
		return err
	}

	logx.Info("Unblock starredRepo update lock success")
	return nil
}

func updateStarredRepoUpdatedAt(c *StarUpdateConsumer, developerId int64) error {
	starredRepoUpdatedAt, err := c.svc.StarredRepoUpdatedAtModel.FindOneByDeveloperId(c.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			starredRepoUpdatedAt = &model.StarredRepoUpdatedAt{
				DeveloperId: developerId,
				UpdatedAt:   time.Now(),
			}
			if _, err = c.svc.StarredRepoUpdatedAtModel.Insert(c.ctx, starredRepoUpdatedAt); err != nil {
				logx.Error("Insert error: ", err)
				return err
			}
			return nil
		default:
			logx.Error("FindOneByDeveloperId error: ", err)
			return err
		}
	}

	starredRepoUpdatedAt.UpdatedAt = time.Now()

	if err = c.svc.StarredRepoUpdatedAtModel.Update(c.ctx, starredRepoUpdatedAt); err != nil {
		logx.Error("Update error: ", err)
		return err
	}

	logx.Info("Update starredRepo update at success")
	return nil
}

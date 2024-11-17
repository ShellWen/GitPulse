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

	switch newFollow.DataId {
	case tasks.FetchFollowingCompletedDataId:
		if err = updateFollowingUpdatedAt(c, newFollow.FollowerId); err != nil {
			return
		}
		if err = unblockFollowingUpdateLock(c, newFollow.FollowerId); err != nil {
			return
		}
	case tasks.FetchFollowerCompletedDataId:
		if err = updateFollowerUpdatedAt(c, newFollow.FollowingId); err != nil {
			return
		}
		if err = unblockFollowerUpdateLock(c, newFollow.FollowingId); err != nil {
			return
		}
	default:
		err = insertNewFollow(c, newFollow)
		return
	}

	return
}

func insertNewFollow(c *FollowUpdateConsumer, newFollow *model.Follow) error {
	if _, err := c.svc.FollowModel.Insert(c.ctx, newFollow); err != nil {
		logx.Error("Insert error: ", err)
		return err
	}

	return nil
}

func unblockFollowingUpdateLock(c *FollowUpdateConsumer, developerId int64) error {
	if _, err := c.svc.RedisClient.LpushCtx(c.ctx, locks.GetNewLockKey(locks.UpdateFollowing, developerId), ""); err != nil {
		logx.Error("LpushCtx error: ", err)
		return err
	}

	logx.Info("Unblock following update lock success")
	return nil
}

func unblockFollowerUpdateLock(c *FollowUpdateConsumer, developerId int64) error {
	if _, err := c.svc.RedisClient.LpushCtx(c.ctx, locks.GetNewLockKey(locks.UpdateFollower, developerId), ""); err != nil {
		logx.Error("LpushCtx error: ", err)
		return err
	}

	logx.Info("Unblock follower update lock success")
	return nil
}

func updateFollowerUpdatedAt(c *FollowUpdateConsumer, developerId int64) error {
	followerUpdatedAt, err := c.svc.FollowerUpdatedAtModel.FindOneByDeveloperId(c.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			followerUpdatedAt = &model.FollowerUpdatedAt{
				DeveloperId: developerId,
				UpdatedAt:   time.Now(),
			}
			if _, err = c.svc.FollowerUpdatedAtModel.Insert(c.ctx, followerUpdatedAt); err != nil {
				logx.Error("Insert error: ", err)
				return err
			}
			return nil
		default:
			logx.Error("FindOneByDeveloperId error: ", err)
			return err
		}
	}

	followerUpdatedAt.UpdatedAt = time.Now()

	if err = c.svc.FollowerUpdatedAtModel.Update(c.ctx, followerUpdatedAt); err != nil {
		logx.Error("Update error: ", err)
		return err
	}

	logx.Info("Update follower update at success")
	return nil
}

func updateFollowingUpdatedAt(c *FollowUpdateConsumer, developerId int64) error {
	followingUpdatedAt, err := c.svc.FollowingUpdatedAtModel.FindOneByDeveloperId(c.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			followingUpdatedAt = &model.FollowingUpdatedAt{
				DeveloperId: developerId,
				UpdatedAt:   time.Now(),
			}
			if _, err = c.svc.FollowingUpdatedAtModel.Insert(c.ctx, followingUpdatedAt); err != nil {
				logx.Error("Insert error: ", err)
				return err
			}
			return nil
		default:
			logx.Error("FindOneByDeveloperId error: ", err)
			return err
		}
	}

	followingUpdatedAt.UpdatedAt = time.Now()

	if err = c.svc.FollowingUpdatedAtModel.Update(c.ctx, followingUpdatedAt); err != nil {
		logx.Error("Update error: ", err)
		return err
	}

	logx.Info("Update following update at success")
	return nil
}

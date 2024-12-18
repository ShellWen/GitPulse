package consumer

import (
	"context"
	"errors"
	locks "github.com/ShellWen/GitPulse/common/lock"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/developer/model"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeveloperUpdateConsumer struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewDeveloperUpdateConsumer(ctx context.Context, svc *svc.ServiceContext) *DeveloperUpdateConsumer {
	return &DeveloperUpdateConsumer{
		ctx: ctx,
		svc: svc,
	}
}

func (c *DeveloperUpdateConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	logx.Info("consume message: ", value)

	var (
		newDeveloper *model.Developer
		oldDeveloper *model.Developer
		exist        bool
	)

	if err = jsonx.UnmarshalFromString(value, &newDeveloper); err != nil {
		logx.Error("UnmarshalFromString error: ", err)
		return
	}

	if oldDeveloper, exist, err = getOldDeveloper(c, newDeveloper); err != nil {
		return
	}

	if exist {
		if err = updateOldDeveloper(c, oldDeveloper, newDeveloper); err != nil {
			return
		}
	} else {
		if err = insertNewDeveloper(c, newDeveloper); err != nil {
			return
		}
	}

	if err = unblockDeveloperUpdateLock(c, newDeveloper.Id); err != nil {
		return
	}

	logx.Info("Update developer: ", newDeveloper.Id, " success")
	return
}

func getOldDeveloper(c *DeveloperUpdateConsumer, newDeveloper *model.Developer) (*model.Developer, bool, error) {
	if oldDeveloper, err := c.svc.DeveloperModel.FindOneById(c.ctx, newDeveloper.Id); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			logx.Info("Developer not found")
			return nil, false, nil
		default:
			logx.Error("FindOneById error: ", err)
			return nil, false, err
		}
	} else {
		logx.Info("Find old developer success")
		return oldDeveloper, true, nil
	}
}

func updateOldDeveloper(c *DeveloperUpdateConsumer, oldDeveloper *model.Developer, newDeveloper *model.Developer) error {
	oldDeveloper.Name = newDeveloper.Name
	oldDeveloper.Login = newDeveloper.Login
	oldDeveloper.AvatarUrl = newDeveloper.AvatarUrl
	oldDeveloper.Company = newDeveloper.Company
	oldDeveloper.Location = newDeveloper.Location
	oldDeveloper.Bio = newDeveloper.Bio
	oldDeveloper.Blog = newDeveloper.Blog
	oldDeveloper.Email = newDeveloper.Email
	oldDeveloper.CreatedAt = newDeveloper.CreatedAt
	oldDeveloper.UpdatedAt = newDeveloper.UpdatedAt
	oldDeveloper.TwitterUsername = newDeveloper.TwitterUsername
	oldDeveloper.Repos = newDeveloper.Repos
	oldDeveloper.Stars = newDeveloper.Stars
	oldDeveloper.Gists = newDeveloper.Gists
	oldDeveloper.Followers = newDeveloper.Followers
	oldDeveloper.Following = newDeveloper.Following

	if err := c.svc.DeveloperModel.Update(c.ctx, oldDeveloper); err != nil {
		logx.Error("Update error: ", err)
		return err
	}

	logx.Info("Update developer success")
	return nil
}

func insertNewDeveloper(c *DeveloperUpdateConsumer, newDeveloper *model.Developer) error {
	if _, err := c.svc.DeveloperModel.Insert(c.ctx, newDeveloper); err != nil {
		logx.Error("Insert error: ", err)
		return err
	}

	logx.Info("Insert new developer success")
	return nil
}

func unblockDeveloperUpdateLock(c *DeveloperUpdateConsumer, developerId int64) (err error) {
	if _, err = c.svc.RedisClient.LpushCtx(c.ctx, locks.GetNewLockKey(locks.UpdateDeveloper, developerId), ""); err != nil {
		logx.Error("LpushCtx error: ", err)
		return
	}

	logx.Info("Unblock developer update lock success")
	return
}

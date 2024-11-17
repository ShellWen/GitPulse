package consumer

import (
	"context"
	"errors"
	locks "github.com/ShellWen/GitPulse/common/lock"
	"github.com/ShellWen/GitPulse/common/tasks"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/contribution/model"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type ContributionUpdateConsumer struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewContributionUpdateConsumer(ctx context.Context, svc *svc.ServiceContext) *ContributionUpdateConsumer {
	return &ContributionUpdateConsumer{
		ctx: ctx,
		svc: svc,
	}
}

func (c *ContributionUpdateConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	logx.Info("consume message: ", value)

	var (
		newContribution *model.Contribution
		oldContribution *model.Contribution
		exist           bool
	)

	if err = jsonx.UnmarshalFromString(value, &newContribution); err != nil {
		logx.Error("UnmarshalFromString error: ", err)
		return
	}

	switch newContribution.DataId {
	case tasks.FetchIssuePROfUserCompletedDataId:
		if err = updateIssuePROfUserUpdatedAt(c, newContribution.UserId); err != nil {
			return
		}
		if err = unblockIssuePROfUserUpdateLock(c, newContribution.UserId); err != nil {
			return
		}
	case tasks.FetchCommentOfUserCompletedDataId:
		if err = updateCommentOfUserUpdatedAt(c, newContribution.UserId); err != nil {
			return
		}
		if err = unblockCommentOfUserUpdateLock(c, newContribution.UserId); err != nil {
			return
		}
	case tasks.FetchReviewOfUserCompletedDataId:
		if err = updateReviewOfUserUpdatedAt(c, newContribution.UserId); err != nil {
			return
		}
		if err = unblockReviewOfUserUpdateLock(c, newContribution.UserId); err != nil {
			return
		}
	default:
		if oldContribution, exist, err = getOldContribution(c, newContribution); err != nil {
			return
		}

		if exist {
			if err = updateOldContribution(c, oldContribution, newContribution); err != nil {
				return
			}
		} else {
			if err = insertNewContribution(c, newContribution); err != nil {
				return
			}
		}
	}

	logx.Info("Update contribution: ", newContribution.ContributionId, " success")
	return
}

func getOldContribution(c *ContributionUpdateConsumer, newContribution *model.Contribution) (*model.Contribution, bool, error) {
	if oldContribution, err := c.svc.ContributionModel.FindOneByCategoryRepoIdContributionId(c.ctx, newContribution.Category, newContribution.RepoId, newContribution.ContributionId); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			return nil, false, nil
		default:
			logx.Error("FindOneById error: ", err)
			return nil, false, err
		}
	} else {
		return oldContribution, true, nil
	}
}

func updateOldContribution(c *ContributionUpdateConsumer, oldContribution *model.Contribution, newContribution *model.Contribution) error {
	oldContribution.Content = newContribution.Content
	oldContribution.UpdatedAt = newContribution.UpdatedAt
	oldContribution.CreatedAt = newContribution.CreatedAt

	if err := c.svc.ContributionModel.Update(c.ctx, oldContribution); err != nil {
		logx.Error("Update error: ", err)
		return err
	}

	return nil
}

func insertNewContribution(c *ContributionUpdateConsumer, newContribution *model.Contribution) error {
	if _, err := c.svc.ContributionModel.Insert(c.ctx, newContribution); err != nil {
		logx.Error("Insert error: ", err)
		return err
	}

	return nil
}

func unblockIssuePROfUserUpdateLock(c *ContributionUpdateConsumer, developerId int64) error {
	if _, err := c.svc.RedisClient.LpushCtx(c.ctx, locks.GetNewLockKey(locks.UpdateIssuePROfUser, developerId), ""); err != nil {
		logx.Error("LpushCtx error: ", err)
		return err
	}

	logx.Info("Unblock issuePR of user update lock success")
	return nil
}

func unblockCommentOfUserUpdateLock(c *ContributionUpdateConsumer, developerId int64) error {
	if _, err := c.svc.RedisClient.LpushCtx(c.ctx, locks.GetNewLockKey(locks.UpdateCommentOfUser, developerId), ""); err != nil {
		logx.Error("LpushCtx error: ", err)
		return err
	}

	logx.Info("Unblock comment of user update lock success")
	return nil
}

func unblockReviewOfUserUpdateLock(c *ContributionUpdateConsumer, developerId int64) error {
	if _, err := c.svc.RedisClient.LpushCtx(c.ctx, locks.GetNewLockKey(locks.UpdateReviewOfUser, developerId), ""); err != nil {
		logx.Error("LpushCtx error: ", err)
		return err
	}

	logx.Info("Unblock review of user update lock success")
	return nil
}

func updateIssuePROfUserUpdatedAt(c *ContributionUpdateConsumer, developerId int64) error {
	issuePROfUserUpdatedAt, err := c.svc.IssuePrOfUserUpdatedAtModel.FindOneByDeveloperId(c.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			issuePROfUserUpdatedAt = &model.IssuePrOfUserUpdatedAt{
				DeveloperId: developerId,
				UpdatedAt:   time.Now(),
			}
			if _, err = c.svc.IssuePrOfUserUpdatedAtModel.Insert(c.ctx, issuePROfUserUpdatedAt); err != nil {
				logx.Error("Insert error: ", err)
				return err
			}
			return nil
		default:
			logx.Error("FindOneByDeveloperId error: ", err)
			return err
		}
	}

	issuePROfUserUpdatedAt.UpdatedAt = time.Now()

	if err = c.svc.IssuePrOfUserUpdatedAtModel.Update(c.ctx, issuePROfUserUpdatedAt); err != nil {
		logx.Error("Update error: ", err)
		return err
	}

	logx.Info("Update issuePR of user update at success")
	return nil
}

func updateCommentOfUserUpdatedAt(c *ContributionUpdateConsumer, developerId int64) error {
	commentOfUserUpdatedAt, err := c.svc.CommentOfUserUpdatedAtModel.FindOneByDeveloperId(c.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			commentOfUserUpdatedAt = &model.CommentOfUserUpdatedAt{
				DeveloperId: developerId,
				UpdatedAt:   time.Now(),
			}
			if _, err = c.svc.CommentOfUserUpdatedAtModel.Insert(c.ctx, commentOfUserUpdatedAt); err != nil {
				logx.Error("Insert error: ", err)
				return err
			}
			return nil
		default:
			logx.Error("FindOneByDeveloperId error: ", err)
			return err
		}
	}

	commentOfUserUpdatedAt.UpdatedAt = time.Now()

	if err = c.svc.CommentOfUserUpdatedAtModel.Update(c.ctx, commentOfUserUpdatedAt); err != nil {
		logx.Error("Update error: ", err)
		return err
	}

	logx.Info("Update comment of user update at success")
	return nil
}

func updateReviewOfUserUpdatedAt(c *ContributionUpdateConsumer, developerId int64) error {
	reviewOfUserUpdatedAt, err := c.svc.ReviewOfUserUpdatedAtModel.FindOneByDeveloperId(c.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			reviewOfUserUpdatedAt = &model.ReviewOfUserUpdatedAt{
				DeveloperId: developerId,
				UpdatedAt:   time.Now(),
			}
			if _, err = c.svc.ReviewOfUserUpdatedAtModel.Insert(c.ctx, reviewOfUserUpdatedAt); err != nil {
				logx.Error("Insert error: ", err)
				return err
			}
			return nil
		default:
			logx.Error("FindOneByDeveloperId error: ", err)
			return err
		}
	}

	reviewOfUserUpdatedAt.UpdatedAt = time.Now()

	if err = c.svc.ReviewOfUserUpdatedAtModel.Update(c.ctx, reviewOfUserUpdatedAt); err != nil {
		logx.Error("Update error: ", err)
		return err
	}

	logx.Info("Update review of user update at success")
	return nil
}

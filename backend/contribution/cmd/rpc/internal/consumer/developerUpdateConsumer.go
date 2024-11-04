package consumer

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/contribution/model"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
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

	if oldContribution, exist, err = getOldContribution(c, newContribution); err != nil {
		return
	}

	if exist {
		err = updateOldContribution(c, oldContribution, newContribution)
	} else {
		err = insertNewContribution(c, newContribution)
	}

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

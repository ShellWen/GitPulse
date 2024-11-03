package consumer

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/repo/model"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type RepoUpdateConsumer struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewRepoUpdateConsumer(ctx context.Context, svc *svc.ServiceContext) *RepoUpdateConsumer {
	return &RepoUpdateConsumer{
		ctx: ctx,
		svc: svc,
	}
}

func (c *RepoUpdateConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	logx.Info("consume message: ", value)

	var (
		newRepo *model.Repo
		oldRepo *model.Repo
		exist   bool
	)

	if err = jsonx.UnmarshalFromString(value, &newRepo); err != nil {
		logx.Error("UnmarshalFromString error: ", err)
		return
	}

	if oldRepo, exist, err = getOldRepo(c, newRepo); err != nil {
		return
	}

	if exist {
		err = updateOldRepo(c, oldRepo, newRepo)
	} else {
		err = insertNewRepo(c, newRepo)
	}

	if c.svc.RepoUpdatedChan[newRepo.Id] == nil {
		c.svc.RepoUpdatedChan[newRepo.Id] = make(chan struct{})
	}

	select {
	case c.svc.RepoUpdatedChan[newRepo.Id] <- struct{}{}:
	default:
	}

	return
}

func getOldRepo(c *RepoUpdateConsumer, newRepo *model.Repo) (*model.Repo, bool, error) {
	if oldRepo, err := c.svc.RepoModel.FindOneById(c.ctx, newRepo.Id); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			return nil, false, nil
		default:
			logx.Error("FindOneById error: ", err)
			return nil, false, err
		}
	} else {
		return oldRepo, true, nil
	}
}

func updateOldRepo(c *RepoUpdateConsumer, oldRepo *model.Repo, newRepo *model.Repo) error {
	oldRepo.Name = newRepo.Name
	oldRepo.StarCount = newRepo.StarCount
	oldRepo.ForkCount = newRepo.ForkCount
	oldRepo.IssueCount = newRepo.IssueCount
	oldRepo.CommitCount = newRepo.CommitCount
	oldRepo.Language = newRepo.Language
	oldRepo.Description = newRepo.Description
	oldRepo.LastFetchForkAt = newRepo.LastFetchForkAt
	oldRepo.LastFetchContributionAt = newRepo.LastFetchContributionAt
	oldRepo.MergedPrCount = newRepo.MergedPrCount
	oldRepo.OpenPrCount = newRepo.OpenPrCount
	oldRepo.CommentCount = newRepo.CommentCount
	oldRepo.ReviewCount = newRepo.ReviewCount
	oldRepo.PrCount = newRepo.PrCount

	if err := c.svc.RepoModel.Update(c.ctx, oldRepo); err != nil {
		logx.Error("Update error: ", err)
		return err
	}

	return nil
}

func insertNewRepo(c *RepoUpdateConsumer, newRepo *model.Repo) error {
	if _, err := c.svc.RepoModel.Insert(c.ctx, newRepo); err != nil {
		logx.Error("Insert error: ", err)
		return err
	}

	return nil
}

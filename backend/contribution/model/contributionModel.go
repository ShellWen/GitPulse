package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContributionModel = (*customContributionModel)(nil)

type (
	// ContributionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContributionModel.
	ContributionModel interface {
		contributionModel
		SearchByCategory(ctx context.Context, category string, page int64, limit int64) (*[]*Contribution, error)
		SearchByUserId(ctx context.Context, userId int64, page int64, limit int64) (*[]*Contribution, error)
		SearchByRepoId(ctx context.Context, repoId int64, page int64, limit int64) (*[]*Contribution, error)
	}

	customContributionModel struct {
		*defaultContributionModel
	}
)

func (m *customContributionModel) SearchByCategory(ctx context.Context, category string, page int64, limit int64) (*[]*Contribution, error) {
	var resp []*Contribution

	query := fmt.Sprintf("select %s from %s where category = '%s' limit %d offset %d", contributionRows, m.table, category, limit, (page-1)*limit)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (m *customContributionModel) SearchByUserId(ctx context.Context, userId int64, page int64, limit int64) (*[]*Contribution, error) {
	var resp []*Contribution

	query := fmt.Sprintf("select %s from %s where user_id = %d limit %d offset %d", contributionRows, m.table, userId, limit, (page-1)*limit)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (m *customContributionModel) SearchByRepoId(ctx context.Context, repoId int64, page int64, limit int64) (*[]*Contribution, error) {
	var resp []*Contribution

	query := fmt.Sprintf("select %s from %s where repo_id = %d limit %d offset %d", contributionRows, m.table, repoId, limit, (page-1)*limit)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query); err != nil {
		return nil, err
	}

	return &resp, nil
}

// NewContributionModel returns a model for the database table.
func NewContributionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ContributionModel {
	return &customContributionModel{
		defaultContributionModel: newContributionModel(conn, c, opts...),
	}
}

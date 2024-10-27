package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StarModel = (*customStarModel)(nil)

type (
	// StarModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStarModel.
	StarModel interface {
		starModel
		SearchStaredRepo(ctx context.Context, developerId uint64, page int64, limit int64) (*[]*Star, error)
		SearchStaringDeveloper(ctx context.Context, repoId uint64, page int64, limit int64) (*[]*Star, error)
	}

	customStarModel struct {
		*defaultStarModel
	}
)

func (m *customStarModel) SearchStaredRepo(ctx context.Context, developerId uint64, page int64, limit int64) (*[]*Star, error) {
	var resp []*Star

	query := fmt.Sprintf("select %s from %s where `developer_id` = ? limit ?,?", starRows, m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query, developerId, (page-1)*limit, limit); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (m *customStarModel) SearchStaringDeveloper(ctx context.Context, repoId uint64, page int64, limit int64) (*[]*Star, error) {
	var resp []*Star

	query := fmt.Sprintf("select %s from %s where `repo_id` = ? limit ?,?", starRows, m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query, repoId, (page-1)*limit, limit); err != nil {
		return nil, err
	}

	return &resp, nil
}

// NewStarModel returns a model for the database table.
func NewStarModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) StarModel {
	return &customStarModel{
		defaultStarModel: newStarModel(conn, c, opts...),
	}
}

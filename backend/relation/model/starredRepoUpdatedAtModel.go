package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StarredRepoUpdatedAtModel = (*customStarredRepoUpdatedAtModel)(nil)

type (
	// StarredRepoUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStarredRepoUpdatedAtModel.
	StarredRepoUpdatedAtModel interface {
		starredRepoUpdatedAtModel
	}

	customStarredRepoUpdatedAtModel struct {
		*defaultStarredRepoUpdatedAtModel
	}
)

// NewStarredRepoUpdatedAtModel returns a model for the database table.
func NewStarredRepoUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) StarredRepoUpdatedAtModel {
	return &customStarredRepoUpdatedAtModel{
		defaultStarredRepoUpdatedAtModel: newStarredRepoUpdatedAtModel(conn, c, opts...),
	}
}

package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CreatedRepoUpdatedAtModel = (*customCreatedRepoUpdatedAtModel)(nil)

type (
	// CreatedRepoUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCreatedRepoUpdatedAtModel.
	CreatedRepoUpdatedAtModel interface {
		createdRepoUpdatedAtModel
	}

	customCreatedRepoUpdatedAtModel struct {
		*defaultCreatedRepoUpdatedAtModel
	}
)

// NewCreatedRepoUpdatedAtModel returns a model for the database table.
func NewCreatedRepoUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CreatedRepoUpdatedAtModel {
	return &customCreatedRepoUpdatedAtModel{
		defaultCreatedRepoUpdatedAtModel: newCreatedRepoUpdatedAtModel(conn, c, opts...),
	}
}

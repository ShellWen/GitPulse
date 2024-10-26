package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RepoModel = (*customRepoModel)(nil)

type (
	// RepoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRepoModel.
	RepoModel interface {
		repoModel
	}

	customRepoModel struct {
		*defaultRepoModel
	}
)

// NewRepoModel returns a model for the database table.
func NewRepoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RepoModel {
	return &customRepoModel{
		defaultRepoModel: newRepoModel(conn, c, opts...),
	}
}

package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FollowerUpdatedAtModel = (*customFollowerUpdatedAtModel)(nil)

type (
	// FollowerUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFollowerUpdatedAtModel.
	FollowerUpdatedAtModel interface {
		followerUpdatedAtModel
	}

	customFollowerUpdatedAtModel struct {
		*defaultFollowerUpdatedAtModel
	}
)

// NewFollowerUpdatedAtModel returns a model for the database table.
func NewFollowerUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) FollowerUpdatedAtModel {
	return &customFollowerUpdatedAtModel{
		defaultFollowerUpdatedAtModel: newFollowerUpdatedAtModel(conn, c, opts...),
	}
}

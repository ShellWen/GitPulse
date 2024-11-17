package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FollowingUpdatedAtModel = (*customFollowingUpdatedAtModel)(nil)

type (
	// FollowingUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFollowingUpdatedAtModel.
	FollowingUpdatedAtModel interface {
		followingUpdatedAtModel
	}

	customFollowingUpdatedAtModel struct {
		*defaultFollowingUpdatedAtModel
	}
)

// NewFollowingUpdatedAtModel returns a model for the database table.
func NewFollowingUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) FollowingUpdatedAtModel {
	return &customFollowingUpdatedAtModel{
		defaultFollowingUpdatedAtModel: newFollowingUpdatedAtModel(conn, c, opts...),
	}
}

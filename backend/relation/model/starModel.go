package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StarModel = (*customStarModel)(nil)

type (
	// StarModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStarModel.
	StarModel interface {
		starModel
	}

	customStarModel struct {
		*defaultStarModel
	}
)

// NewStarModel returns a model for the database table.
func NewStarModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) StarModel {
	return &customStarModel{
		defaultStarModel: newStarModel(conn, c, opts...),
	}
}

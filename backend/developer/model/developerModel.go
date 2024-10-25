package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DeveloperModel = (*customDeveloperModel)(nil)

type (
	// DeveloperModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeveloperModel.
	DeveloperModel interface {
		developerModel
	}

	customDeveloperModel struct {
		*defaultDeveloperModel
	}
)

// NewDeveloperModel returns a model for the database table.
func NewDeveloperModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DeveloperModel {
	return &customDeveloperModel{
		defaultDeveloperModel: newDeveloperModel(conn, c, opts...),
	}
}

package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NationModel = (*customNationModel)(nil)

type (
	// NationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNationModel.
	NationModel interface {
		nationModel
	}

	customNationModel struct {
		*defaultNationModel
	}
)

// NewNationModel returns a model for the database table.
func NewNationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) NationModel {
	return &customNationModel{
		defaultNationModel: newNationModel(conn, c, opts...),
	}
}

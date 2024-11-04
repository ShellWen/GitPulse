package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RegionModel = (*customRegionModel)(nil)

type (
	// RegionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRegionModel.
	RegionModel interface {
		regionModel
	}

	customRegionModel struct {
		*defaultRegionModel
	}
)

// NewRegionModel returns a model for the database table.
func NewRegionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RegionModel {
	return &customRegionModel{
		defaultRegionModel: newRegionModel(conn, c, opts...),
	}
}

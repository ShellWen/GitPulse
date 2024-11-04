package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PulsePointModel = (*customPulsePointModel)(nil)

type (
	// PulsePointModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPulsePointModel.
	PulsePointModel interface {
		pulsePointModel
	}

	customPulsePointModel struct {
		*defaultPulsePointModel
	}
)

// NewPulsePointModel returns a model for the database table.
func NewPulsePointModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PulsePointModel {
	return &customPulsePointModel{
		defaultPulsePointModel: newPulsePointModel(conn, c, opts...),
	}
}

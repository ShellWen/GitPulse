package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ForkUpdatedAtModel = (*customForkUpdatedAtModel)(nil)

type (
	// ForkUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customForkUpdatedAtModel.
	ForkUpdatedAtModel interface {
		forkUpdatedAtModel
	}

	customForkUpdatedAtModel struct {
		*defaultForkUpdatedAtModel
	}
)

// NewForkUpdatedAtModel returns a model for the database table.
func NewForkUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ForkUpdatedAtModel {
	return &customForkUpdatedAtModel{
		defaultForkUpdatedAtModel: newForkUpdatedAtModel(conn, c, opts...),
	}
}

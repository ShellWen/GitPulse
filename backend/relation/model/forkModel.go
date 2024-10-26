package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ForkModel = (*customForkModel)(nil)

type (
	// ForkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customForkModel.
	ForkModel interface {
		forkModel
	}

	customForkModel struct {
		*defaultForkModel
	}
)

// NewForkModel returns a model for the database table.
func NewForkModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ForkModel {
	return &customForkModel{
		defaultForkModel: newForkModel(conn, c, opts...),
	}
}

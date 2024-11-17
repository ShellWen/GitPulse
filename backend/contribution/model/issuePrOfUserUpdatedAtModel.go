package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ IssuePrOfUserUpdatedAtModel = (*customIssuePrOfUserUpdatedAtModel)(nil)

type (
	// IssuePrOfUserUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIssuePrOfUserUpdatedAtModel.
	IssuePrOfUserUpdatedAtModel interface {
		issuePrOfUserUpdatedAtModel
	}

	customIssuePrOfUserUpdatedAtModel struct {
		*defaultIssuePrOfUserUpdatedAtModel
	}
)

// NewIssuePrOfUserUpdatedAtModel returns a model for the database table.
func NewIssuePrOfUserUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) IssuePrOfUserUpdatedAtModel {
	return &customIssuePrOfUserUpdatedAtModel{
		defaultIssuePrOfUserUpdatedAtModel: newIssuePrOfUserUpdatedAtModel(conn, c, opts...),
	}
}

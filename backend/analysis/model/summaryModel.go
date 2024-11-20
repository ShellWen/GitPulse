package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SummaryModel = (*customSummaryModel)(nil)

type (
	// SummaryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSummaryModel.
	SummaryModel interface {
		summaryModel
	}

	customSummaryModel struct {
		*defaultSummaryModel
	}
)

// NewSummaryModel returns a model for the database table.
func NewSummaryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SummaryModel {
	return &customSummaryModel{
		defaultSummaryModel: newSummaryModel(conn, c, opts...),
	}
}

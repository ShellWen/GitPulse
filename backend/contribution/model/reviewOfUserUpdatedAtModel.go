package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ReviewOfUserUpdatedAtModel = (*customReviewOfUserUpdatedAtModel)(nil)

type (
	// ReviewOfUserUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customReviewOfUserUpdatedAtModel.
	ReviewOfUserUpdatedAtModel interface {
		reviewOfUserUpdatedAtModel
	}

	customReviewOfUserUpdatedAtModel struct {
		*defaultReviewOfUserUpdatedAtModel
	}
)

// NewReviewOfUserUpdatedAtModel returns a model for the database table.
func NewReviewOfUserUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ReviewOfUserUpdatedAtModel {
	return &customReviewOfUserUpdatedAtModel{
		defaultReviewOfUserUpdatedAtModel: newReviewOfUserUpdatedAtModel(conn, c, opts...),
	}
}

package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentOfUserUpdatedAtModel = (*customCommentOfUserUpdatedAtModel)(nil)

type (
	// CommentOfUserUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentOfUserUpdatedAtModel.
	CommentOfUserUpdatedAtModel interface {
		commentOfUserUpdatedAtModel
	}

	customCommentOfUserUpdatedAtModel struct {
		*defaultCommentOfUserUpdatedAtModel
	}
)

// NewCommentOfUserUpdatedAtModel returns a model for the database table.
func NewCommentOfUserUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CommentOfUserUpdatedAtModel {
	return &customCommentOfUserUpdatedAtModel{
		defaultCommentOfUserUpdatedAtModel: newCommentOfUserUpdatedAtModel(conn, c, opts...),
	}
}

package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LanguagesModel = (*customLanguagesModel)(nil)

type (
	// LanguagesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLanguagesModel.
	LanguagesModel interface {
		languagesModel
	}

	customLanguagesModel struct {
		*defaultLanguagesModel
	}
)

// NewLanguagesModel returns a model for the database table.
func NewLanguagesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LanguagesModel {
	return &customLanguagesModel{
		defaultLanguagesModel: newLanguagesModel(conn, c, opts...),
	}
}

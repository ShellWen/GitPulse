package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AnalysisModel = (*customAnalysisModel)(nil)

type (
	// AnalysisModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAnalysisModel.
	AnalysisModel interface {
		analysisModel
	}

	customAnalysisModel struct {
		*defaultAnalysisModel
	}
)

// NewAnalysisModel returns a model for the database table.
func NewAnalysisModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AnalysisModel {
	return &customAnalysisModel{
		defaultAnalysisModel: newAnalysisModel(conn, c, opts...),
	}
}

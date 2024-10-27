package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ForkModel = (*customForkModel)(nil)

type (
	// ForkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customForkModel.
	ForkModel interface {
		forkModel
		SearchFork(ctx context.Context, originalRepoId int64, page int64, limit int64) (*[]*Fork, error)
	}

	customForkModel struct {
		*defaultForkModel
	}
)

func (m *customForkModel) SearchFork(ctx context.Context, originalRepoId int64, page int64, limit int64) (*[]*Fork, error) {
	var resp []*Fork

	query := fmt.Sprintf("select %s from %s where original_repo_id = %d limit %d offset %d", forkRows, m.table, originalRepoId, limit, (page-1)*limit)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query); err != nil {
		return nil, err
	}

	return &resp, nil
}

// NewForkModel returns a model for the database table.
func NewForkModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ForkModel {
	return &customForkModel{
		defaultForkModel: newForkModel(conn, c, opts...),
	}
}

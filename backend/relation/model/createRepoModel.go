package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CreateRepoModel = (*customCreateRepoModel)(nil)

type (
	// CreateRepoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCreateRepoModel.
	CreateRepoModel interface {
		createRepoModel
		SearchCreatedRepo(ctx context.Context, developerId uint64, page int64, limit int64) (*[]*CreateRepo, error)
	}

	customCreateRepoModel struct {
		*defaultCreateRepoModel
	}
)

func (m *customCreateRepoModel) SearchCreatedRepo(ctx context.Context, developerId uint64, page int64, limit int64) (*[]*CreateRepo, error) {
	var resp []*CreateRepo

	query := fmt.Sprintf("select %s from %s where `developer_id` = ? limit ?,?", createRepoRows, m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query, developerId, (page-1)*limit, limit); err != nil {
		return nil, err
	}

	return &resp, nil
}

// NewCreateRepoModel returns a model for the database table.
func NewCreateRepoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CreateRepoModel {
	return &customCreateRepoModel{
		defaultCreateRepoModel: newCreateRepoModel(conn, c, opts...),
	}
}

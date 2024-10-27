package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FollowModel = (*customFollowModel)(nil)

type (
	// FollowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFollowModel.
	FollowModel interface {
		followModel
		SearchFollowed(ctx context.Context, following uint64, page int64, limit int64) (*[]*Follow, error)
		SearchFollowing(ctx context.Context, followed uint64, page int64, limit int64) (*[]*Follow, error)
	}

	customFollowModel struct {
		*defaultFollowModel
	}
)

func (m *customFollowModel) SearchFollowed(ctx context.Context, following uint64, page int64, limit int64) (*[]*Follow, error) {
	var resp []*Follow

	query := fmt.Sprintf("select %s from %s where `following` = ? limit ?,?", followRows, m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query, following, (page-1)*limit, limit); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (m *customFollowModel) SearchFollowing(ctx context.Context, followed uint64, page int64, limit int64) (*[]*Follow, error) {
	var resp []*Follow

	query := fmt.Sprintf("select %s from %s where `followed` = ? limit ?,?", followRows, m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query, followed, (page-1)*limit, limit); err != nil {
		return nil, err
	}

	return &resp, nil
}

// NewFollowModel returns a model for the database table.
func NewFollowModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) FollowModel {
	return &customFollowModel{
		defaultFollowModel: newFollowModel(conn, c, opts...),
	}
}

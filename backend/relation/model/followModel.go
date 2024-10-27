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
		SearchFollowed(ctx context.Context, following int64, page int64, limit int64) (*[]*Follow, error)
		SearchFollowing(ctx context.Context, followed int64, page int64, limit int64) (*[]*Follow, error)
	}

	customFollowModel struct {
		*defaultFollowModel
	}
)

func (m *customFollowModel) SearchFollowed(ctx context.Context, following int64, page int64, limit int64) (*[]*Follow, error) {
	var resp []*Follow

	query := fmt.Sprintf("select %s from %s where following_id = %d limit %d offset %d", followRows, m.table, following, limit, (page-1)*limit)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (m *customFollowModel) SearchFollowing(ctx context.Context, followed int64, page int64, limit int64) (*[]*Follow, error) {
	var resp []*Follow

	query := fmt.Sprintf("select %s from %s where followed_id = %d limit %d offset %d", followRows, m.table, followed, limit, (page-1)*limit)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query); err != nil {
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

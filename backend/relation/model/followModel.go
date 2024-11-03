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
		SearchFollowingByDeveloperId(ctx context.Context, developerId int64, page int64, limit int64) (*[]*Follow, error)
		SearchFollowerByDeveloperId(ctx context.Context, developerId int64, page int64, limit int64) (*[]*Follow, error)
	}

	customFollowModel struct {
		*defaultFollowModel
	}
)

func (m *customFollowModel) SearchFollowingByDeveloperId(ctx context.Context, developerId int64, page int64, limit int64) (*[]*Follow, error) {
	var resp []*Follow

	query := fmt.Sprintf("select %s from %s where follower_id = %d limit %d offset %d", followRows, m.table, developerId, limit, (page-1)*limit)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (m *customFollowModel) SearchFollowerByDeveloperId(ctx context.Context, developerId int64, page int64, limit int64) (*[]*Follow, error) {
	var resp []*Follow

	query := fmt.Sprintf("select %s from %s where following_id = %d limit %d offset %d", followRows, m.table, developerId, limit, (page-1)*limit)
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

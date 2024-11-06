package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PulsePointModel = (*customPulsePointModel)(nil)

type (
	// PulsePointModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPulsePointModel.
	PulsePointModel interface {
		pulsePointModel
		FindAll(ctx context.Context) ([]*PulsePoint, error)
	}

	customPulsePointModel struct {
		*defaultPulsePointModel
	}
)

func (m *customPulsePointModel) FindAll(ctx context.Context) ([]*PulsePoint, error) {
	var resp []*PulsePoint

	query := fmt.Sprintf("select %s from %s", pulsePointRows, m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query); err != nil {
		return nil, err
	}

	return resp, nil
}

// NewPulsePointModel returns a model for the database table.
func NewPulsePointModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PulsePointModel {
	return &customPulsePointModel{
		defaultPulsePointModel: newPulsePointModel(conn, c, opts...),
	}
}

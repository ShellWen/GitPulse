// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	followerUpdatedAtFieldNames          = builder.RawFieldNames(&FollowerUpdatedAt{}, true)
	followerUpdatedAtRows                = strings.Join(followerUpdatedAtFieldNames, ",")
	followerUpdatedAtRowsExpectAutoSet   = strings.Join(stringx.Remove(followerUpdatedAtFieldNames, "data_id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	followerUpdatedAtRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(followerUpdatedAtFieldNames, "data_id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cacheRelationFollowerUpdatedAtDataIdPrefix      = "cache:relation:followerUpdatedAt:dataId:"
	cacheRelationFollowerUpdatedAtDeveloperIdPrefix = "cache:relation:followerUpdatedAt:developerId:"
)

type (
	followerUpdatedAtModel interface {
		Insert(ctx context.Context, data *FollowerUpdatedAt) (sql.Result, error)
		FindOne(ctx context.Context, dataId int64) (*FollowerUpdatedAt, error)
		FindOneByDeveloperId(ctx context.Context, developerId int64) (*FollowerUpdatedAt, error)
		Update(ctx context.Context, data *FollowerUpdatedAt) error
		Delete(ctx context.Context, dataId int64) error
	}

	defaultFollowerUpdatedAtModel struct {
		sqlc.CachedConn
		table string
	}

	FollowerUpdatedAt struct {
		UpdatedAt   time.Time `db:"updated_at"`
		DataId      int64     `db:"data_id"`
		DeveloperId int64     `db:"developer_id"`
	}
)

func newFollowerUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultFollowerUpdatedAtModel {
	return &defaultFollowerUpdatedAtModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"relation"."follower_updated_at"`,
	}
}

func (m *defaultFollowerUpdatedAtModel) Delete(ctx context.Context, dataId int64) error {
	data, err := m.FindOne(ctx, dataId)
	if err != nil {
		return err
	}

	relationFollowerUpdatedAtDataIdKey := fmt.Sprintf("%s%v", cacheRelationFollowerUpdatedAtDataIdPrefix, dataId)
	relationFollowerUpdatedAtDeveloperIdKey := fmt.Sprintf("%s%v", cacheRelationFollowerUpdatedAtDeveloperIdPrefix, data.DeveloperId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where data_id = $1", m.table)
		return conn.ExecCtx(ctx, query, dataId)
	}, relationFollowerUpdatedAtDataIdKey, relationFollowerUpdatedAtDeveloperIdKey)
	return err
}

func (m *defaultFollowerUpdatedAtModel) FindOne(ctx context.Context, dataId int64) (*FollowerUpdatedAt, error) {
	relationFollowerUpdatedAtDataIdKey := fmt.Sprintf("%s%v", cacheRelationFollowerUpdatedAtDataIdPrefix, dataId)
	var resp FollowerUpdatedAt
	err := m.QueryRowCtx(ctx, &resp, relationFollowerUpdatedAtDataIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where data_id = $1 limit 1", followerUpdatedAtRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, dataId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFollowerUpdatedAtModel) FindOneByDeveloperId(ctx context.Context, developerId int64) (*FollowerUpdatedAt, error) {
	relationFollowerUpdatedAtDeveloperIdKey := fmt.Sprintf("%s%v", cacheRelationFollowerUpdatedAtDeveloperIdPrefix, developerId)
	var resp FollowerUpdatedAt
	err := m.QueryRowIndexCtx(ctx, &resp, relationFollowerUpdatedAtDeveloperIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where developer_id = $1 limit 1", followerUpdatedAtRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, developerId); err != nil {
			return nil, err
		}
		return resp.DataId, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFollowerUpdatedAtModel) Insert(ctx context.Context, data *FollowerUpdatedAt) (sql.Result, error) {
	relationFollowerUpdatedAtDataIdKey := fmt.Sprintf("%s%v", cacheRelationFollowerUpdatedAtDataIdPrefix, data.DataId)
	relationFollowerUpdatedAtDeveloperIdKey := fmt.Sprintf("%s%v", cacheRelationFollowerUpdatedAtDeveloperIdPrefix, data.DeveloperId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1)", m.table, followerUpdatedAtRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeveloperId)
	}, relationFollowerUpdatedAtDataIdKey, relationFollowerUpdatedAtDeveloperIdKey)
	return ret, err
}

func (m *defaultFollowerUpdatedAtModel) Update(ctx context.Context, newData *FollowerUpdatedAt) error {
	data, err := m.FindOne(ctx, newData.DataId)
	if err != nil {
		return err
	}

	relationFollowerUpdatedAtDataIdKey := fmt.Sprintf("%s%v", cacheRelationFollowerUpdatedAtDataIdPrefix, data.DataId)
	relationFollowerUpdatedAtDeveloperIdKey := fmt.Sprintf("%s%v", cacheRelationFollowerUpdatedAtDeveloperIdPrefix, data.DeveloperId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where data_id = $1", m.table, followerUpdatedAtRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.DataId, newData.DeveloperId)
	}, relationFollowerUpdatedAtDataIdKey, relationFollowerUpdatedAtDeveloperIdKey)
	return err
}

func (m *defaultFollowerUpdatedAtModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheRelationFollowerUpdatedAtDataIdPrefix, primary)
}

func (m *defaultFollowerUpdatedAtModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where data_id = $1 limit 1", followerUpdatedAtRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFollowerUpdatedAtModel) tableName() string {
	return m.table
}
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
	createdRepoUpdatedAtFieldNames          = builder.RawFieldNames(&CreatedRepoUpdatedAt{}, true)
	createdRepoUpdatedAtRows                = strings.Join(createdRepoUpdatedAtFieldNames, ",")
	createdRepoUpdatedAtRowsExpectAutoSet   = strings.Join(stringx.Remove(createdRepoUpdatedAtFieldNames, "data_id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	createdRepoUpdatedAtRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(createdRepoUpdatedAtFieldNames, "data_id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cacheRelationCreatedRepoUpdatedAtDataIdPrefix      = "cache:relation:createdRepoUpdatedAt:dataId:"
	cacheRelationCreatedRepoUpdatedAtDeveloperIdPrefix = "cache:relation:createdRepoUpdatedAt:developerId:"
)

type (
	createdRepoUpdatedAtModel interface {
		Insert(ctx context.Context, data *CreatedRepoUpdatedAt) (sql.Result, error)
		FindOne(ctx context.Context, dataId int64) (*CreatedRepoUpdatedAt, error)
		FindOneByDeveloperId(ctx context.Context, developerId int64) (*CreatedRepoUpdatedAt, error)
		Update(ctx context.Context, data *CreatedRepoUpdatedAt) error
		Delete(ctx context.Context, dataId int64) error
	}

	defaultCreatedRepoUpdatedAtModel struct {
		sqlc.CachedConn
		table string
	}

	CreatedRepoUpdatedAt struct {
		DataId      int64     `db:"data_id"`
		DeveloperId int64     `db:"developer_id"`
		UpdatedAt   time.Time `db:"updated_at"`
	}
)

func newCreatedRepoUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultCreatedRepoUpdatedAtModel {
	return &defaultCreatedRepoUpdatedAtModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"relation"."created_repo_updated_at"`,
	}
}

func (m *defaultCreatedRepoUpdatedAtModel) Delete(ctx context.Context, dataId int64) error {
	data, err := m.FindOne(ctx, dataId)
	if err != nil {
		return err
	}

	relationCreatedRepoUpdatedAtDataIdKey := fmt.Sprintf("%s%v", cacheRelationCreatedRepoUpdatedAtDataIdPrefix, dataId)
	relationCreatedRepoUpdatedAtDeveloperIdKey := fmt.Sprintf("%s%v", cacheRelationCreatedRepoUpdatedAtDeveloperIdPrefix, data.DeveloperId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where data_id = $1", m.table)
		return conn.ExecCtx(ctx, query, dataId)
	}, relationCreatedRepoUpdatedAtDataIdKey, relationCreatedRepoUpdatedAtDeveloperIdKey)
	return err
}

func (m *defaultCreatedRepoUpdatedAtModel) FindOne(ctx context.Context, dataId int64) (*CreatedRepoUpdatedAt, error) {
	relationCreatedRepoUpdatedAtDataIdKey := fmt.Sprintf("%s%v", cacheRelationCreatedRepoUpdatedAtDataIdPrefix, dataId)
	var resp CreatedRepoUpdatedAt
	err := m.QueryRowCtx(ctx, &resp, relationCreatedRepoUpdatedAtDataIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where data_id = $1 limit 1", createdRepoUpdatedAtRows, m.table)
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

func (m *defaultCreatedRepoUpdatedAtModel) FindOneByDeveloperId(ctx context.Context, developerId int64) (*CreatedRepoUpdatedAt, error) {
	relationCreatedRepoUpdatedAtDeveloperIdKey := fmt.Sprintf("%s%v", cacheRelationCreatedRepoUpdatedAtDeveloperIdPrefix, developerId)
	var resp CreatedRepoUpdatedAt
	err := m.QueryRowIndexCtx(ctx, &resp, relationCreatedRepoUpdatedAtDeveloperIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where developer_id = $1 limit 1", createdRepoUpdatedAtRows, m.table)
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

func (m *defaultCreatedRepoUpdatedAtModel) Insert(ctx context.Context, data *CreatedRepoUpdatedAt) (sql.Result, error) {
	relationCreatedRepoUpdatedAtDataIdKey := fmt.Sprintf("%s%v", cacheRelationCreatedRepoUpdatedAtDataIdPrefix, data.DataId)
	relationCreatedRepoUpdatedAtDeveloperIdKey := fmt.Sprintf("%s%v", cacheRelationCreatedRepoUpdatedAtDeveloperIdPrefix, data.DeveloperId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1)", m.table, createdRepoUpdatedAtRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeveloperId)
	}, relationCreatedRepoUpdatedAtDataIdKey, relationCreatedRepoUpdatedAtDeveloperIdKey)
	return ret, err
}

func (m *defaultCreatedRepoUpdatedAtModel) Update(ctx context.Context, newData *CreatedRepoUpdatedAt) error {
	data, err := m.FindOne(ctx, newData.DataId)
	if err != nil {
		return err
	}

	relationCreatedRepoUpdatedAtDataIdKey := fmt.Sprintf("%s%v", cacheRelationCreatedRepoUpdatedAtDataIdPrefix, data.DataId)
	relationCreatedRepoUpdatedAtDeveloperIdKey := fmt.Sprintf("%s%v", cacheRelationCreatedRepoUpdatedAtDeveloperIdPrefix, data.DeveloperId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where data_id = $1", m.table, createdRepoUpdatedAtRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.DataId, newData.DeveloperId)
	}, relationCreatedRepoUpdatedAtDataIdKey, relationCreatedRepoUpdatedAtDeveloperIdKey)
	return err
}

func (m *defaultCreatedRepoUpdatedAtModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheRelationCreatedRepoUpdatedAtDataIdPrefix, primary)
}

func (m *defaultCreatedRepoUpdatedAtModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where data_id = $1 limit 1", createdRepoUpdatedAtRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultCreatedRepoUpdatedAtModel) tableName() string {
	return m.table
}

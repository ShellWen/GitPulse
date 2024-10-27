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
	starFieldNames          = builder.RawFieldNames(&Star{}, true)
	starRows                = strings.Join(starFieldNames, ",")
	starRowsExpectAutoSet   = strings.Join(stringx.Remove(starFieldNames, "data_id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	starRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(starFieldNames, "data_id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cacheRelationStarDataIdPrefix            = "cache:relation:star:dataId:"
	cacheRelationStarDeveloperIdRepoIdPrefix = "cache:relation:star:developerId:repoId:"
)

type (
	starModel interface {
		Insert(ctx context.Context, data *Star) (sql.Result, error)
		FindOne(ctx context.Context, dataId int64) (*Star, error)
		FindOneByDeveloperIdRepoId(ctx context.Context, developerId int64, repoId int64) (*Star, error)
		Update(ctx context.Context, data *Star) error
		Delete(ctx context.Context, dataId int64) error
	}

	defaultStarModel struct {
		sqlc.CachedConn
		table string
	}

	Star struct {
		DataId       int64     `db:"data_id"`
		DataCreateAt time.Time `db:"data_create_at"`
		DataUpdateAt time.Time `db:"data_update_at"`
		DeveloperId  int64     `db:"developer_id"`
		RepoId       int64     `db:"repo_id"`
	}
)

func newStarModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultStarModel {
	return &defaultStarModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"relation"."star"`,
	}
}

func (m *defaultStarModel) Delete(ctx context.Context, dataId int64) error {
	data, err := m.FindOne(ctx, dataId)
	if err != nil {
		return err
	}

	relationStarDataIdKey := fmt.Sprintf("%s%v", cacheRelationStarDataIdPrefix, dataId)
	relationStarDeveloperIdRepoIdKey := fmt.Sprintf("%s%v:%v", cacheRelationStarDeveloperIdRepoIdPrefix, data.DeveloperId, data.RepoId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where data_id = $1", m.table)
		return conn.ExecCtx(ctx, query, dataId)
	}, relationStarDataIdKey, relationStarDeveloperIdRepoIdKey)
	return err
}

func (m *defaultStarModel) FindOne(ctx context.Context, dataId int64) (*Star, error) {
	relationStarDataIdKey := fmt.Sprintf("%s%v", cacheRelationStarDataIdPrefix, dataId)
	var resp Star
	err := m.QueryRowCtx(ctx, &resp, relationStarDataIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where data_id = $1 limit 1", starRows, m.table)
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

func (m *defaultStarModel) FindOneByDeveloperIdRepoId(ctx context.Context, developerId int64, repoId int64) (*Star, error) {
	relationStarDeveloperIdRepoIdKey := fmt.Sprintf("%s%v:%v", cacheRelationStarDeveloperIdRepoIdPrefix, developerId, repoId)
	var resp Star
	err := m.QueryRowIndexCtx(ctx, &resp, relationStarDeveloperIdRepoIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where developer_id = $1 and repo_id = $2 limit 1", starRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, developerId, repoId); err != nil {
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

func (m *defaultStarModel) Insert(ctx context.Context, data *Star) (sql.Result, error) {
	relationStarDataIdKey := fmt.Sprintf("%s%v", cacheRelationStarDataIdPrefix, data.DataId)
	relationStarDeveloperIdRepoIdKey := fmt.Sprintf("%s%v:%v", cacheRelationStarDeveloperIdRepoIdPrefix, data.DeveloperId, data.RepoId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4)", m.table, starRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DataCreateAt, data.DataUpdateAt, data.DeveloperId, data.RepoId)
	}, relationStarDataIdKey, relationStarDeveloperIdRepoIdKey)
	return ret, err
}

func (m *defaultStarModel) Update(ctx context.Context, newData *Star) error {
	data, err := m.FindOne(ctx, newData.DataId)
	if err != nil {
		return err
	}

	relationStarDataIdKey := fmt.Sprintf("%s%v", cacheRelationStarDataIdPrefix, data.DataId)
	relationStarDeveloperIdRepoIdKey := fmt.Sprintf("%s%v:%v", cacheRelationStarDeveloperIdRepoIdPrefix, data.DeveloperId, data.RepoId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where data_id = $1", m.table, starRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.DataId, newData.DataCreateAt, newData.DataUpdateAt, newData.DeveloperId, newData.RepoId)
	}, relationStarDataIdKey, relationStarDeveloperIdRepoIdKey)
	return err
}

func (m *defaultStarModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheRelationStarDataIdPrefix, primary)
}

func (m *defaultStarModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where data_id = $1 limit 1", starRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultStarModel) tableName() string {
	return m.table
}

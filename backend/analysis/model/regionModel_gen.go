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
	regionFieldNames          = builder.RawFieldNames(&Region{}, true)
	regionRows                = strings.Join(regionFieldNames, ",")
	regionRowsExpectAutoSet   = strings.Join(stringx.Remove(regionFieldNames, "data_id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	regionRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(regionFieldNames, "data_id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cacheAnalysisRegionDataIdPrefix      = "cache:analysis:region:dataId:"
	cacheAnalysisRegionDeveloperIdPrefix = "cache:analysis:region:developerId:"
)

type (
	regionModel interface {
		Insert(ctx context.Context, data *Region) (sql.Result, error)
		FindOne(ctx context.Context, dataId int64) (*Region, error)
		FindOneByDeveloperId(ctx context.Context, developerId int64) (*Region, error)
		Update(ctx context.Context, data *Region) error
		Delete(ctx context.Context, dataId int64) error
	}

	defaultRegionModel struct {
		sqlc.CachedConn
		table string
	}

	Region struct {
		DataId        int64     `db:"data_id"`
		DataCreatedAt time.Time `db:"data_created_at"`
		DataUpdatedAt time.Time `db:"data_updated_at"`
		DeveloperId   int64     `db:"developer_id"`
		Region        string    `db:"region"`
		Confidence    float64   `db:"confidence"`
	}
)

func newRegionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultRegionModel {
	return &defaultRegionModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"analysis"."region"`,
	}
}

func (m *defaultRegionModel) Delete(ctx context.Context, dataId int64) error {
	data, err := m.FindOne(ctx, dataId)
	if err != nil {
		return err
	}

	analysisRegionDataIdKey := fmt.Sprintf("%s%v", cacheAnalysisRegionDataIdPrefix, dataId)
	analysisRegionDeveloperIdKey := fmt.Sprintf("%s%v", cacheAnalysisRegionDeveloperIdPrefix, data.DeveloperId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where data_id = $1", m.table)
		return conn.ExecCtx(ctx, query, dataId)
	}, analysisRegionDataIdKey, analysisRegionDeveloperIdKey)
	return err
}

func (m *defaultRegionModel) FindOne(ctx context.Context, dataId int64) (*Region, error) {
	analysisRegionDataIdKey := fmt.Sprintf("%s%v", cacheAnalysisRegionDataIdPrefix, dataId)
	var resp Region
	err := m.QueryRowCtx(ctx, &resp, analysisRegionDataIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where data_id = $1 limit 1", regionRows, m.table)
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

func (m *defaultRegionModel) FindOneByDeveloperId(ctx context.Context, developerId int64) (*Region, error) {
	analysisRegionDeveloperIdKey := fmt.Sprintf("%s%v", cacheAnalysisRegionDeveloperIdPrefix, developerId)
	var resp Region
	err := m.QueryRowIndexCtx(ctx, &resp, analysisRegionDeveloperIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where developer_id = $1 limit 1", regionRows, m.table)
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

func (m *defaultRegionModel) Insert(ctx context.Context, data *Region) (sql.Result, error) {
	analysisRegionDataIdKey := fmt.Sprintf("%s%v", cacheAnalysisRegionDataIdPrefix, data.DataId)
	analysisRegionDeveloperIdKey := fmt.Sprintf("%s%v", cacheAnalysisRegionDeveloperIdPrefix, data.DeveloperId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5)", m.table, regionRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DataCreatedAt, data.DataUpdatedAt, data.DeveloperId, data.Region, data.Confidence)
	}, analysisRegionDataIdKey, analysisRegionDeveloperIdKey)
	return ret, err
}

func (m *defaultRegionModel) Update(ctx context.Context, newData *Region) error {
	data, err := m.FindOne(ctx, newData.DataId)
	if err != nil {
		return err
	}

	analysisRegionDataIdKey := fmt.Sprintf("%s%v", cacheAnalysisRegionDataIdPrefix, data.DataId)
	analysisRegionDeveloperIdKey := fmt.Sprintf("%s%v", cacheAnalysisRegionDeveloperIdPrefix, data.DeveloperId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where data_id = $1", m.table, regionRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.DataId, newData.DataCreatedAt, newData.DataUpdatedAt, newData.DeveloperId, newData.Region, newData.Confidence)
	}, analysisRegionDataIdKey, analysisRegionDeveloperIdKey)
	return err
}

func (m *defaultRegionModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheAnalysisRegionDataIdPrefix, primary)
}

func (m *defaultRegionModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where data_id = $1 limit 1", regionRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRegionModel) tableName() string {
	return m.table
}

package svc

import (
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/internal/config"
	"github.com/ShellWen/GitPulse/repo/model"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	RedisClient     *redis.Redis
	RepoModel       model.RepoModel
	RepoUpdatedChan map[int64]chan struct{}
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		RepoModel:       model.NewRepoModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		RepoUpdatedChan: make(map[int64]chan struct{}),
	}
}

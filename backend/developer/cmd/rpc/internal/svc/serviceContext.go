package svc

import (
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/internal/config"
	"github.com/ShellWen/GitPulse/developer/model"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	RedisClient    *redis.Redis
	DeveloperModel model.DeveloperModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		DeveloperModel: model.NewDeveloperModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
	}
}

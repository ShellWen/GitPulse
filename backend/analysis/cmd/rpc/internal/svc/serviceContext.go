package svc

import (
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/config"
	"github.com/ShellWen/GitPulse/analysis/model"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	RedisClient   *redis.Redis
	AnalysisModel model.AnalysisModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		AnalysisModel: model.NewAnalysisModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
	}
}

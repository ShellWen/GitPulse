package svc

import (
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/config"
	"github.com/ShellWen/GitPulse/analysis/model"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config             config.Config
	RedisClient        *redis.Redis
	RegionModel        model.RegionModel
	LanguagesModel     model.LanguagesModel
	PulsePointModel    model.PulsePointModel
	RpcClient          zrpc.Client
	DeveloperRpcClient zrpc.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		RegionModel:     model.NewRegionModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		LanguagesModel:  model.NewLanguagesModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		PulsePointModel: model.NewPulsePointModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		RpcClient:       zrpc.MustNewClient(c.RpcClientConf),
	}
}

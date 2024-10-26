package svc

import (
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/config"
	"github.com/ShellWen/GitPulse/relation/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	RedisClient     *redis.Redis
	CreateRepoModel model.CreateRepoModel
	FollowModel     model.FollowModel
	ForkModel       model.ForkModel
	StarModel       model.StarModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		CreateRepoModel: model.NewCreateRepoModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		FollowModel:     model.NewFollowModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		ForkModel:       model.NewForkModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		StarModel:       model.NewStarModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}

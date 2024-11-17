package svc

import (
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/config"
	"github.com/ShellWen/GitPulse/relation/model"
	"github.com/hashicorp/consul/api"
	"github.com/hibiken/asynq"
	_ "github.com/jackc/pgx/v5/stdlib"
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

	CreatedRepoUpdatedAtModel model.CreatedRepoUpdatedAtModel
	FollowingUpdatedAtModel   model.FollowingUpdatedAtModel
	FollowerUpdatedAtModel    model.FollowerUpdatedAtModel
	ForkUpdatedAtModel        model.ForkUpdatedAtModel
	StarredRepoUpdatedAtModel model.StarredRepoUpdatedAtModel

	AsynqClient  *asynq.Client
	ConsulClient *api.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		CreateRepoModel: model.NewCreateRepoModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		FollowModel:     model.NewFollowModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		ForkModel:       model.NewForkModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		StarModel:       model.NewStarModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),

		CreatedRepoUpdatedAtModel: model.NewCreatedRepoUpdatedAtModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		FollowingUpdatedAtModel:   model.NewFollowingUpdatedAtModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		FollowerUpdatedAtModel:    model.NewFollowerUpdatedAtModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		ForkUpdatedAtModel:        model.NewForkUpdatedAtModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		StarredRepoUpdatedAtModel: model.NewStarredRepoUpdatedAtModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),

		RedisClient: redis.MustNewRedis(c.Redis.RedisConf),
		AsynqClient: asynq.NewClient(asynq.RedisClientOpt{
			Addr:     c.AsynqRedisConf.Addr,
			Password: c.AsynqRedisConf.Password,
			DB:       c.AsynqRedisConf.DB,
		}),
		ConsulClient: func() *api.Client {
			client, err := api.NewClient(&api.Config{
				Address: c.Consul.Host,
			})

			if err != nil {
				panic(err)
			}

			return client
		}(),
	}
}

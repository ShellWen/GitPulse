package svc

import (
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/config"
	"github.com/ShellWen/GitPulse/contribution/model"
	"github.com/hashicorp/consul/api"
	"github.com/hibiken/asynq"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	RedisClient       *redis.Redis
	ContributionModel model.ContributionModel

	IssuePrOfUserUpdatedAtModel model.IssuePrOfUserUpdatedAtModel
	ReviewOfUserUpdatedAtModel  model.ReviewOfUserUpdatedAtModel
	CommentOfUserUpdatedAtModel model.CommentOfUserUpdatedAtModel

	AsynqClient  *asynq.Client
	ConsulClient *api.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		ContributionModel: model.NewContributionModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),

		IssuePrOfUserUpdatedAtModel: model.NewIssuePrOfUserUpdatedAtModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		ReviewOfUserUpdatedAtModel:  model.NewReviewOfUserUpdatedAtModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		CommentOfUserUpdatedAtModel: model.NewCommentOfUserUpdatedAtModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),

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

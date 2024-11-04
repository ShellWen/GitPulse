package svc

import (
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/config"
	"github.com/ShellWen/GitPulse/contribution/model"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                   config.Config
	RedisClient              *redis.Redis
	ContributionModel        model.ContributionModel
	IssuePrUpdatedChan       map[int64]chan struct{}
	CommentReviewUpdatedChan map[int64]chan struct{}
	AllUpdatedChan           map[int64]chan struct{}
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                   c,
		ContributionModel:        model.NewContributionModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		IssuePrUpdatedChan:       make(map[int64]chan struct{}),
		CommentReviewUpdatedChan: make(map[int64]chan struct{}),
		AllUpdatedChan:           make(map[int64]chan struct{}),
	}
}

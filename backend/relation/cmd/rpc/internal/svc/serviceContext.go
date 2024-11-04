package svc

import (
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/config"
	"github.com/ShellWen/GitPulse/relation/model"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                 config.Config
	RedisClient            *redis.Redis
	CreateRepoModel        model.CreateRepoModel
	FollowModel            model.FollowModel
	ForkModel              model.ForkModel
	StarModel              model.StarModel
	CreatedRepoUpdatedChan map[int64]chan struct{}
	StarredRepoUpdatedChan map[int64]chan struct{}
	ForkUpdatedChan        map[int64]chan struct{}
	FollowingUpdatedChan   map[int64]chan struct{}
	FollowerUpdatedChan    map[int64]chan struct{}
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                 c,
		CreateRepoModel:        model.NewCreateRepoModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		FollowModel:            model.NewFollowModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		ForkModel:              model.NewForkModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		StarModel:              model.NewStarModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		CreatedRepoUpdatedChan: make(map[int64]chan struct{}),
		StarredRepoUpdatedChan: make(map[int64]chan struct{}),
		ForkUpdatedChan:        make(map[int64]chan struct{}),
		FollowingUpdatedChan:   make(map[int64]chan struct{}),
		FollowerUpdatedChan:    make(map[int64]chan struct{}),
	}
}

package svc

import (
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/config"
	"github.com/ShellWen/GitPulse/analysis/model"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/contribution"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/relation"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/repo"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	RedisClient     *redis.Redis
	RegionModel     model.RegionModel
	LanguagesModel  model.LanguagesModel
	PulsePointModel model.PulsePointModel

	DeveloperRpcClient    developer.DeveloperZrpcClient
	RepoRpcClient         repo.RepoZrpcClient
	ContributionRpcClient contribution.ContributionZrpcClient
	RelationRpcClient     relation.Relation
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		RedisClient:     redis.MustNewRedis(c.Redis),
		RegionModel:     model.NewRegionModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		LanguagesModel:  model.NewLanguagesModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),
		PulsePointModel: model.NewPulsePointModel(sqlx.NewSqlConn("pgx", c.DB.DataSource), c.Cache),

		DeveloperRpcClient:    developer.NewDeveloperZrpcClient(zrpc.MustNewClient(c.DeveloperRpcConf)),
		RepoRpcClient:         repo.NewRepoZrpcClient(zrpc.MustNewClient(c.RepoRpcConf)),
		ContributionRpcClient: contribution.NewContributionZrpcClient(zrpc.MustNewClient(c.ContributionRpcConf)),
		RelationRpcClient:     relation.NewRelation(zrpc.MustNewClient(c.RelationRpcConf)),
	}
}

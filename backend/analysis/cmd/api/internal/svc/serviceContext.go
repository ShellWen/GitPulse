package svc

import (
	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/config"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/analysis"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/contribution"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/relation"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/repo"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	AsynqClient *asynq.Client
	RedisClient *redis.Redis

	DeveloperRpcClient    developer.DeveloperZrpcClient
	RepoRpcClient         repo.RepoZrpcClient
	ContributionRpcClient contribution.ContributionZrpcClient
	RelationRpcClient     relation.Relation
	AnalysisRpcClient     analysis.Analysis

	LanguageUpdating   bool
	PulsePointUpdating bool
	RegionUpdating     bool

	LanguagesUpdatedChan  map[int64]chan struct{}
	PulsePointUpdatedChan map[int64]chan struct{}
	RegionUpdatedChan     map[int64]chan struct{}
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		AsynqClient: asynq.NewClient(&asynq.RedisClientOpt{
			Addr:     c.AsynqRedisConf.Addr,
			Password: c.AsynqRedisConf.Password,
			DB:       c.AsynqRedisConf.DB,
		}),
		RedisClient: redis.MustNewRedis(c.Redis),

		LanguagesUpdatedChan:  make(map[int64]chan struct{}),
		PulsePointUpdatedChan: make(map[int64]chan struct{}),
		RegionUpdatedChan:     make(map[int64]chan struct{}),

		LanguageUpdating:   false,
		PulsePointUpdating: false,
		RegionUpdating:     false,

		DeveloperRpcClient:    developer.NewDeveloperZrpcClient(zrpc.MustNewClient(c.DeveloperRpcConf)),
		RepoRpcClient:         repo.NewRepoZrpcClient(zrpc.MustNewClient(c.RepoRpcConf)),
		ContributionRpcClient: contribution.NewContributionZrpcClient(zrpc.MustNewClient(c.ContributionRpcConf)),
		RelationRpcClient:     relation.NewRelation(zrpc.MustNewClient(c.RelationRpcConf)),
		AnalysisRpcClient:     analysis.NewAnalysis(zrpc.MustNewClient(c.AnalysisRpcConf)),
	}
}

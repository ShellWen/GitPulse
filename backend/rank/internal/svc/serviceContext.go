package svc

import (
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/analysis"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"github.com/ShellWen/GitPulse/rank/internal/config"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	AsynqClient *asynq.Client
	RedisClient *redis.Redis

	DeveloperRpcClient developer.DeveloperZrpcClient
	AnalysisRpcClient  analysis.Analysis
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

		DeveloperRpcClient: developer.NewDeveloperZrpcClient(zrpc.MustNewClient(c.DeveloperRpcConf)),
		AnalysisRpcClient:  analysis.NewAnalysis(zrpc.MustNewClient(c.AnalysisRpcConf)),
	}
}

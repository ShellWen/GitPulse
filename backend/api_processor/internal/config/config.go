package config

import (
	"github.com/ShellWen/GitPulse/common/config"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf
	DeveloperRpcConf    zrpc.RpcClientConf
	RelationRpcConf     zrpc.RpcClientConf
	ContributionRpcConf zrpc.RpcClientConf
	RepoRpcConf         zrpc.RpcClientConf
	AnalysisRpcConf     zrpc.RpcClientConf

	RedisClient redis.RedisConf

	AsynqRedisConf config.AsynqRedisConf
}

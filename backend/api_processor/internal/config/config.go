package config

import (
	"github.com/ShellWen/GitPulse/common/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	DeveloperRpcConf    zrpc.RpcClientConf
	RelationRpcConf     zrpc.RpcClientConf
	ContributionRpcConf zrpc.RpcClientConf
	RepoRpcConf         zrpc.RpcClientConf
	AnalysisRpcConf     zrpc.RpcClientConf

	Log         logx.LogConf
	RedisClient redis.RedisConf

	AsynqRedisConf config.AsynqRedisConf
}

package config

import (
	"github.com/ShellWen/GitPulse/common/config"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	DB struct {
		DataSource string
	}
	Cache          cache.CacheConf
	AsynqRedisConf config.AsynqRedisConf

	DeveloperRpcConf    zrpc.RpcClientConf
	RelationRpcConf     zrpc.RpcClientConf
	RepoRpcConf         zrpc.RpcClientConf
	ContributionRpcConf zrpc.RpcClientConf
	AnalysisRpcConf     zrpc.RpcClientConf
}

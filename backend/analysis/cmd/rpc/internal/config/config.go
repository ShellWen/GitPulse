package config

import (
	"github.com/ShellWen/GitPulse/common/config"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	Consul              consul.Conf
	DeveloperRpcConf    zrpc.RpcClientConf
	RelationRpcConf     zrpc.RpcClientConf
	RepoRpcConf         zrpc.RpcClientConf
	ContributionRpcConf zrpc.RpcClientConf

	DB struct {
		DataSource string
	}
	Cache          cache.CacheConf
	RedisConf      redis.RedisConf
	SparkModelConf config.SparkModelConf
}

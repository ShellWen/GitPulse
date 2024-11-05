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
	Cache            cache.CacheConf
	DeveloperRpcConf zrpc.RpcClientConf
	AsynqRedisConf   config.AsynqRedisConf
}

package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	RpcServerConf zrpc.RpcServerConf
	RpcClientConf zrpc.RpcClientConf
	DB            struct {
		DataSource string
	}
	Cache cache.CacheConf
}

package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	Cache                                 cache.CacheConf
	KqDeveloperUpdateConsumerConf         kq.KqConf
	KqDeveloperUpdateCompleteConsumerConf kq.KqConf
}

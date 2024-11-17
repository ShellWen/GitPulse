package config

import (
	"github.com/ShellWen/GitPulse/common/config"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	Consul consul.Conf
	DB     struct {
		DataSource string
	}
	Cache                    cache.CacheConf
	KqRepoUpdateConsumerConf kq.KqConf
	AsynqRedisConf           config.AsynqRedisConf
}

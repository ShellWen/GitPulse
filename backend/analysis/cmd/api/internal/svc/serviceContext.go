package svc

import (
	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/config"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	RpcClient           zrpc.Client
	KqFetcherTaskPusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		RpcClient:           zrpc.MustNewClient(c.RpcClientConf),
		KqFetcherTaskPusher: kq.NewPusher(c.KqFetcherTaskPusherConf.Brokers, c.KqFetcherTaskPusherConf.Topic, kq.WithAllowAutoTopicCreation(), kq.WithSyncPush()),
	}
}

package svc

import (
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/config"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	DeveloperRpc        developer.DeveloperZrpcClient
	KqFetcherTaskPusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		DeveloperRpc:        developer.NewDeveloperZrpcClient(zrpc.MustNewClient(c.DeveloperRpcConf)),
		KqFetcherTaskPusher: kq.NewPusher(c.KqFetcherTaskPusherConf.Brokers, c.KqFetcherTaskPusherConf.Topic, kq.WithAllowAutoTopicCreation(), kq.WithSyncPush()),
	}
}

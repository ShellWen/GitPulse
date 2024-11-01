package svc

import (
	"github.com/ShellWen/GitPulse/fetcher/internal/config"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config               config.Config
	RpcClient            zrpc.Client
	KqDeveloperPusher    *kq.Pusher
	KqContributionPusher *kq.Pusher
	KqCreateRepoPusher   *kq.Pusher
	KqForkPusher         *kq.Pusher
	KqStarPusher         *kq.Pusher
	KqFollowPusher       *kq.Pusher
	KqRepoPusher         *kq.Pusher
	RedisClient          *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:               c,
		RpcClient:            zrpc.MustNewClient(c.RpcClientConf),
		KqDeveloperPusher:    kq.NewPusher(c.KqDeveloperPusherConf.Brokers, c.KqDeveloperPusherConf.Topic, kq.WithAllowAutoTopicCreation()),
		KqContributionPusher: kq.NewPusher(c.KqContributionPusherConf.Brokers, c.KqContributionPusherConf.Topic, kq.WithAllowAutoTopicCreation()),
		KqCreateRepoPusher:   kq.NewPusher(c.KqCreateRepoPusherConf.Brokers, c.KqCreateRepoPusherConf.Topic, kq.WithAllowAutoTopicCreation()),
		KqForkPusher:         kq.NewPusher(c.KqForkPusherConf.Brokers, c.KqForkPusherConf.Topic, kq.WithAllowAutoTopicCreation()),
		KqStarPusher:         kq.NewPusher(c.KqStarPusherConf.Brokers, c.KqStarPusherConf.Topic, kq.WithAllowAutoTopicCreation()),
		KqFollowPusher:       kq.NewPusher(c.KqFollowPusherConf.Brokers, c.KqFollowPusherConf.Topic, kq.WithAllowAutoTopicCreation()),
		KqRepoPusher:         kq.NewPusher(c.KqRepoPusherConf.Brokers, c.KqRepoPusherConf.Topic, kq.WithAllowAutoTopicCreation()),
		RedisClient:          redis.MustNewRedis(c.RedisClient),
	}
}

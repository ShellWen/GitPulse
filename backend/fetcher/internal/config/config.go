package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type KqPusherConf struct {
	Brokers []string
	Topic   string
}

type Config struct {
	zrpc.RpcClientConf
	logx.LogConf
	KqFetcherTaskConsumerConf kq.KqConf
	KqDeveloperPusherConf     KqPusherConf
	KqContributionPusherConf  KqPusherConf
	KqCreateRepoPusherConf    KqPusherConf
	KqForkPusherConf          KqPusherConf
	KqStarPusherConf          KqPusherConf
	KqFollowPusherConf        KqPusherConf
	KqRepoPusherConf          KqPusherConf
	RedisClient               redis.RedisConf
}

package config

import (
	"github.com/ShellWen/GitPulse/common/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	DeveloperRpcConf    zrpc.RpcClientConf
	RelationRpcConf     zrpc.RpcClientConf
	ContributionRpcConf zrpc.RpcClientConf
	RepoRpcConf         zrpc.RpcClientConf

	Log                                    logx.LogConf
	KqDeveloperPusherConf                  config.KqPusherConf
	KqContributionPusherConf               config.KqPusherConf
	KqCreateRepoPusherConf                 config.KqPusherConf
	KqForkPusherConf                       config.KqPusherConf
	KqStarPusherConf                       config.KqPusherConf
	KqFollowPusherConf                     config.KqPusherConf
	KqRepoPusherConf                       config.KqPusherConf
	RedisClient                            redis.RedisConf
	KqDeveloperUpdateCompletePusherConf    config.KqPusherConf
	KqRepoUpdateCompletePusherConf         config.KqPusherConf
	KqContributionUpdateCompletePusherConf config.KqPusherConf
	KqRelationUpdateCompletePusherConf     config.KqPusherConf

	AsynqRedisConf config.AsynqRedisConf
}

package svc

import (
	"github.com/ShellWen/GitPulse/common/tasks"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/contribution"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"github.com/ShellWen/GitPulse/fetcher/internal/config"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/relation"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/repo"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"time"
)

type ServiceContext struct {
	Config                config.Config
	DeveloperRpcClient    developer.DeveloperZrpcClient
	RelationRpcClient     relation.Relation
	ContributionRpcClient contribution.ContributionZrpcClient
	RepoRpcClient         repo.RepoZrpcClient

	KqDeveloperPusher                  *kq.Pusher
	KqContributionPusher               *kq.Pusher
	KqCreateRepoPusher                 *kq.Pusher
	KqForkPusher                       *kq.Pusher
	KqStarPusher                       *kq.Pusher
	KqFollowPusher                     *kq.Pusher
	KqRepoPusher                       *kq.Pusher
	RedisClient                        *redis.Redis
	KqDeveloperUpdateCompletePusher    *kq.Pusher
	KqRepoUpdateCompletePusher         *kq.Pusher
	KqContributionUpdateCompletePusher *kq.Pusher
	KqRelationUpdateCompletePusher     *kq.Pusher

	AsynqServer *asynq.Server
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		DeveloperRpcClient:    developer.NewDeveloperZrpcClient(zrpc.MustNewClient(c.DeveloperRpcConf)),
		RelationRpcClient:     relation.NewRelation(zrpc.MustNewClient(c.RelationRpcConf)),
		ContributionRpcClient: contribution.NewContributionZrpcClient(zrpc.MustNewClient(c.ContributionRpcConf)),
		RepoRpcClient:         repo.NewRepoZrpcClient(zrpc.MustNewClient(c.RepoRpcConf)),

		KqDeveloperPusher:                  kq.NewPusher(c.KqDeveloperPusherConf.Brokers, c.KqDeveloperPusherConf.Topic, kq.WithAllowAutoTopicCreation(), kq.WithSyncPush()),
		KqContributionPusher:               kq.NewPusher(c.KqContributionPusherConf.Brokers, c.KqContributionPusherConf.Topic, kq.WithAllowAutoTopicCreation(), kq.WithSyncPush()),
		KqCreateRepoPusher:                 kq.NewPusher(c.KqCreateRepoPusherConf.Brokers, c.KqCreateRepoPusherConf.Topic, kq.WithAllowAutoTopicCreation(), kq.WithSyncPush()),
		KqForkPusher:                       kq.NewPusher(c.KqForkPusherConf.Brokers, c.KqForkPusherConf.Topic, kq.WithAllowAutoTopicCreation(), kq.WithSyncPush()),
		KqStarPusher:                       kq.NewPusher(c.KqStarPusherConf.Brokers, c.KqStarPusherConf.Topic, kq.WithAllowAutoTopicCreation(), kq.WithSyncPush()),
		KqFollowPusher:                     kq.NewPusher(c.KqFollowPusherConf.Brokers, c.KqFollowPusherConf.Topic, kq.WithAllowAutoTopicCreation(), kq.WithSyncPush()),
		KqRepoPusher:                       kq.NewPusher(c.KqRepoPusherConf.Brokers, c.KqRepoPusherConf.Topic, kq.WithAllowAutoTopicCreation(), kq.WithSyncPush()),
		RedisClient:                        redis.MustNewRedis(c.RedisClient),
		KqDeveloperUpdateCompletePusher:    kq.NewPusher(c.KqDeveloperUpdateCompletePusherConf.Brokers, c.KqDeveloperUpdateCompletePusherConf.Topic, kq.WithAllowAutoTopicCreation(), kq.WithSyncPush()),
		KqRepoUpdateCompletePusher:         kq.NewPusher(c.KqRepoUpdateCompletePusherConf.Brokers, c.KqRepoUpdateCompletePusherConf.Topic, kq.WithAllowAutoTopicCreation(), kq.WithSyncPush()),
		KqContributionUpdateCompletePusher: kq.NewPusher(c.KqContributionUpdateCompletePusherConf.Brokers, c.KqContributionUpdateCompletePusherConf.Topic, kq.WithAllowAutoTopicCreation(), kq.WithSyncPush()),
		KqRelationUpdateCompletePusher:     kq.NewPusher(c.KqRelationUpdateCompletePusherConf.Brokers, c.KqRelationUpdateCompletePusherConf.Topic, kq.WithAllowAutoTopicCreation(), kq.WithSyncPush()),

		AsynqServer: asynq.NewServer(
			asynq.RedisClientOpt{
				Addr:     c.AsynqRedisConf.Addr,
				Password: c.AsynqRedisConf.Password,
				DB:       c.AsynqRedisConf.DB,
			}, asynq.Config{
				Concurrency: 20,
				RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
					return tasks.FetchRetryDelay
				},
				Queues: map[string]int{tasks.FetcherTaskQueue: 1},
			}),
	}
}

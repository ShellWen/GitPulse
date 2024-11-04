package svc

import (
	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/config"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/analysis"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/contribution"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/relation"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/repo"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	RpcClient           zrpc.Client
	KqFetcherTaskPusher *kq.Pusher

	DeveloperRpcClient    developer.DeveloperZrpcClient
	RepoRpcClient         repo.RepoZrpcClient
	ContributionRpcClient contribution.ContributionZrpcClient
	RelationRpcClient     relation.Relation
	AnalysisRpcClient     analysis.Analysis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		KqFetcherTaskPusher: kq.NewPusher(c.KqFetcherTaskPusherConf.Brokers, c.KqFetcherTaskPusherConf.Topic, kq.WithAllowAutoTopicCreation(), kq.WithSyncPush()),

		DeveloperRpcClient:    developer.NewDeveloperZrpcClient(zrpc.MustNewClient(c.DeveloperRpcConf)),
		RepoRpcClient:         repo.NewRepoZrpcClient(zrpc.MustNewClient(c.RepoRpcConf)),
		ContributionRpcClient: contribution.NewContributionZrpcClient(zrpc.MustNewClient(c.ContributionRpcConf)),
		RelationRpcClient:     relation.NewRelation(zrpc.MustNewClient(c.RelationRpcConf)),
		AnalysisRpcClient:     analysis.NewAnalysis(zrpc.MustNewClient(c.AnalysisRpcConf)),
	}
}

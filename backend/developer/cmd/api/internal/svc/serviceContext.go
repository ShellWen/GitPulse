package svc

import (
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/config"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	DeveloperRpc developer.DeveloperZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		DeveloperRpc: developer.NewDeveloperZrpcClient(zrpc.MustNewClient(c.DeveloperRpcConf)),
	}
}

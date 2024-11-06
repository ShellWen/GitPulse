package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/logic"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"

	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/config"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/server"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "analysis/cmd/rpc/etc/analysis.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterAnalysisServer(grpcServer, server.NewAnalysisServer(ctx))

		if c.RpcServerConf.Mode == service.DevMode || c.RpcServerConf.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	_ = consul.RegisterService(c.RpcServerConf.ListenOn, c.Consul)

	if err := logic.BuildRank(context.Background(), ctx); err != nil {
		panic(err)
	}

	fmt.Printf("Starting rpc server at %s...\n", c.RpcServerConf.ListenOn)
	s.Start()
}

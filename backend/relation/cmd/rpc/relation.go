package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/consumer"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/config"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/server"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "relation/cmd/rpc/etc/relation.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterRelationServer(grpcServer, server.NewRelationServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	serviceGroup.Add(s)

	_ = consul.RegisterService(c.ListenOn, c.Consul)

	for _, s := range consumer.Consumers(c, context.Background(), ctx) {
		serviceGroup.Add(s)
	}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	serviceGroup.Start()
}

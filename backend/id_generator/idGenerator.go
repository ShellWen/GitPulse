package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"

	"github.com/ShellWen/GitPulse/id_generator/internal/config"
	"github.com/ShellWen/GitPulse/id_generator/internal/server"
	"github.com/ShellWen/GitPulse/id_generator/internal/svc"
	"github.com/ShellWen/GitPulse/id_generator/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "id_generator/etc/idGenerator.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterIdGeneratorServer(grpcServer, server.NewIdGeneratorServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	_ = consul.RegisterService(c.ListenOn, c.Consul)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

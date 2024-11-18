package main

import (
	"context"
	"flag"
	"github.com/ShellWen/GitPulse/api_processor/internal/config"
	"github.com/ShellWen/GitPulse/api_processor/internal/consumer"
	"github.com/ShellWen/GitPulse/api_processor/internal/svc"
	"github.com/joho/godotenv"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

var configFile = flag.String("f", "api_processor/etc/apiProcessor.yaml", "the config file")

func main() {
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		logx.Error("load .env file failed: %v", err)
	}

	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.MustSetup(c.Log)
	logx.Info("api processor started")

	ctx := context.Background()
	svcContext := svc.NewServiceContext(c)
	taskConsumer := consumer.NewAPITaskConsumer(ctx, svcContext)
	mux := taskConsumer.Register()

	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.Error(err)
		return
	}
}

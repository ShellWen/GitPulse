package main

import (
	"context"
	"flag"
	"github.com/ShellWen/GitPulse/fetcher/internal/config"
	"github.com/ShellWen/GitPulse/fetcher/internal/consumer"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/joho/godotenv"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "fetcher/etc/fetcher.yaml", "the config file")

func main() {
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		logx.Error("load .env file failed: %v", err)
	}

	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.MustSetup(c.LogConf)
	logx.Info("fetcher started")

	ctx := context.Background()
	svcContext := svc.NewServiceContext(c)

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, s := range consumer.Consumers(c, ctx, svcContext) {
		serviceGroup.Add(s)
	}

	serviceGroup.Start()
}

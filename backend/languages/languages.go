package main

import (
	"flag"
	"fmt"
	"github.com/ShellWen/GitPulse/languages/internal/config"
	"github.com/ShellWen/GitPulse/languages/internal/handler"
	"github.com/ShellWen/GitPulse/languages/internal/svc"
	"github.com/joho/godotenv"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

var configFile = flag.String("f", "languages/etc/languages.yaml", "the config file")

func main() {
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		logx.Error("load .env file failed: %v", err)
	}

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

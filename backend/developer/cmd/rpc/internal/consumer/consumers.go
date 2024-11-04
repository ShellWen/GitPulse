package consumer

import (
	"context"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/internal/config"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/internal/svc"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

func Consumers(c config.Config, ctx context.Context, svc *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.KqDeveloperUpdateConsumerConf, NewDeveloperUpdateConsumer(ctx, svc)),
		kq.MustNewQueue(c.KqDeveloperUpdateCompleteConsumerConf, NewDeveloperUpdateCompleteConsumer(ctx, svc)),
	}
}

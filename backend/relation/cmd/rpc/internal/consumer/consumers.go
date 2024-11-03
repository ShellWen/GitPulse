package consumer

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/config"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

func Consumers(c config.Config, ctx context.Context, svc *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.KqCreateRepoUpdateConsumerConf, NewCreateRepoUpdateConsumer(ctx, svc)),
		kq.MustNewQueue(c.KqFollowUpdateConsumerConf, NewFollowUpdateConsumer(ctx, svc)),
		kq.MustNewQueue(c.KqForkUpdateConsumerConf, NewForkUpdateConsumer(ctx, svc)),
		kq.MustNewQueue(c.KqStarUpdateConsumerConf, NewStarUpdateConsumer(ctx, svc)),
	}
}

package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/api_processor/internal/svc"
	"time"
)

const expireTime = time.Minute * 10

func SaveToRedis(ctx context.Context, svcCtx *svc.ServiceContext, key string, data string) error {
	err := svcCtx.RedisClient.SetCtx(ctx, key, data)
	if err != nil {
		return err
	}

	err = svcCtx.RedisClient.ExpireCtx(ctx, key, int(expireTime.Seconds()))
	if err != nil {
		return err
	}

	return nil
}

package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/model"
	"strconv"
)

func BuildRank(ctx context.Context, svcCtx *svc.ServiceContext) (err error) {
	var (
		pulsePoints []*model.PulsePoint
	)

	if pulsePoints, err = svcCtx.PulsePointModel.FindAll(ctx); err != nil {
		return
	}

	for _, pulsePoint := range pulsePoints {
		if _, err = svcCtx.RedisClient.ZaddFloatCtx(ctx, "pulse_point_rank", pulsePoint.PulsePoint, strconv.FormatInt(pulsePoint.DeveloperId, 10)); err != nil {
			return
		}
	}

	return
}

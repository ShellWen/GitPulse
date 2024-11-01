package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/common/consts"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

func fetchWithRetry(ctx context.Context, svcContext *svc.ServiceContext, Id int64, topic string, fetchFunc func(context.Context, *svc.ServiceContext, int64) error) (err error) {
	var counter int

	err = fetchFunc(ctx, svcContext, Id)
	for counter = 1; err != nil && counter <= consts.FetchRetryingTimes; counter++ {
		logx.Info("Retrying to fetch...  " + topic + "  Retry time: " + strconv.Itoa(counter))
		err = fetchFunc(ctx, svcContext, Id)
	}
	if err != nil {
		logx.Error(errors.New("Failed to fetch " + topic + ": " + err.Error()))
	}

	return
}

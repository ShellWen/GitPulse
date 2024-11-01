package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strconv"
)

func statusCodeHandler(ctx context.Context, svcContext *svc.ServiceContext, topic string, f func() error, statusCode int) (err error) {
	switch statusCode {
	case http.StatusNotModified:
		logx.Info(topic + " not modified")
		err = nil
	case http.StatusOK:
		err = f()
		logx.Info("Successfully push the update task of " + topic)
	default:
		err = errors.New("unexpected status code " + strconv.Itoa(statusCode))
	}
	return
}

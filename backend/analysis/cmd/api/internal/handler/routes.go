// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"
	"time"

	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/developers/:login/languages",
				Handler: getLanguageUsageHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/developers/:login/pulse-point",
				Handler: getPulsePointHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/developers/:login/region",
				Handler: getRegionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/languages",
				Handler: getLanguagesHandler(serverCtx),
			},
		},
		rest.WithTimeout(60000*time.Millisecond),
	)
}

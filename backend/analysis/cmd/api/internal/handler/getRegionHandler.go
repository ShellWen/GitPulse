package handler

import (
	"github.com/google/go-github/v66/github"
	"net/http"

	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/logic"
	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getRegionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRegionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetRegionLogic(r.Context(), svcCtx)
		resp, err := l.GetRegion(&req)
		if err != nil {
			if err.(*github.ErrorResponse) != nil {
				w.WriteHeader(err.(*github.ErrorResponse).Response.StatusCode)
			}
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package handler

import (
	"net/http"

	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/logic"
	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getPulsePointHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPulsePointReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetPulsePointLogic(r.Context(), svcCtx)
		resp, err := l.GetPulsePoint(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

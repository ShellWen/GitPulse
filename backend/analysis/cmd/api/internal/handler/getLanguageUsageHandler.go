package handler

import (
	"errors"
	"github.com/google/go-github/v66/github"
	"net/http"

	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/logic"
	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getLanguageUsageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetLanguageUsageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetLanguageUsageLogic(r.Context(), svcCtx)
		resp, err := l.GetLanguageUsage(&req)
		if err != nil {
			var value *github.ErrorResponse
			ok := errors.As(err, &value)
			if ok {
				w.WriteHeader(value.Response.StatusCode)
			}
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

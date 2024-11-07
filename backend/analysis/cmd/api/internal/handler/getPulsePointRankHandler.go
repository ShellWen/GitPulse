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

func getPulsePointRankHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPulsePointRankReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if req.Limit < 1 {
			req.Limit = 1
		} else if req.Limit > 100 {
			req.Limit = 100
		}

		l := logic.NewGetPulsePointRankLogic(r.Context(), svcCtx)
		resp, err := l.GetPulsePointRank(&req)
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

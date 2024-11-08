package developer

import (
	"errors"
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/logic/developer"
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/types"
	"github.com/google/go-github/v66/github"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetDeveloperHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDeveloperReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := developer.NewGetDeveloperLogic(r.Context(), svcCtx)
		resp, err := l.GetDeveloper(&req)
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

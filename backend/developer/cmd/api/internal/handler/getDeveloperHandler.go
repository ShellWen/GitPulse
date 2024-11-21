package handler

import (
	"errors"
	"github.com/google/go-github/v66/github"
	zeroErrors "github.com/zeromicro/x/errors"
	"net/http"

	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/logic"
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
)

func getDeveloperHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDeveloperReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetDeveloperLogic(r.Context(), svcCtx)
		devResp, err := l.GetDeveloper(&req)
		var codeMsg *zeroErrors.CodeMsg
		var ghErrResp *github.ErrorResponse
		if err != nil {
			switch {
			case errors.As(err, &codeMsg):
				w.WriteHeader(codeMsg.Code)
			case errors.As(err, &ghErrResp):
				err = zeroErrors.New(ghErrResp.Response.StatusCode, "GitHub: "+ghErrResp.Message)
				w.WriteHeader(ghErrResp.Response.StatusCode)
			default:
				err = zeroErrors.New(http.StatusInternalServerError, err.Error())
				w.WriteHeader(http.StatusInternalServerError)
			}
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else if devResp != nil {
			httpx.OkJsonCtx(r.Context(), w, devResp)
		} else {
			err = zeroErrors.New(http.StatusInternalServerError, "Internal server error")
			w.WriteHeader(http.StatusInternalServerError)
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		}
	}
}

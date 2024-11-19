package handler

import (
	"errors"
	"github.com/ShellWen/GitPulse/rank/internal/svc"
	"github.com/ShellWen/GitPulse/rank/internal/types"
	"github.com/google/go-github/v66/github"
	zeroErrors "github.com/zeromicro/x/errors"
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"github.com/ShellWen/GitPulse/rank/internal/logic"

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
		} else if resp != nil {
			httpx.OkJsonCtx(r.Context(), w, resp)
		} else {
			err = zeroErrors.New(http.StatusInternalServerError, "Internal server error")
			w.WriteHeader(http.StatusInternalServerError)
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		}
	}
}

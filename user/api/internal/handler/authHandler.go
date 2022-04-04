package handler

import (
	"devops/common/response"
	"net/http"

	"devops/user/api/internal/logic"
	"devops/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAuthLogic(r.Context(), svcCtx)
		err := l.Auth(r, w)
		if err != nil {
			httpx.WriteJson(w, http.StatusUnauthorized, response.HandlerError(err))
		} else {
			httpx.Ok(w)
		}
	}
}

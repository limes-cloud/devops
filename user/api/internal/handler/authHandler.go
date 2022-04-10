package handler

import (
	"devops/common/errorx"
	"devops/common/response"
	"net/http"

	"devops/user/api/internal/logic"
	"devops/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAuthLogic(r.Context(), svcCtx)
		if err := l.Auth(r, w); err != nil {
			if err.Error() == errorx.AuthErr {
				httpx.WriteJson(w, http.StatusUnauthorized, response.HandlerError(err))
			}
			if err.Error() == errorx.RbacErr {
				httpx.WriteJson(w, http.StatusForbidden, response.HandlerError(err))
			}
		} else {
			httpx.Ok(w)
		}
	}
}

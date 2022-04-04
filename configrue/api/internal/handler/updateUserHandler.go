package handler

import (
	"configure/api/internal/logic"
	"configure/api/internal/svc"
	"configure/api/internal/types"
	"configure/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UpdateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJson(w, 200, response.HandlerError(err))
			return
		}

		l := logic.NewUpdateUserLogic(r.Context(), svcCtx)
		err := l.UpdateUser(&req)
		if err != nil {
			httpx.WriteJson(w, 200, response.HandlerError(err))
		} else {
			httpx.Ok(w)
		}
	}
}

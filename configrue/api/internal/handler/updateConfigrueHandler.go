package handler

import (
	"devops/common/response"
	"devops/configrue/api/internal/logic"
	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UpdateConfigrueHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateConfigureRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJson(w, 200, response.HandlerError(err))
			return
		}

		l := logic.NewUpdateConfigrueLogic(r.Context(), svcCtx)
		err := l.UpdateConfigrue(&req)
		if err != nil {
			httpx.WriteJson(w, 200, response.HandlerError(err))
		} else {
			httpx.OkJson(w, response.HandlerResp(nil))
		}
	}
}

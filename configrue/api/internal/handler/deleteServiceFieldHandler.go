package handler

import (
	"devops/common/response"
	"devops/configrue/api/internal/logic"
	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func DeleteServiceFieldHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteServiceFieldRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJson(w, 200, response.HandlerError(err))
			return
		}

		l := logic.NewDeleteServiceFieldLogic(r.Context(), svcCtx)
		err := l.DeleteServiceField(&req)
		if err != nil {
			httpx.WriteJson(w, 200, response.HandlerError(err))
		} else {
			httpx.OkJson(w, response.HandlerResp(nil))
		}
	}
}

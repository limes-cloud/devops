package handler

import (
	"configure/api/internal/logic"
	"configure/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"video/common/response"
)

func GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser()
		if err != nil {
			httpx.WriteJson(w, 200, response.HandlerError(err))
		} else {
			httpx.OkJson(w, response.HandlerResp(resp))
		}
	}
}

package handler

import (
	"devops/common/response"
	"devops/configrue/api/internal/logic"
	"devops/configrue/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GetEnvironmentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetEnvironmentLogic(r.Context(), svcCtx)
		resp, err := l.GetEnvironment()
		if err != nil {
			httpx.WriteJson(w, 200, response.HandlerError(err))
		} else {
			httpx.OkJson(w, response.HandlerResp(resp))
		}
	}
}

package handler

import (
	"configure/api/internal/logic"
	"configure/api/internal/svc"
	"configure/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GetTeamHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetTeamLogic(r.Context(), svcCtx)
		resp, err := l.GetTeam()
		if err != nil {
			httpx.WriteJson(w, 200, response.HandlerError(err))
		} else {
			httpx.OkJson(w, response.HandlerResp(resp))
		}
	}
}

package {{.PkgName}}

import (
	"net/http"
    "devops/common/response"
	{{if .After1_1_10}}"github.com/zeromicro/go-zero/rest/httpx"{{end}}
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJson(w, 200, response.HandlerError(err))
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
			httpx.WriteJson(w, 200, response.HandlerError(err))
		} else {
		    {{if .HasResp}}httpx.OkJson(w, response.HandlerResp(resp)){{else}}httpx.OkJson(w, response.HandlerResp(nil)){{end}}
		}
	}
}

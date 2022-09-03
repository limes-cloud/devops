package types

type PageOperatorLogRequest struct {
	ServiceKeyword string `json:"service_keyword" form:"service_keyword"`
	ServiceName    string `json:"service_name"  form:"service_name" sql:"like '%?%'"`
	Page           int    `json:"page" form:"page" sql:"-"`
	Count          int    `json:"count" form:"count" sql:"-"`
}

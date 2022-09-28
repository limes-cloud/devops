package types

type GetLogRequest struct {
	Page  int    `json:"page" form:"page" sql:"-"`
	Count int    `json:"count" form:"count" sql:"-"`
	Title string `json:"title" form:"log"  sql:"like '%?%'"`
	Start int64  `json:"start" form:"start" sql:"> ?" field:"created_at"`
	End   int64  `json:"end" form:"end" sql:"< ?" field:"created_at"`
}

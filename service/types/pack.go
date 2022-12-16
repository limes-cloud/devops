package types

type PagePackRequest struct {
	Page           int    `json:"page" form:"page" binding:"required" sql:"-"`
	Count          int    `json:"count" form:"count"  binding:"required,max=50"  sql:"-"`
	ServiceKeyword string `json:"service_keyword" form:"service_keyword"  binding:"required"`
	IsFinish       *bool  `json:"is_finish" form:"is_finish"`
	IsClear        *bool  `json:"is_clear" form:"is_finish"`
	Status         *bool  `json:"status" form:"status"`
	Start          int64  `json:"start" form:"start" sql:"> ?" field:"created_at"`
	End            int64  `json:"end" form:"end" sql:"< ?" field:"created_at"`
}

type AddPackRequest struct {
	CloneType      string `json:"clone_type"  binding:"required"`
	CloneValue     string `json:"clone_value" binding:"required"`
	CommitID       string `json:"commit_id"  binding:"required"`
	ServiceKeyword string `json:"service_keyword" binding:"required"`
}

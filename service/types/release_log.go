package types

type PageReleaseLogRequest struct {
	Page           int     `json:"page" form:"page" binding:"required" sql:"-"`
	Count          int     `json:"count" form:"count"  binding:"required,max=50"  sql:"-"`
	ServiceKeyword string  `json:"service_keyword" form:"service_keyword"  binding:"required"`
	IsFinish       *bool   `json:"is_finish" form:"is_finish"`
	Status         *string `json:"status" form:"status"`
	Start          int64   `json:"start" form:"start" sql:"> ?" field:"created_at"`
	End            int64   `json:"end" form:"end" sql:"< ?" field:"created_at"`
}

type AddReleaseLogRequest struct {
	EnvKeyword     string `json:"env_keyword" binding:"required"`
	ServiceKeyword string `json:"service_keyword" binding:"required"`
	PackID         int64  `json:"pack_id"`
}

type AllReleaseImagesRequest struct {
	ServiceKeyword string `form:"service_keyword" json:"service_keyword" binding:"required"`
}

package types

type PageReleaseRequest struct {
	Page  int    `json:"page" form:"page" binding:"required" sql:"-"`
	Count int    `json:"count" form:"count"  binding:"required,max=50"  sql:"-"`
	Name  string `json:"name" form:"name" sql:"like '%?%'"`
	Type  string `json:"type" form:"type"`
}

type AddReleaseRequest struct {
	Name     string `json:"name" binding:"required"`
	Desc     string `json:"desc"`
	Type     string `json:"type"  binding:"required"`
	Template string `json:"template" binding:"required"`
}

type UpdateReleaseRequest struct {
	ID       int64  `json:"id"  binding:"required"`
	Name     string `json:"name" binding:"required"`
	Desc     string `json:"desc"`
	Type     string `json:"type"  binding:"required"`
	Template string `json:"template" binding:"required"`
}

type DeleteReleaseRequest struct {
	ID int64 `json:"id"  binding:"required"`
}

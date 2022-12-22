package types

type PageDockerfileRequest struct {
	Page  int    `json:"page" form:"page" binding:"required" sql:"-"`
	Count int    `json:"count" form:"count"  binding:"required,max=50"  sql:"-"`
	Name  string `json:"name" form:"name" sql:"like '%?%'"`
}

type AddDockerfileRequest struct {
	Name     string `json:"name" binding:"required"`
	Desc     string `json:"desc"`
	Template string `json:"template" binding:"required"`
}

type UpdateDockerfileRequest struct {
	ID       int64  `json:"id"  binding:"required"`
	Name     string `json:"name" binding:"required"`
	Desc     string `json:"desc"`
	Template string `json:"template" binding:"required"`
}

type DeleteDockerfileRequest struct {
	ID int64 `json:"id"  binding:"required"`
}

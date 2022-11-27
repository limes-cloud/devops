package types

type AllServiceFieldRequest struct {
	Keyword string `json:"keyword" form:"keyword"  binding:"required"`
}

type AddServiceRequest struct {
	Keyword     string  `json:"keyword" binding:"required"`
	Name        string  `json:"name"  binding:"required"`
	Description *string `json:"description"`
	EnvIds      []int64 `json:"env_ids" binding:"required"`
}

type AllServiceRequest struct {
	Keyword string `json:"keyword" form:"keyword"`
	Name    string `json:"name" form:"name" sql:"like '%?%'"`
}

type UpdateServiceRequest struct {
	ID          int64   `json:"id" binding:"required"`
	Keyword     string  `json:"keyword"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	EnvIds      []int64 `json:"env_ids" binding:"required"`
}

type DeleteServiceRequest struct {
	ID int64 `json:"id" binding:"required"`
}

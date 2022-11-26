package types

type AllServiceEnvRequest struct {
	SrvId int64 `json:"srv_id" form:"srv_id"  binding:"required"`
}

type AllServiceFieldRequest struct {
	SrvId int64 `json:"srv_id" form:"srv_id"  binding:"required"`
}

type AddServiceRequest struct {
	Keyword     string  `json:"keyword" binding:"required"`
	Name        string  `json:"name"  binding:"required"`
	IsPrivate   *bool   `json:"is_private"  binding:"required"`
	TeamID      *int64  `json:"team_id"`
	Description *string `json:"description"`
	EnvIds      []int64 `json:"env_ids" binding:"required"`
}

type PageServiceRequest struct {
	Keyword   string `json:"keyword" form:"keyword"`
	Name      string `json:"name" form:"name" sql:"like '%?%'"`
	IsPrivate *bool  `json:"is_private" form:"is_private"`
	TeamID    *int64 `json:"team_id"  form:"team_id" sql:"-"`
	Page      int    `json:"page" form:"page" binding:"required" sql:"-"`
	Count     int    `json:"count" form:"count"  binding:"required,max=50"  sql:"-"`
}

type UpdateServiceRequest struct {
	ID          int64   `json:"id" binding:"required"`
	Keyword     string  `json:"keyword"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	IsPrivate   *bool   `json:"is_private"`
	TeamID      *int64  `json:"team_id"`
	EnvIds      []int64 `json:"env_ids" binding:"required"`
}

type DeleteServiceRequest struct {
	ID int64 `json:"id" binding:"required"`
}

package types

type AllEnvironmentRequest struct {
	EnvKeyword string `json:"env_keyword" form:"env_id"`
	Drive      string `json:"drive" form:"drive" sql:"like '%?%'"`
	Status     *bool  `json:"status" form:"status"`
}

type EnvironmentConnectRequest struct {
	Keyword string `json:"keyword"  binding:"required"`
}

type AddEnvironmentRequest struct {
	EnvKeyword string `json:"env_keyword" binding:"required"`
	Drive      string `json:"drive" binding:"required"`
	Config     string `json:"config" binding:"required"`
	Prefix     string `json:"prefix" binding:"required"`
	Status     bool   `json:"status"`
}

type UpdateEnvironmentRequest struct {
	ID         int64  `json:"id"  binding:"required"`
	EnvKeyword string `json:"env_keyword" binding:"required"`
	Drive      string `json:"drive"`
	Config     string `json:"config"`
	Prefix     string `json:"prefix"`
	Token      string `json:"token"`
}

type DeleteEnvironmentRequest struct {
	ID int64 `json:"id"  binding:"required"`
}

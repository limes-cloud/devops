package types

type AllEnvironmentRequest struct {
	EnvID  int64  `json:"env_id" form:"env_id"`
	Drive  string `json:"drive" form:"drive" sql:"like '%?%'"`
	Status *bool  `json:"status" form:"status"`
}

type EnvironmentConnectRequest struct {
	ID int64 `json:"id"  binding:"required"`
}

type AddEnvironmentRequest struct {
	EnvID  int64  `json:"env_id" binding:"required"`
	Drive  string `json:"drive" binding:"required"`
	Config string `json:"config" binding:"required"`
	Prefix string `json:"prefix" binding:"required"`
	Status bool   `json:"status"`
}

type UpdateEnvironmentRequest struct {
	ID     int64  `json:"id"  binding:"required"`
	EnvID  int64  `json:"env_id" binding:"required"`
	Drive  string `json:"drive"`
	Config string `json:"config"`
	Prefix string `json:"prefix"`
	Token  string `json:"token"`
}

type DeleteEnvironmentRequest struct {
	ID int64 `json:"id"  binding:"required"`
}

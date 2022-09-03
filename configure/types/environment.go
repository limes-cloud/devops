package types

type AllEnvironmentRequest struct {
	Keyword string `json:"keyword" form:"keyword"`
	Name    string `json:"name" form:"name" sql:"like '%?%'"`
	Drive   string `json:"drive" form:"drive" sql:"like '%?%'"`
	Status  *bool  `json:"status" form:"status"`
}

type EnvironmentConnectRequest struct {
	ID int64 `json:"id"  binding:"required"`
}

type AddEnvironmentRequest struct {
	Keyword     string `json:"keyword" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Drive       string `json:"drive" binding:"required"`
	Config      string `json:"config" binding:"required"`
	Prefix      string `json:"prefix" binding:"required"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type UpdateEnvironmentRequest struct {
	ID          int64  `json:"id"  binding:"required"`
	Name        string `json:"name"`
	Keyword     string `json:"keyword"`
	Drive       string `json:"drive"`
	Config      string `json:"config"`
	Prefix      string `json:"prefix"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	Token       string `json:"token"`
}

type DeleteEnvironmentRequest struct {
	ID int64 `json:"id"  binding:"required"`
}

type UpdateEnvServiceRequest struct {
	ID     int64   `json:"id"  binding:"required"`
	SrvIds []int64 `json:"srv_ids" binding:"required"`
}

type AllEnvServiceRequest struct {
	EnvId int64 `json:"env_id" form:"env_id"`
	SrvId int64 `json:"srv_id" form:"srv_id"`
}

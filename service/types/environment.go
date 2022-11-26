package types

type AllEnvironmentRequest struct {
	Keyword string `json:"keyword" form:"keyword"`
	Name    string `json:"name" form:"name" sql:"like '%?%'"`
	Status  *bool  `json:"status" form:"status"`
}

type EnvironmentConnectRequest struct {
	ID int64 `json:"id"  binding:"required"`
}

type AddEnvironmentRequest struct {
	Keyword     string `json:"keyword" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type UpdateEnvironmentRequest struct {
	ID          int64  `json:"id"  binding:"required"`
	Name        string `json:"name"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type DeleteEnvironmentRequest struct {
	ID int64 `json:"id"  binding:"required"`
}

type UpdateServiceEnvRequest struct {
	ID     int64   `json:"id"  binding:"required"`
	SrvIds []int64 `json:"srv_ids" binding:"required"`
}

//type AllServiceEnvRequest struct {
//	EnvId int64 `json:"env_id" form:"env_id"`
//	SrvId int64 `json:"srv_id" form:"srv_id"`
//}

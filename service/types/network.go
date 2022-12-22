package types

type PageNetworkRequest struct {
	Page  int    `json:"page" form:"page" binding:"required" sql:"-"`
	Count int    `json:"count" form:"count"  binding:"required,max=50"  sql:"-"`
	Host  string `json:"host" form:"host" sql:"like '%?%'"`
	SrvID int64  `json:"srv_id" form:"srv_id"`
	EnvID int64  `json:"env_id" form:"env_id"`
}

type AddNetworkRequest struct {
	SrvID    int64  `json:"srv_id"  binding:"required"`
	EnvID    int64  `json:"env_id"  binding:"required"`
	Host     string `json:"host"  binding:"required"`
	Cert     string `json:"cert"`
	Key      string `json:"key"`
	Redirect bool   `json:"redirect"`
}

type UpdateNetworkRequest struct {
	ID       int64   `json:"id"  binding:"required"`
	SrvID    int64   `json:"srv_id"`
	EnvID    int64   `json:"env_id"`
	Host     string  `json:"host"`
	Cert     *string `json:"cert"`
	Key      *string `json:"key"`
	Redirect *bool   `json:"redirect"`
}

type DeleteNetworkRequest struct {
	ID int64 `json:"id"  binding:"required"`
}

package types

type CompareConfigRequest struct {
	SrvKeyword string `json:"srv_keyword" form:"srv_keyword" binding:"required"`
	EnvKeyword string `json:"env_keyword" form:"env_keyword" binding:"required"`
}

type RollbackConfigRequest struct {
	ID int64 `json:"id"  form:"id" binding:"required"`
}

type AllConfigLogRequest struct {
	ServiceKeyword string `json:"service_keyword" form:"service_keyword" binding:"required"`
	EnvKeyword     string `json:"env_keyword" form:"env_keyword" binding:"required"`
}

type DriverConfigRequest struct {
	SrvKeyword string `json:"srv_keyword" form:"srv_keyword" binding:"required"`
	EnvKeyword string `json:"env_keyword" form:"env_keyword" binding:"required"`
}

type ConfigRequest struct {
	Service string `json:"service"  form:"service" binding:"required"`
	Token   string `json:"token" form:"token" binding:"required"`
}

type ConfigLogRequest struct {
	ID int64 `json:"id" form:"id"  binding:"required"`
}

type SyncConfigRequest struct {
	SrvKeyword string `json:"srv_keyword" form:"srv_keyword" binding:"required"`
	EnvKeyword string `json:"env_keyword" form:"env_keyword" binding:"required"`
}

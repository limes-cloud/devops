package types

type CompareConfigRequest struct {
	SrvId int64 `json:"srv_id" binding:"required"`
	EnvId int64 `json:"env_id" binding:"required"`
}

type RollbackConfigRequest struct {
	ID int64 `json:"id"  form:"id" binding:"required"`
}

type AllConfigLogRequest struct {
	SrvId int64 `json:"srv_id" form:"srv_id" binding:"required"`
	EnvId int64 `json:"env_id" form:"env_id" binding:"required"`
}

type DriverConfigRequest struct {
	SrvId int64 `json:"srv_id" form:"srv_id" binding:"required"`
	EnvId int64 `json:"env_id" form:"env_id" binding:"required"`
}

type ConfigLogRequest struct {
	ID int64 `json:"id" form:"id"  binding:"required"`
}

type SyncConfigRequest struct {
	SrvId int64 `json:"srv_id" binding:"required"`
	EnvId int64 `json:"env_id" binding:"required"`
}

type ParseTemplateRequest struct {
	SrvId int64 `json:"srv_id" form:"srv_id" binding:"required"`
	EnvId int64 `json:"env_id" form:"env_id" binding:"required"`
}

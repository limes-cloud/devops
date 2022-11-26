package types

type AddResourceValueRequest struct {
	ResourceID int64 `json:"resource_id"`
	Data       []struct {
		EnvId int64  `json:"env_id"`
		Value string `json:"value"`
	} `json:"data" binding:"required"`
}

type AllResourceValueRequest struct {
	ResourceID int64 `json:"resource_id" form:"resource_id" binding:"required"`
}

package types

type AddResourceValueRequest struct {
	ResourceID int64 `json:"resource_id"`
	Data       []struct {
		EnvKeyword string `json:"env_keyword"`
		Value      string `json:"value"`
	} `json:"data" binding:"required"`
}

type AllResourceValueRequest struct {
	ResourceID int64 `json:"resource_id" form:"resource_id" binding:"required"`
}

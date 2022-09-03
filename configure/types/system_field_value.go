package types

type AddSystemFieldValueRequest struct {
	FieldId int64 `json:"field_id"`
	Data    []struct {
		EnvId int64  `json:"env_id"`
		Value string `json:"value"`
	} `json:"data" binding:"required"`
}

type AllSystemFieldValueRequest struct {
	FieldId int64 `json:"field_id" form:"field_id" binding:"required"`
}

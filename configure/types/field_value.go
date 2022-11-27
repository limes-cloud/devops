package types

type AddFieldValueRequest struct {
	FieldId int64 `json:"field_id"`
	Data    []struct {
		EnvKeyword string `json:"env_keyword"`
		Value      string `json:"value"`
	} `json:"data" binding:"required"`
}

type AllFieldValueRequest struct {
	FieldId int64 `json:"field_id" form:"field_id" binding:"required"`
}

package models

import "configure/common/model"

type Role struct {
	model.BaseModel
	Name        string `json:"name" `
	KeyWord     string `json:"key_word"`
	Status      *bool  `json:"status" `
	Weight      int    `json:"weight"`
	Description string `json:"description"`
	DataScope   string `json:"data_scope" `
	Operator    string `json:"operator"`
	OperatorID  int    `json:"operator_id"`
}

func (u Role) tb() string {
	return "team"
}

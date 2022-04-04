package models

import "devops/common/model"

type Team struct {
	model.BaseModel
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	ParentID    int    `json:"parent_id"`
	Operator    string `json:"operator"`
	OperatorID  int    `json:"operator_id"`
}

func (u Team) Table() string {
	return "team"
}

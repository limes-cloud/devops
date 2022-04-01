package models

import "configure/common/model"

type User struct {
	model.BaseModel
	RoleID     int    `json:"role_id"`
	TeamID     int    `json:"team_id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	Password   string `json:"password,omitempty"`
	Status     *bool  `json:"status" `
	Operator   string `json:"operator"`
	OperatorID int    `json:"operator_id"`
}

func (u User) tb() string {
	return "user"
}

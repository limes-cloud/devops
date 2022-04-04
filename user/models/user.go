package models

import (
	"devops/common/model"
)

type User struct {
	model.BaseModel
	RoleID     int    `json:"role_id"`
	TeamID     int    `json:"team_id"`
	RoleName   string `json:"role_name" gorm:"->"`
	TeamName   string `json:"team_name" gorm:"->"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	Password   string `json:"password,omitempty"`
	Status     *bool  `json:"status" `
	Operator   string `json:"operator"`
	OperatorID int    `json:"operator_id"`
}

// 如果没有自定义缓存。那就应该按照ID来进行缓存数据

func (u User) Table() string {
	return "user"
}

func (u *User) GetUser(userId int64) error {
	return database().Table(u.Table()).
		Select("user.*,team.name team_name,role.name role_name").
		Joins("left join role on role.id = user.role_id").
		Joins("left join team on team.id = user.team_id").
		First(&u, userId).Error
}

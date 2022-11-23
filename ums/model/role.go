package model

import (
	"github.com/limeschool/gin"
)

type Role struct {
	gin.BaseModel
	Name        string `json:"name" `
	Keyword     string `json:"keyword"`
	Status      *bool  `json:"status" `
	Weight      int    `json:"weight"`
	Description string `json:"description"`
	TeamIds     string `json:"team_ids"`
	DataScope   string `json:"data_scope" `
	Operator    string `json:"operator"`
	OperatorID  int64  `json:"operator_id"`
}

func (u *Role) Table() string {
	return "role"
}

func (u *Role) Create(ctx *gin.Context) error {
	user := CurUser(ctx)
	u.OperatorID = user.ID
	u.Operator = user.Name
	return transferErr(database(ctx).Table(u.Table()).Create(&u).Error)
}

func (u *Role) OneById(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).First(u, "id = ?", id).Error)
}

func (u *Role) All(ctx *gin.Context, m interface{}) ([]Role, error) {
	var list []Role
	db := database(ctx).Table(u.Table()).Order(u.Table() + ".weight desc")
	db = gin.GormWhere(db, u.Table(), m)
	return list, transferErr(db.Find(&list).Error)
}

func (u *Role) UpdateByID(ctx *gin.Context) error {
	user := CurUser(ctx)
	u.Operator = user.Name
	u.OperatorID = user.ID
	return transferErr(database(ctx).Table(u.Table()).Updates(u).Error)
}

func (u *Role) DeleteByID(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).Where("id = ?", id).Delete(&u).Error)
}

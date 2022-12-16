package model

import (
	"github.com/limeschool/gin"
	"service/meta"
)

type CodeRegistry struct {
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Type       string `json:"type"`
	Host       string `json:"host"`
	Token      string `json:"token"`
	CloneType  string `json:"clone_type"`
	Operator   string `json:"operator,omitempty"`
	OperatorId int64  `json:"operator_id,omitempty"`
	gin.BaseModel
}

func (u *CodeRegistry) Table() string {
	return "code_registry"
}

func (u *CodeRegistry) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Create(&u).Error)
}

func (u *CodeRegistry) OneById(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).First(u, "id = ?", id).Error)
}

func (u *CodeRegistry) All(ctx *gin.Context, m any, fs ...callback) ([]CodeRegistry, error) {
	var list []CodeRegistry
	db := database(ctx).Table(u.Table())
	db = execCallback(db, fs...)
	db = gin.GormWhere(db, u.Table(), m)
	return list, transferErr(db.Find(&list).Error)
}

func (u *CodeRegistry) UpdateByID(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Updates(u).Error)
}

func (u *CodeRegistry) DeleteByID(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).Where("id = ?", id).Delete(&u).Error)
}

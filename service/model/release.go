package model

import (
	"github.com/limeschool/gin"
	"service/meta"
)

type Release struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Desc       string `json:"desc"`
	Template   string `json:"template"`
	Operator   string `json:"operator,omitempty"`
	OperatorId int64  `json:"operator_id,omitempty"`
	gin.BaseModel
}

func (u *Release) Table() string {
	return "release_template"
}

func (u *Release) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Create(&u).Error)
}

func (u *Release) OneById(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).First(u, "id = ?", id).Error)
}

func (u *Release) All(ctx *gin.Context, m any, fs ...callback) ([]Release, error) {
	var list []Release
	db := database(ctx).Table(u.Table())
	db = execCallback(db, fs...)
	db = gin.GormWhere(db, u.Table(), m)
	return list, transferErr(db.Find(&list).Error)
}

func (u *Release) UpdateByID(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Updates(u).Error)
}

func (u *Release) DeleteByID(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).Where("id = ?", id).Delete(&u).Error)
}

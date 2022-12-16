package model

import (
	"github.com/limeschool/gin"
	"service/meta"
)

type ImageRegistry struct {
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	Host         string `json:"host"`
	Protocol     string `json:"protocol"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	HistoryCount int64  `json:"history_count"`
	Operator     string `json:"operator,omitempty"`
	OperatorId   int64  `json:"operator_id,omitempty"`
	gin.BaseModel
}

func (u *ImageRegistry) Table() string {
	return "image_registry"
}

func (u *ImageRegistry) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Create(&u).Error)
}

func (u *ImageRegistry) OneById(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).First(u, "id = ?", id).Error)
}

func (u *ImageRegistry) All(ctx *gin.Context, m any, fs ...callback) ([]ImageRegistry, error) {
	var list []ImageRegistry
	db := database(ctx).Table(u.Table())
	db = execCallback(db, fs...)
	db = gin.GormWhere(db, u.Table(), m)
	return list, transferErr(db.Find(&list).Error)
}

func (u *ImageRegistry) UpdateByID(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Updates(u).Error)
}

func (u *ImageRegistry) DeleteByID(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).Where("id = ?", id).Delete(&u).Error)
}

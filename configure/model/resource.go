package model

import (
	"configure/meta"
	"github.com/limeschool/gin"
)

type Resource struct {
	gin.BaseModel
	Field       string  `json:"field"`
	Type        string  `json:"type"`
	ChildField  string  `json:"child_field"`
	Description *string `json:"description"`
	Operator    string  `json:"operator"`
	OperatorId  int64   `json:"operator_id"`
	FieldValue  string  `json:"field_value,omitempty" gorm:"->"`
}

func (u *Resource) Table() string {
	return "resource"
}

func (u *Resource) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Create(&u).Error)
}

func (u *Resource) Page(ctx *gin.Context, page, count int, m interface{}) ([]Resource, int64, error) {
	var list []Resource
	var total int64
	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	if err := db.Count(&total).Error; err != nil {
		return nil, total, transferErr(database(ctx).Table(u.Table()).Create(&u).Error)
	}
	return list, total, transferErr(db.Offset((page - 1) * count).Limit(count).Find(&list).Error)
}

func (u *Resource) AllByCallback(ctx *gin.Context, fs ...callback) ([]Resource, error) {
	var list []Resource
	db := database(ctx).Table(u.Table())
	db = execCallback(db, fs...)
	return list, db.Find(&list).Error
}

func (u *Resource) UpdateByID(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Updates(u).Error)
}

func (u *Resource) DeleteByID(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).Where("id = ?", id).Delete(&u).Error)
}

package model

import (
	"configure/meta"
	"github.com/limeschool/gin"
)

type SystemField struct {
	gin.BaseModel
	Field       string  `json:"field"`
	Type        string  `json:"type"`
	ChildField  string  `json:"child_field"`
	Description *string `json:"description"`
	Operator    string  `json:"operator"`
	OperatorId  int64   `json:"operator_id"`
	FieldValue  string  `json:"field_value,omitempty" gorm:"->"`
}

func (u *SystemField) Table() string {
	return "system_field"
}

func (u *SystemField) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return database(ctx).Table(u.Table()).Create(&u).Error
}

func (u *SystemField) OneById(ctx *gin.Context, id int64) error {
	return database(ctx).Table(u.Table()).First(u, "id = ?", id).Error
}

func (u *SystemField) One(ctx *gin.Context, conds ...interface{}) error {
	return database(ctx).Table(u.Table()).First(u, conds...).Error
}

func (u *SystemField) Page(ctx *gin.Context, page, count int, m interface{}) ([]SystemField, int64, error) {
	var list []SystemField
	var total int64
	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}
	return list, total, db.Offset((page - 1) * count).Limit(count).Find(&list).Error
}

func (u *SystemField) All(ctx *gin.Context, conds ...interface{}) ([]SystemField, error) {
	var list []SystemField
	db := database(ctx).Table(u.Table())
	return list, db.Find(&list, conds...).Error
}

func (u *SystemField) AllByCallback(ctx *gin.Context, fs ...callback) ([]SystemField, error) {
	var list []SystemField
	db := database(ctx).Table(u.Table())
	for _, f := range fs {
		f(db)
	}
	return list, db.Find(&list).Error
}

func (u *SystemField) UpdateByID(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return database(ctx).Table(u.Table()).Updates(u).Error
}

func (u *SystemField) DeleteByID(ctx *gin.Context, id int64) error {
	return database(ctx).Table(u.Table()).Where("id = ?", id).Delete(&u).Error
}

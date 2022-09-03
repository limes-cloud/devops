package model

import (
	"configure/meta"
	"github.com/limeschool/gin"
)

type ServiceField struct {
	gin.BaseModel
	ServiceId   int64   `json:"service_id"`
	Field       string  `json:"field"`
	Description *string `json:"description"`
	Operator    string  `json:"operator"`
	OperatorId  int64   `json:"operator_id"`
	FieldValue  string  `json:"field_value,omitempty" gorm:"->"`
}

func (u *ServiceField) Table() string {
	return "service_field"
}

func (u *ServiceField) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return database(ctx).Table(u.Table()).Create(&u).Error
}

func (u *ServiceField) OneById(ctx *gin.Context, id int64) error {
	return database(ctx).Table(u.Table()).First(u, "id = ?", id).Error
}

func (u *ServiceField) One(ctx *gin.Context, conds ...interface{}) error {
	return database(ctx).Table(u.Table()).First(u, conds...).Error
}

func (u *ServiceField) Page(ctx *gin.Context, page, count int, m interface{}) ([]ServiceField, int64, error) {
	var list []ServiceField
	var total int64
	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}
	return list, total, db.Offset((page - 1) * count).Limit(count).Find(&list).Error
}

func (u *ServiceField) All(ctx *gin.Context, conds ...interface{}) ([]ServiceField, error) {
	var list []ServiceField
	db := database(ctx).Table(u.Table())
	return list, db.Find(&list, conds...).Error
}

func (u *ServiceField) AllByCallback(ctx *gin.Context, fs ...callback) ([]ServiceField, error) {
	var list []ServiceField
	db := database(ctx).Table(u.Table())
	for _, f := range fs {
		f(db)
	}
	return list, db.Find(&list).Error
}

func (u *ServiceField) UpdateByID(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return database(ctx).Table(u.Table()).Updates(u).Error
}

func (u *ServiceField) DeleteByID(ctx *gin.Context, id int64) error {
	return database(ctx).Table(u.Table()).Where("id = ?", id).Delete(&u).Error
}

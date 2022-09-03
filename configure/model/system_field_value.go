package model

import (
	"github.com/limeschool/gin"
	"gorm.io/gorm"
)

type SystemFieldValue struct {
	gin.BaseModel
	EnvId      int64  `json:"env_id"`
	FieldId    int64  `json:"field_id"`
	Value      string `json:"value"`
	Operator   string `json:"operator"`
	OperatorId int64  `json:"operator_id"`
}

func (u *SystemFieldValue) Table() string {
	return "system_field_value"
}

func (u *SystemFieldValue) CreateAll(ctx *gin.Context, list []SystemFieldValue) error {
	return database(ctx).Table(u.Table()).Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(u, "field_id = ?", u.FieldId).Error; err != nil {
			return err
		}
		return tx.Create(list).Error
	})
}

func (u *SystemFieldValue) All(ctx *gin.Context, conds ...interface{}) ([]SystemFieldValue, error) {
	var list []SystemFieldValue
	db := database(ctx).Table(u.Table())
	return list, db.Find(&list, conds...).Error
}

package model

import (
	"github.com/limeschool/gin"
	"gorm.io/gorm"
)

type FieldValue struct {
	gin.BaseModel
	EnvKeyword string `json:"env_keyword"`
	FieldId    int64  `json:"field_id"`
	Value      string `json:"value"`
	Operator   string `json:"operator"`
	OperatorId int64  `json:"operator_id"`
}

func (u *FieldValue) Table() string {
	return "field_value"
}

func (u *FieldValue) CreateAll(ctx *gin.Context, list []FieldValue) error {
	return database(ctx).Table(u.Table()).Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(u, "field_id = ?", u.FieldId).Error; err != nil {
			return err
		}
		return tx.Create(list).Error
	})
}

func (u *FieldValue) All(ctx *gin.Context, conds ...interface{}) ([]FieldValue, error) {
	var list []FieldValue
	db := database(ctx).Table(u.Table())
	return list, db.Find(&list, conds...).Error
}

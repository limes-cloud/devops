package model

import (
	"github.com/limeschool/gin"
	"gorm.io/gorm"
)

type ResourceValue struct {
	gin.BaseModel
	EnvKeyword string `json:"env_keyword"`
	ResourceId int64  `json:"resource_id"`
	Value      string `json:"value"`
	Operator   string `json:"operator"`
	OperatorId int64  `json:"operator_id"`
}

func (u *ResourceValue) Table() string {
	return "resource_value"
}

func (u *ResourceValue) CreateAll(ctx *gin.Context, list []ResourceValue) error {
	return database(ctx).Table(u.Table()).Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(u, "resource_id = ?", u.ResourceId).Error; err != nil {
			return err
		}
		return transferErr(tx.Create(list).Error)
	})
}

func (u *ResourceValue) All(ctx *gin.Context, cond ...interface{}) ([]ResourceValue, error) {
	var list []ResourceValue
	db := database(ctx).Table(u.Table())
	return list, transferErr(db.Find(&list, cond...).Error)
}

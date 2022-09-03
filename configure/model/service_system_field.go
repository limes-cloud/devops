package model

import (
	"github.com/limeschool/gin"
	"gorm.io/gorm"
)

type ServiceSystemField struct {
	ServiceId     int64  `json:"service_id"`
	SystemFieldId int64  `json:"system_field_id"`
	Operator      string `json:"operator"`
	OperatorId    int64  `json:"operator_id"`
}

func (u *ServiceSystemField) Table() string {
	return "service_system_field"
}

func (u *ServiceSystemField) All(ctx *gin.Context, m interface{}) ([]ServiceSystemField, error) {
	var list []ServiceSystemField
	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	return list, db.Find(&list).Error
}

func (u *ServiceSystemField) CreateAll(ctx *gin.Context, list []ServiceSystemField) error {
	return database(ctx).Table(u.Table()).Transaction(func(tx *gorm.DB) error {
		if tx.Where("service_id = ?", u.ServiceId).Delete(&u).Error != nil {
			tx.Rollback()
		}
		if len(list) == 0 {
			return nil
		}
		return tx.Create(&list).Error
	})
}

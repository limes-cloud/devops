package model

import (
	"github.com/limeschool/gin"
	"gorm.io/gorm"
)

type ServiceResource struct {
	ServiceKeyword string `json:"service_keyword"`
	ResourceID     int64  `json:"resource_id"`
	Operator       string `json:"operator"`
	OperatorID     int64  `json:"operator_id"`
}

func (u *ServiceResource) Table() string {
	return "service_resource"
}

func (u *ServiceResource) All(ctx *gin.Context, m interface{}) ([]ServiceResource, error) {
	var list []ServiceResource
	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	return list, db.Find(&list).Error
}

func (u *ServiceResource) CreateAll(ctx *gin.Context, list []ServiceResource) error {
	return database(ctx).Table(u.Table()).Transaction(func(tx *gorm.DB) error {
		if tx.Where("resource_id=?", u.ResourceID).Delete(&u).Error != nil {
			tx.Rollback()
		}
		if len(list) == 0 {
			return nil
		}
		return tx.Create(&list).Error
	})
}

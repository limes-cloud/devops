package service

import (
	"configure/model"
	"configure/types"
	"github.com/limeschool/gin"
	"gorm.io/gorm"
)

func AllServiceFieldAndResource(ctx *gin.Context, in *types.AllServiceFieldRequest) (interface{}, error) {
	// 查询业务字段
	srvField := model.Field{}
	srvFields, _ := srvField.All(ctx, "service_keyword = ?", in.Keyword)
	if srvFields == nil {
		srvFields = []model.Field{}
	}

	// 查询资源字段
	sysFiled := model.Resource{}
	sysFields, _ := sysFiled.AllByCallback(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (select resource_id from service_resource where service_keyword=?)", in.Keyword)
	})
	if sysFields == nil {
		sysFields = []model.Resource{}
	}

	// 返回数据
	return gin.H{
		"service":  srvFields,
		"resource": sysFields,
	}, nil
}

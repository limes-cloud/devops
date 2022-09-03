package model

import (
	"configure/consts"
	"github.com/limeschool/gin"
	"gorm.io/gorm"
)

type callback func(db *gorm.DB) *gorm.DB

func database(ctx *gin.Context) *gorm.DB {
	return ctx.Mysql(consts.DATABASE)
}

package model

import (
	"github.com/limeschool/gin"
	"gorm.io/gorm"
	"notice/consts"
)

type callback func(db *gorm.DB) *gorm.DB

func database(ctx *gin.Context) *gorm.DB {
	return ctx.Mysql(consts.DATABASE)
}

func execCallback(db *gorm.DB, fs ...callback) *gorm.DB {
	if fs != nil {
		for _, f := range fs {
			db = f(db)
		}
	}
	return db
}

package model

import (
	"github.com/limeschool/gin"
	"gorm.io/gorm"
	"service/consts"
	"service/errors"
	"strings"
)

var dataMap = map[string]string{
	"name": "名称",
}

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

func transferErr(err error) error {
	if err == nil {
		return nil
	}

	if customErr, ok := err.(*gin.CustomError); ok {
		return customErr
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.DBNotFoundError
	}

	if strings.Contains(err.Error(), "Duplicate") {
		str := err.Error()
		str = strings.ReplaceAll(str, "'", "")
		str = strings.TrimPrefix(str, "Error 1062: Duplicate entry ")
		arr := strings.Split(str, " for key ")
		return errors.NewF(`%v "%v" 已存在`, dataMap[arr[1]], arr[0])
	}

	if strings.Contains(err.Error(), "a foreign key constraint fails") {
		return errors.New("当前数据正在被使用")
	}
	return errors.DBError
}

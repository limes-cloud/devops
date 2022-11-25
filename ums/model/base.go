package model

import (
	"github.com/go-redis/redis/v8"
	"github.com/limeschool/gin"
	"gorm.io/gorm"
	"strings"
	"time"
	"ums/consts"
	"ums/errors"
)

var dataMap = map[string]string{
	"phone":   "手机号码",
	"email":   "电子邮箱",
	"keyword": "标志",
	"name":    "名称",
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

	return errors.DBError
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

func delayDelCache(ctx *gin.Context, key string) {
	ctx.Redis(consts.REDIS).Del(ctx, key)
	go func() {
		time.Sleep(1 * time.Second)
		ctx.Redis(consts.REDIS).Del(ctx, key)
	}()
}

func cache(ctx *gin.Context) *redis.Client {
	return ctx.Redis(consts.REDIS)
}

package models

import (
	"devops/common/drive/mysqlx"
	"devops/common/drive/redisx"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type callback func(db *gorm.DB) *gorm.DB

func database() *gorm.DB {
	return mysqlx.DB
}

func catch() *redis.Client {
	return redisx.Client
}

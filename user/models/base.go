package models

import (
	"devops/common/drive/mysqlx"
	"gorm.io/gorm"
)

func database() *gorm.DB {
	return mysqlx.DB
}

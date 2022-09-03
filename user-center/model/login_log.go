package model

import (
	"github.com/limeschool/gin"
)

type LoginLog struct {
	gin.CreateModel
	Username    string `json:"username"`
	IP          string `json:"ip"`
	Address     string `json:"address"`
	Browser     string `json:"browser"`
	Device      string `json:"device"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
}

func (u LoginLog) Table() string {
	return "login_log"
}

func (u *LoginLog) Create(ctx *gin.Context) error {
	return database(ctx).Table(u.Table()).Create(u).Error
}

func (u *LoginLog) Page(ctx *gin.Context, page, count int, m interface{}, fs ...callback) ([]LoginLog, int64, error) {
	var list []LoginLog
	var total int64
	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	db = execCallback(db, fs...)
	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}
	return list, total, db.Order("created_at desc").Offset((page - 1) * count).Limit(count).Find(&list).Error
}

func (u *LoginLog) Count(ctx *gin.Context, fs ...callback) (int64, error) {
	var total int64
	db := database(ctx).Table(u.Table())
	db = execCallback(db, fs...)
	return total, db.Count(&total).Error
}

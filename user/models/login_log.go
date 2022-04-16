package models

import (
	"devops/common/model"
)

type LoginLog struct {
	model.CreateModel
	Username    string `json:"username"`
	Password    string `json:"password"`
	IP          string `json:"ip"`
	Address     string `json:"address"`
	Browser     string `json:"browser"`
	Device      string `json:"device"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
}

func (u *LoginLog) Table() string {
	return "login_log"
}

func (u *LoginLog) OneByID() error {
	return database().Table(u.Table()).First(&u, u.ID).Error
}

func (u *LoginLog) One(query interface{}, f ...callback) error {
	db := database().Table(u.Table())
	if query != nil {
		db = model.SqlWhere(db, query)
	}
	for _, fun := range f {
		db = fun(db)
	}
	return db.First(&u).Error
}

func (u *LoginLog) Create() error {
	return database().Table(u.Table()).Create(u).Error
}

func (u *LoginLog) Page(query interface{}, page, count int64, f ...callback) ([]LoginLog, int64, error) {
	var list []LoginLog
	var total int64
	db := database().Table(u.Table())
	if query != nil {
		db = model.SqlWhere(db, query, "page", "count", "start", "end")
	}
	for _, fun := range f {
		db = fun(db)
	}
	db.Count(&total)
	db = db.Offset(int((page - 1) * count)).Limit(int(count))
	return list, total, db.Find(&list).Error
}

func (u *LoginLog) All(query interface{}, f ...callback) ([]LoginLog, int64, error) {
	var list []LoginLog
	var total int64
	db := database().Table(u.Table())
	if query != nil {
		db = model.SqlWhere(db, query)
	}
	for _, fun := range f {
		db = fun(db)
	}
	db.Count(&total)
	return list, total, db.Find(&list).Error
}

func (u *LoginLog) Count(query interface{}, f ...callback) (int64, error) {
	var total int64
	db := database().Table(u.Table())
	if query != nil {
		db = model.SqlWhere(db, query)
	}
	for _, fun := range f {
		db = fun(db)
	}
	return total, db.Count(&total).Error
}

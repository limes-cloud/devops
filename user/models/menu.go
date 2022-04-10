package models

import (
	"context"
	"devops/common/meta"
	"devops/common/model"
	"devops/common/tools"
	"time"
)

type Menu struct {
	Title      string `json:"title"`
	Icon       string `json:"icon"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Permission string `json:"permission"`
	Method     string `json:"method"`
	Component  string `json:"component"`
	ParentID   int64  `json:"parent_id"`
	Weight     int    `json:"weight"`
	Hidden     bool   `json:"hidden"`
	IsFrame    bool   `json:"is_frame"`
	Operator   string `json:"operator"`
	OperatorID int64  `json:"operator_id"`
	model.DeleteModel
}

func (u Menu) Table() string {
	return "menu"
}

func (u *Menu) OneByID() error {
	return database().Table(u.Table()).First(&u, u.ID).Error
}

func (u *Menu) One(query interface{}) error {
	db := database().Table(u.Table())
	if query != nil {
		db = model.SqlWhere(db, query)
	}
	return db.First(&u).Error
}

func (u *Menu) Create(ctx context.Context) error {
	u.OperatorID = meta.UserId(ctx)
	u.Operator = meta.UserName(ctx)
	return database().Table(u.Table()).Create(&u).Error
}

func (u *Menu) Page(query interface{}, page, count int64) ([]Menu, int64, error) {
	var list []Menu
	var total int64
	db := database().Table(u.Table())
	if query != nil {
		db = model.SqlWhere(db, query, "page", "count")
	}
	db.Count(&total)
	db = db.Order(u.Table() + ".weight desc")
	db = db.Offset(int((page - 1) * count)).Limit(int(count))
	return list, total, db.Find(&list).Error
}

func (u *Menu) All(query interface{}) ([]Menu, int64, error) {
	var list []Menu
	var total int64
	db := database().Table(u.Table())
	if query != nil {
		db = model.SqlWhere(db, query)
	}
	db.Count(&total)
	db = db.Order(u.Table() + ".weight desc")
	return list, total, db.Find(&list).Error
}

func (u *Menu) AllCall(f callback) ([]Menu, int64, error) {
	var list []Menu
	var total int64
	db := database().Table(u.Table())
	if f != nil {
		db = f(db)
	}
	db = db.Order(u.Table() + ".weight desc")
	db.Count(&total)
	return list, total, db.Find(&list).Error
}

func (u *Menu) UpdateByFields(ctx context.Context, c interface{}, m interface{}) error {
	fields := tools.ToMap(m)
	db := database().Table(u.Table())
	db = model.SqlWhere(db, c)
	fields["created_at"] = time.Now().Unix()
	fields["updated_at"] = time.Now().Unix()
	fields["operator"] = meta.UserName(ctx)
	fields["operator_id"] = meta.UserId(ctx)
	return db.Updates(fields).Error
}

func (u *Menu) UpdateByID(ctx context.Context, m interface{}) error {
	fields := tools.ToMap(m)
	fields["created_at"] = time.Now().Unix()
	fields["updated_at"] = time.Now().Unix()
	fields["operator"] = meta.UserName(ctx)
	fields["operator_id"] = meta.UserId(ctx)
	return database().Table(u.Table()).Where("id = ?", u.ID).Updates(fields).Error
}

func (u *Menu) Update(ctx context.Context) error {
	u.OperatorID = meta.UserId(ctx)
	u.Operator = meta.UserName(ctx)
	return database().Table(u.Table()).Updates(u).Error
}

func (u *Menu) DeleteByID(ctx context.Context) error {
	u.OperatorID = meta.UserId(ctx)
	u.Operator = meta.UserName(ctx)
	return database().Table(u.Table()).Delete(&u).Error
}

func (u *Menu) Delete(ctx context.Context, m interface{}) error {
	u.OperatorID = meta.UserId(ctx)
	u.Operator = meta.UserName(ctx)
	db := database().Table(u.Table())
	db = model.SqlWhere(db, m)
	return db.Delete(&u).Error
}

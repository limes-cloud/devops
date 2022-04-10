package models

import (
	"context"
	"devops/common/meta"
	"devops/common/model"
	"devops/common/tools"
	"time"
)

type Role struct {
	model.BaseModel
	Name        string `json:"name" `
	Keyword     string `json:"keyword"`
	Status      *bool  `json:"status" `
	Weight      int    `json:"weight"`
	Description string `json:"description"`
	DataScope   string `json:"data_scope" `
	Operator    string `json:"operator"`
	OperatorID  int64  `json:"operator_id"`
}

func (u *Role) Table() string {
	return "role"
}

func (u *Role) OneByID() error {
	return database().Table(u.Table()).First(&u, u.ID).Error
}

func (u *Role) One(query interface{}) error {
	db := database().Table(u.Table())
	if query != nil {
		db = model.SqlWhere(db, query)
	}
	return db.First(&u).Error
}

func (u *Role) Create(ctx context.Context) error {
	u.OperatorID = meta.UserId(ctx)
	u.Operator = meta.UserName(ctx)
	return database().Table(u.Table()).Create(&u).Error
}

func (u *Role) Page(query interface{}, page, count int64) ([]Role, int64, error) {
	var list []Role
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

func (u *Role) All(query interface{}) ([]Role, int64, error) {
	var list []Role
	var total int64
	db := database().Table(u.Table())
	if query != nil {
		db = model.SqlWhere(db, query)
	}
	db.Count(&total)
	db = db.Order(u.Table() + ".weight desc")
	return list, total, db.Find(&list).Error
}

func (u *Role) UpdateByFields(ctx context.Context, c interface{}, m interface{}) error {
	fields := tools.ToMap(m)
	db := database().Table(u.Table())
	db = model.SqlWhere(db, c)
	fields["created_at"] = time.Now().Unix()
	fields["updated_at"] = time.Now().Unix()
	fields["operator"] = meta.UserName(ctx)
	fields["operator_id"] = meta.UserId(ctx)
	return db.Updates(fields).Error
}

func (u *Role) UpdateByID(ctx context.Context, m interface{}) error {
	fields := tools.ToMap(m)
	fields["created_at"] = time.Now().Unix()
	fields["updated_at"] = time.Now().Unix()
	fields["operator"] = meta.UserName(ctx)
	fields["operator_id"] = meta.UserId(ctx)
	return database().Table(u.Table()).Where("id = ?", u.ID).Updates(fields).Error
}

func (u *Role) Update(ctx context.Context) error {
	u.OperatorID = meta.UserId(ctx)
	u.Operator = meta.UserName(ctx)
	return database().Table(u.Table()).Updates(u).Error
}

func (u *Role) DeleteByID(ctx context.Context) error {
	u.OperatorID = meta.UserId(ctx)
	u.Operator = meta.UserName(ctx)
	return database().Table(u.Table()).Delete(&u).Error
}

func (u *Role) Delete(ctx context.Context, m interface{}) error {
	u.OperatorID = meta.UserId(ctx)
	u.Operator = meta.UserName(ctx)
	db := database().Table(u.Table())
	db = model.SqlWhere(db, m)
	return db.Delete(&u).Error
}

package models

import (
	"context"
	"devops/common/meta"
	"devops/common/model"
	"devops/common/tools"
	"time"
)

type Configure struct {
	ServiceId   int64  `json:"service_id"`
	Template    string `json:"template"`
	Version     string `json:"version"`
	IsUse       bool   `json:"is_use"`
	Description string `json:"description"`
	Operator    string `json:"operator"`
	OperatorID  int64  `json:"operator_id"`
	model.BaseModel
}

func (e Configure) Table() string {
	return "configure"
}

func (e *Configure) OneByID() error {
	return database().Table(e.Table()).First(&e, e.ID).Error
}

func (e *Configure) One(query interface{}, f ...callback) error {
	db := database().Table(e.Table())
	db = model.SqlWhere(db, query)
	for _, fun := range f {
		db = fun(db)
	}
	return db.First(&e).Error
}

func (e *Configure) Create(ctx context.Context) error {
	e.OperatorID = meta.UserId(ctx)
	e.Operator = meta.UserName(ctx)
	return database().Table(e.Table()).Create(&e).Error
}

func (e *Configure) Page(query interface{}, page, count int64, f ...callback) ([]Configure, int64, error) {
	var list []Configure
	var total int64
	db := database().Table(e.Table())
	db = model.SqlWhere(db, query, "page", "count")

	for _, fun := range f {
		db = fun(db)
	}

	db.Count(&total)
	db = db.Offset(int((page - 1) * count)).Limit(int(count))
	return list, total, db.Find(&list).Error
}

func (e *Configure) All(query interface{}, f ...callback) ([]Configure, int64, error) {
	var list []Configure
	var total int64

	db := database().Table(e.Table()).
		Select("id,service_id,version, is_use, description, operator, operator_id,created_at")
	db = model.SqlWhere(db, query)

	for _, fun := range f {
		db = fun(db)
	}

	db.Count(&total)
	return list, total, db.Find(&list).Error
}

func (e *Configure) Update(ctx context.Context, c interface{}, m interface{}, f ...callback) error {
	fields := tools.ToMap(m)
	db := database().Table(e.Table())
	db = model.SqlWhere(db, c)
	for _, fun := range f {
		db = fun(db)
	}
	fields["created_at"] = time.Now().Unix()
	fields["updated_at"] = time.Now().Unix()
	fields["operator"] = meta.UserName(ctx)
	fields["operator_id"] = meta.UserId(ctx)
	return db.Updates(m).Error
}

func (e *Configure) UpdateByID(ctx context.Context, m interface{}) error {
	fields := tools.ToMap(m)
	fields["created_at"] = time.Now().Unix()
	fields["updated_at"] = time.Now().Unix()
	fields["operator"] = meta.UserName(ctx)
	fields["operator_id"] = meta.UserId(ctx)
	return database().Table(e.Table()).Where("id = ?", e.ID).Updates(m).Error
}

func (e *Configure) DeleteByID(ctx context.Context) error {
	e.OperatorID = meta.UserId(ctx)
	e.Operator = meta.UserName(ctx)
	return database().Table(e.Table()).Delete(&e).Error
}

func (e *Configure) Delete(ctx context.Context, m interface{}, f ...callback) error {
	db := database().Table(e.Table())
	db = model.SqlWhere(db, m)
	for _, fun := range f {
		db = fun(db)
	}

	return db.Delete(e).Error
}

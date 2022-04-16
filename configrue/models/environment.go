package models

import (
	"context"
	"devops/common/meta"
	"devops/common/model"
	"devops/common/tools"
	"time"
)

type Environment struct {
	Keyword     string `json:"keyword"`
	Name        string `json:"name"`
	Config      string `json:"config"`
	Prefix      string `json:"prefix"`
	Description string `json:"description"`
	Operator    string `json:"operator"`
	OperatorID  int64  `json:"operator_id"`
	model.DeleteModel
}

func (e Environment) Table() string {
	return "environment"
}

func (e *Environment) OneByID() error {
	return database().Table(e.Table()).First(&e, e.ID).Error
}

func (e *Environment) One(query interface{}, f ...callback) error {
	db := database().Table(e.Table())
	db = model.SqlWhere(db, query)
	for _, fun := range f {
		db = fun(db)
	}
	return db.First(&e).Error
}

func (e *Environment) Create(ctx context.Context) error {
	e.OperatorID = meta.UserId(ctx)
	e.Operator = meta.UserName(ctx)
	return database().Table(e.Table()).Create(&e).Error
}

func (e *Environment) Page(query interface{}, page, count int64, f ...callback) ([]Environment, int64, error) {
	var list []Environment
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

func (e *Environment) All(query interface{}, f ...callback) ([]Environment, int64, error) {
	var list []Environment
	var total int64
	db := database().Table(e.Table())
	db = model.SqlWhere(db, query)

	for _, fun := range f {
		db = fun(db)
	}

	db.Count(&total)
	return list, total, db.Find(&list).Error
}

func (e *Environment) Update(ctx context.Context, c interface{}, m interface{}, f ...callback) error {
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

func (e *Environment) UpdateByID(ctx context.Context, m interface{}) error {
	fields := tools.ToMap(m)
	fields["created_at"] = time.Now().Unix()
	fields["updated_at"] = time.Now().Unix()
	fields["operator"] = meta.UserName(ctx)
	fields["operator_id"] = meta.UserId(ctx)
	return database().Table(e.Table()).Where("id = ?", e.ID).Updates(m).Error
}

func (e *Environment) DeleteByID(ctx context.Context) error {
	e.OperatorID = meta.UserId(ctx)
	e.Operator = meta.UserName(ctx)
	return database().Table(e.Table()).Delete(&e).Error
}

func (e *Environment) Delete(ctx context.Context, m interface{}, f ...callback) error {
	e.OperatorID = meta.UserId(ctx)
	e.Operator = meta.UserName(ctx)
	db := database().Table(e.Table())
	db = model.SqlWhere(db, m)
	for _, fun := range f {
		db = fun(db)
	}
	return db.Delete(&e).Error
}

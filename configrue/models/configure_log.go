package models

import (
	"context"
	"devops/common/meta"
	"devops/common/model"
	"devops/common/tools"
	"time"
)

type ConfigureLog struct {
	ServiceName string `json:"service_name"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Operator    string `json:"operator"`
	OperatorID  int64  `json:"operator_id"`
	model.CreateModel
}

func (e ConfigureLog) Table() string {
	return "configure_log"
}

func (e *ConfigureLog) OneByID() error {
	return database().Table(e.Table()).First(&e, e.ID).Error
}

func (e *ConfigureLog) One(query interface{}, f ...callback) error {
	db := database().Table(e.Table())
	db = model.SqlWhere(db, query)
	for _, fun := range f {
		db = fun(db)
	}
	return db.First(&e).Error
}

func (e *ConfigureLog) Create(ctx context.Context) error {
	e.OperatorID = meta.UserId(ctx)
	e.Operator = meta.UserName(ctx)
	return database().Table(e.Table()).Create(&e).Error
}

func (e *ConfigureLog) Page(query interface{}, page, count int64, f ...callback) ([]ConfigureLog, int64, error) {
	var list []ConfigureLog
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

func (e *ConfigureLog) All(query interface{}, f ...callback) ([]ConfigureLog, int64, error) {
	var list []ConfigureLog
	var total int64
	db := database().Table(e.Table())
	db = model.SqlWhere(db, query)

	for _, fun := range f {
		db = fun(db)
	}

	db.Count(&total)
	return list, total, db.Find(&list).Error
}

func (e *ConfigureLog) Update(ctx context.Context, c interface{}, m interface{}, f ...callback) error {
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

func (e *ConfigureLog) UpdateByID(ctx context.Context, m interface{}) error {
	fields := tools.ToMap(m)
	fields["created_at"] = time.Now().Unix()
	fields["updated_at"] = time.Now().Unix()
	fields["operator"] = meta.UserName(ctx)
	fields["operator_id"] = meta.UserId(ctx)
	return database().Table(e.Table()).Where("id = ?", e.ID).Updates(m).Error
}

func (e *ConfigureLog) DeleteByID(ctx context.Context) error {
	e.OperatorID = meta.UserId(ctx)
	e.Operator = meta.UserName(ctx)
	return database().Table(e.Table()).Delete(&e).Error
}

func (e *ConfigureLog) Delete(ctx context.Context, m interface{}, f ...callback) error {
	db := database().Table(e.Table())
	db = model.SqlWhere(db, m)
	for _, fun := range f {
		db = fun(db)
	}

	return db.Delete(e).Error
}

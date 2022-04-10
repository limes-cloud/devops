package models

import (
	"context"
	"devops/common/meta"
	"devops/common/model"
	"devops/common/tools"
	"time"
)

type Team struct {
	model.BaseModel
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	ParentID    int64  `json:"parent_id"`
	Operator    string `json:"operator"`
	OperatorID  int64  `json:"operator_id"`
}

func (u *Team) Table() string {
	return "team"
}

func (u *Team) OneByID(userId int64) error {
	return database().Table(u.Table()).First(&u, userId).Error
}

func (u *Team) One(query interface{}) error {
	db := database().Table(u.Table())
	if query != nil {
		db = model.SqlWhere(db, query)
	}
	return db.First(&u).Error
}

func (u *Team) Create(ctx context.Context) error {
	u.OperatorID = meta.UserId(ctx)
	u.Operator = meta.UserName(ctx)
	return database().Table(u.Table()).Create(&u).Error
}

func (u *Team) Page(query interface{}, page, count int64) ([]Team, int64, error) {
	var list []Team
	var total int64
	db := database().Table(u.Table())
	if query != nil {
		db = model.SqlWhere(db, query, "page", "count")
	}
	db.Count(&total)
	db = db.Offset(int((page - 1) * count)).Limit(int(count))
	return list, total, db.Find(&list).Error
}

func (u *Team) All(query interface{}) ([]Team, int64, error) {
	var list []Team
	var total int64
	db := database().Table(u.Table())
	if query != nil {
		db = model.SqlWhere(db, query)
	}
	db.Count(&total)
	return list, total, db.Find(&list).Error
}

func (u *Team) UpdateByFields(ctx context.Context, c interface{}, m interface{}) error {
	fields := tools.ToMap(m)
	db := database().Table(u.Table())
	db = model.SqlWhere(db, c)
	fields["created_at"] = time.Now().Unix()
	fields["updated_at"] = time.Now().Unix()
	fields["operator"] = meta.UserName(ctx)
	fields["operator_id"] = meta.UserId(ctx)
	return db.Updates(fields).Error
}

func (u *Team) UpdateByID(ctx context.Context, m interface{}) error {
	fields := tools.ToMap(m)
	fields["created_at"] = time.Now().Unix()
	fields["updated_at"] = time.Now().Unix()
	fields["operator"] = meta.UserName(ctx)
	fields["operator_id"] = meta.UserId(ctx)
	return database().Table(u.Table()).Where("id = ?", u.ID).Updates(fields).Error
}

func (u *Team) Update(ctx context.Context) error {
	u.OperatorID = meta.UserId(ctx)
	u.Operator = meta.UserName(ctx)
	return database().Table(u.Table()).Updates(u).Error
}

func (u *Team) DeleteByID(ctx context.Context) error {
	u.OperatorID = meta.UserId(ctx)
	u.Operator = meta.UserName(ctx)
	return database().Table(u.Table()).Delete(&u).Error
}

func (u *Team) Delete(ctx context.Context, m interface{}) error {
	u.OperatorID = meta.UserId(ctx)
	u.Operator = meta.UserName(ctx)
	db := database().Table(u.Table())
	db = model.SqlWhere(db, m)
	return db.Delete(&u).Error
}

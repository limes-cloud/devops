package models

import (
	"context"
	"devops/common/meta"
	"devops/common/model"
	"time"
)

type RoleMenu struct {
	model.BaseModel
	RoleID     int64  `json:"role_id"`
	MenuID     int64  `json:"menu_id"`
	Operator   string `json:"operator"`
	OperatorID int64  `json:"operator_id"`
}

func (u RoleMenu) Table() string {
	return "role_menu"
}

type RoleMenus []RoleMenu

func (u *RoleMenu) Create(ctx context.Context) error {
	u.OperatorID = meta.UserId(ctx)
	u.Operator = meta.UserName(ctx)
	return database().Table(u.Table()).Create(&u).Error
}

func (u RoleMenus) Create(ctx context.Context) error {
	userId := meta.UserId(ctx)
	userName := meta.UserName(ctx)
	for _, item := range u {
		item.OperatorID = userId
		item.Operator = userName
		item.CreatedAt = time.Now().Unix()
		item.UpdatedAt = time.Now().Unix()
	}
	return database().Table(RoleMenu{}.Table()).Create(&u).Error
}

func (u *RoleMenu) All(query interface{}) ([]RoleMenu, int64, error) {
	var list []RoleMenu
	var total int64
	db := database().Table(u.Table())
	db = model.SqlWhere(db, query)
	db.Count(&total)
	return list, total, db.Find(&list).Error
}

func (u *RoleMenu) Delete(ctx context.Context, m interface{}) error {
	u.OperatorID = meta.UserId(ctx)
	u.Operator = meta.UserName(ctx)
	db := database().Table(u.Table())
	if m != nil {
		db = model.SqlWhere(db, m)
	}
	return db.Delete(&u).Error
}

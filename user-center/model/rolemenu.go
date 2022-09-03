package model

import (
	"github.com/limeschool/gin"
	"gorm.io/gorm"
)

type RoleMenu struct {
	gin.BaseModel
	RoleID     int64  `json:"role_id"`
	MenuID     int64  `json:"menu_id"`
	Operator   string `json:"operator"`
	OperatorID int64  `json:"operator_id"`
}

func (RoleMenu) Table() string {
	return "role_menu"
}

func (u *RoleMenu) Create(ctx *gin.Context, roleId int64, menuIds []int64) error {
	user := CurUser(ctx)
	var list []RoleMenu
	for _, menuId := range menuIds {
		list = append(list, RoleMenu{
			RoleID:     roleId,
			MenuID:     menuId,
			OperatorID: user.BaseModel.ID,
			Operator:   user.Name,
		})
	}
	return database(ctx).Table(u.Table()).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_id = ?", roleId).Error; err != nil {
			return err
		}
		return tx.Create(&list).Error
	})
}

func (u *RoleMenu) All(ctx *gin.Context, conds ...interface{}) ([]RoleMenu, error) {
	var list []RoleMenu
	db := database(ctx).Table(u.Table())
	return list, db.Find(&list, conds...).Error
}

func (u *RoleMenu) Delete(ctx *gin.Context, conds ...interface{}) error {
	return database(ctx).Table(u.Table()).Delete(u, conds...).Error
}

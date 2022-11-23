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

// Create 新建角色所属菜单
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
		if err := tx.Where("role_id=?", roleId).Delete(u).Error; err != nil {
			return transferErr(err)
		}
		return transferErr(tx.Create(&list).Error)
	})
}

// RoleMenus 通过角色ID获取角色菜单
func (u *RoleMenu) RoleMenus(ctx *gin.Context, roleId int64) ([]RoleMenu, error) {
	var list []RoleMenu
	db := database(ctx).Table(u.Table())
	return list, transferErr(db.Find(&list, "role_id=?", roleId).Error)
}

// Delete 通过角色id删除 角色所属菜单
func (u *RoleMenu) Delete(ctx *gin.Context, roleId int64) error {
	return transferErr(database(ctx).Table(u.Table()).Delete(u, "role_id=?", roleId).Error)
}

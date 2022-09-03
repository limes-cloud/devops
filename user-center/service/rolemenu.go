package service

import (
	"github.com/limeschool/gin"
	"ums/errors"
	"ums/meta"
	"ums/model"
	"ums/tools/tree"
	"ums/types"
)

func AddRoleMenu(ctx *gin.Context, in *types.AddRoleMenuRequest) error {
	if in.RoleID == 1 {
		return errors.SuperAdminEditError
	}
	user := model.CurUser(ctx)
	rm := model.RoleMenu{}
	// 查询这个角色下的所有菜单
	list, _ := rm.All(ctx, "role_id = ?", in.RoleID)
	var menuIds []int64
	for _, item := range list {
		menuIds = append(menuIds, item.MenuID)
	}
	// 查询历史菜单详细数据
	menu := model.Menu{}
	var policies [][]string
	menuList, _ := menu.All(ctx, "id in ? and type = 'A'", menuIds)
	for _, item := range menuList {
		policies = append(policies, []string{user.Role.Keyword, item.Path, item.Method})
	}
	// 删除所有的此用户的当前所有权限
	ctx.Rbac().RemovePolicies(policies)
	_ = rm.Delete(ctx, "role_id = ?", in.RoleID)

	// 获取当前修改菜单的信息
	policies = [][]string{}
	menuList, _ = menu.All(ctx, "id in ? and type = 'A'", in.MenuIds)
	for _, item := range menuList {
		policies = append(policies, []string{user.Role.Keyword, item.Path, item.Method})
	}
	ctx.Rbac().AddPolicies(policies)
	return rm.Create(ctx, in.RoleID, in.MenuIds)
}

func RoleMenu(ctx *gin.Context) (tree.Tree, error) {
	user := model.CurUser(ctx)
	rm := model.RoleMenu{}
	if user.Role.Keyword == meta.SuperAdmin {
		return AllMenu(ctx, &types.AllMenuRequest{})
	}
	rmList, err := rm.All(ctx, "role_id = ?", user.RoleID)
	if err != nil {
		return nil, err
	}
	var ids []int64
	for _, item := range rmList {
		ids = append(ids, item.MenuID)
	}
	var menu model.Menu
	menuList, err := menu.All(ctx, "id in ?", ids)
	if err != nil {
		return nil, err
	}
	var listTree []tree.Tree
	for _, item := range menuList {
		listTree = append(listTree, item)
	}
	return tree.BuildTree(listTree), nil
}

func RoleMenuIds(ctx *gin.Context, in *types.RoleMenuIdsRequest) ([]int64, error) {
	rm := model.RoleMenu{}
	rmList, err := rm.All(ctx, "role_id = ?", in.RoleID)
	if err != nil {
		return nil, err
	}
	var ids []int64
	for _, item := range rmList {
		ids = append(ids, item.MenuID)
	}
	return ids, nil
}

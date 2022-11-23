package service

import (
	"github.com/limeschool/gin"
	"ums/errors"
	"ums/meta"
	"ums/model"
	"ums/tools/tree"
	"ums/types"
)

// AddRoleMenu 修改角色所属菜单
func AddRoleMenu(ctx *gin.Context, in *types.AddRoleMenuRequest) error {
	// 菜单不允许修改
	if in.RoleID == 1 {
		return errors.SuperAdminEditError
	}

	// 获取当前role的数据
	role := model.Role{}
	if err := role.OneById(ctx, in.RoleID); err != nil {
		return err
	}

	// 查询这个角色当前的所有菜单id
	rm := model.RoleMenu{}
	list, _ := rm.RoleMenus(ctx, in.RoleID)
	var menuIds []int64
	for _, item := range list {
		menuIds = append(menuIds, item.MenuID)
	}

	// 查询所有所有的api数据
	menu := model.Menu{}
	var policies [][]string
	menuList, _ := menu.All(ctx, "id in ? and type = 'A'", menuIds)
	for _, item := range menuList {
		policies = append(policies, []string{role.Keyword, item.Path, item.Method})
	}

	// 删除所有的此用户的当前api所有权限
	_, _ = ctx.Rbac().RemovePolicies(policies)
	_ = rm.Delete(ctx, in.RoleID)

	// 获取当前修改菜单的信息
	policies = [][]string{}
	menuList, _ = menu.All(ctx, "id in ? and type = 'A'", in.MenuIds)
	for _, item := range menuList {
		policies = append(policies, []string{role.Keyword, item.Path, item.Method})
	}

	// 将新的策略的策略写入rbac
	_, _ = ctx.Rbac().AddPolicies(policies)

	return rm.Create(ctx, in.RoleID, in.MenuIds)
}

func RoleMenu(ctx *gin.Context) (tree.Tree, error) {
	user := model.CurUser(ctx)

	// 如果是超级管理员就直接返回全部菜单
	rm := model.RoleMenu{}
	if user.Role.Keyword == meta.SuperAdmin {
		return AllMenu(ctx, &types.AllMenuRequest{})
	}

	// 查询角色所属菜单
	rmList, err := rm.RoleMenus(ctx, user.RoleID)
	if err != nil {
		return nil, err
	}

	// 获取菜单的所有id
	var ids []int64
	for _, item := range rmList {
		ids = append(ids, item.MenuID)
	}

	// 获取指定id的所有菜单
	var menu model.Menu
	menuList, err := menu.All(ctx, "id in ?", ids)
	if err != nil {
		return nil, err
	}

	// 通过菜单列表构建菜单树
	var listTree []tree.Tree
	for _, item := range menuList {
		listTree = append(listTree, item)
	}
	return tree.BuildTree(listTree), nil
}

// RoleMenuIds 获取角色菜单的所有id
func RoleMenuIds(ctx *gin.Context, in *types.RoleMenuIdsRequest) ([]int64, error) {

	// 获取当前角色的所有菜单
	rm := model.RoleMenu{}
	rmList, err := rm.RoleMenus(ctx, in.RoleID)
	if err != nil {
		return nil, err
	}

	// 组装所有的菜单id
	var ids []int64
	for _, item := range rmList {
		ids = append(ids, item.MenuID)
	}

	return ids, nil
}

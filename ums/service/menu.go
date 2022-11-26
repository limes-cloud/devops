package service

import (
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
	"ums/consts"
	"ums/errors"
	"ums/model"
	"ums/tools/tree"
	"ums/types"
)

// AllMenu 获取菜单树
func AllMenu(ctx *gin.Context, in *types.AllMenuRequest) (tree.Tree, error) {
	menu := model.Menu{}
	if in.IsFilter != nil && *in.IsFilter {
		return menu.Tree(ctx, "permission != ?", consts.BaseApi)
	} else {
		return menu.Tree(ctx)
	}
}

// AddMenu 新增菜单
func AddMenu(ctx *gin.Context, in *types.AddMenuRequest) error {
	menu := model.Menu{}
	if copier.Copy(&menu, in) != nil {
		return errors.AssignError
	}

	if in.Name != "" && menu.One(ctx, "id!=? and name=?", menu.ID(), in.Name) == nil {
		return errors.DulMenuNameError
	}

	return menu.Create(ctx)
}

func UpdateMenu(ctx *gin.Context, in *types.UpdateMenuRequest) error {
	menu := model.Menu{}
	if menu.OneById(ctx, in.ID) != nil {
		return errors.DBNotFoundError
	}

	if in.ParentID != 0 && in.ID == in.ParentID {
		return errors.MenuParentIdError
	}

	if menu.Name != in.Name && menu.One(ctx, "id != ? and name=?", menu.ID(), in.Name) == nil {
		return errors.DulMenuNameError

	}

	// 删除rbac权限
	if menu.Type == "A" && in.Type != "A" {
		ctx.Rbac().RemoveFilteredPolicy(1, menu.Path, menu.Method)
	}

	// 更新rbac权限
	if menu.Type == "A" && in.Type == "A" {
		if menu.Method != in.Method || menu.Path != in.Path {
			oldPolices := ctx.Rbac().GetFilteredPolicy(1, menu.Path, menu.Method)
			if len(oldPolices) != 0 {
				var newPolices [][]string
				for _, val := range oldPolices {
					newPolices = append(newPolices, []string{val[0], in.Path, in.Method})
				}
				ctx.Rbac().UpdatePolicies(oldPolices, newPolices)
			}
		}
	}

	// 新增rbac权限
	if menu.Type != "A" && in.Type == "A" {
		roles := ctx.Rbac().GetAllSubjects()
		var newPolices [][]string

		for _, val := range roles {
			newPolices = append(newPolices, []string{val, in.Path, in.Method})
		}
		ctx.Rbac().AddPolicies(newPolices)
	}

	inMenu := model.Menu{}
	if copier.Copy(&inMenu, in) != nil {
		return errors.AssignError
	}

	return inMenu.UpdateByID(ctx)
}

func DeleteMenu(ctx *gin.Context, in *types.DeleteMenuRequest) error {
	menu := model.Menu{}
	list, _ := menu.All(ctx)
	var treeList []tree.Tree
	for _, item := range list {
		treeList = append(treeList, item)
	}

	// 获取指定id下的所有id
	t := tree.BuildTreeByID(treeList, in.ID)
	ids := tree.GetTreeID(t)

	// 获取当前id中的菜单
	apiList, _ := menu.All(ctx, "id in ? and type='A'", ids)
	for _, item := range apiList {
		ctx.Rbac().RemoveFilteredPolicy(1, item.Path, item.Method)
	}

	return menu.Delete(ctx, "id in ?", ids)
}

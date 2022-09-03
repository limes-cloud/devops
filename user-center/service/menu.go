package service

import (
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
	"gorm.io/gorm"
	"ums/consts"
	"ums/errors"
	"ums/model"
	"ums/tools"
	"ums/tools/tree"
	"ums/types"
)

func AllMenu(ctx *gin.Context, in *types.AllMenuRequest) (tree.Tree, error) {
	menu := model.Menu{}
	if in.IsFilter != nil && *in.IsFilter {
		return menu.Tree(ctx, "permission != ?", consts.BaseApi)
	} else {
		return menu.Tree(ctx)
	}
}

func AddMenu(ctx *gin.Context, in *types.AddMenuRequest) error {
	menu := model.Menu{}
	if copier.Copy(&menu, in) != nil {
		return errors.AssignError
	}

	if in.Name != "" && !errors.Is(menu.One(ctx, "id != ? and name=?", menu.BaseModel.ID, in.Name), gorm.ErrRecordNotFound) {
		return errors.DulMenuNameError
	}

	if in.Permission == consts.BaseApi { //如果是白名单接口，则删除缓存
		ctx.Redis(consts.REDIS).Del(ctx, consts.RedisBaseApi)
		defer func() {
			tools.DelRedis(ctx, consts.RedisBaseApi)
		}()
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

	if copier.Copy(&menu, in) != nil {
		return errors.AssignError
	}

	if menu.Name != in.Name {
		if !errors.Is(menu.One(ctx, "id != ? and name=?", menu.ID, in.Name), gorm.ErrRecordNotFound) {
			return errors.DulMenuNameError
		}
	}

	if menu.Type == "A" && in.Type != "A" {
		ctx.Rbac().RemoveFilteredPolicy(1, menu.Method, menu.Path)
	}
	if menu.Type == "A" && in.Type == "A" {
		if menu.Method != in.Method || menu.Path != in.Path {
			oldPolicys := ctx.Rbac().GetFilteredPolicy(1, menu.Method, menu.Path)
			newPolicys := oldPolicys
			for key, val := range newPolicys {
				if len(val) < 3 {
					continue
				}
				newPolicys[key][1] = in.Method
				newPolicys[key][2] = in.Path
			}
			ctx.Rbac().UpdatePolicies(oldPolicys, newPolicys)
		}
	}
	return menu.UpdateByID(ctx)
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
		ctx.Rbac().RemoveFilteredPolicy(1, item.Method, item.Path)
	}
	return menu.Delete(ctx, "id in ?", ids)
}

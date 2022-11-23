package service

import (
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
	"ums/consts"
	"ums/errors"
	"ums/model"
	"ums/types"
)

func RoleDataScope(ctx *gin.Context) (map[string]string, error) {
	return map[string]string{
		consts.ALLTEAM:  "当前部门及部门以下",
		consts.CURTEAM:  "当前部门权限",
		consts.DOWNTEAM: "部门以下权限",
		consts.CUSTOM:   "自定义权限",
	}, nil
}

func AllRole(ctx *gin.Context, in *types.AllRoleRequest) ([]model.Role, error) {
	role := model.Role{}
	return role.All(ctx, in)
}

func AddRole(ctx *gin.Context, in *types.AddRoleRequest) error {
	role := model.Role{}
	if copier.Copy(&role, in) != nil {
		return errors.AssignError
	}
	return role.Create(ctx)
}

func UpdateRole(ctx *gin.Context, in *types.UpdateRoleRequest) error {
	if in.ID == 1 {
		return errors.SuperAdminEditError
	}
	role := model.Role{}
	if copier.Copy(&role, in) != nil {
		return errors.AssignError
	}
	return role.UpdateByID(ctx)
}

func DeleteRole(ctx *gin.Context, in *types.DeleteRoleRequest) error {
	if in.ID == 1 {
		return errors.SuperAdminDelError
	}
	role := model.Role{}
	if err := role.OneById(ctx, in.ID); err != nil {
		return err
	}
	ctx.Rbac().RemoveFilteredPolicy(0, role.Keyword)
	return role.DeleteByID(ctx, in.ID)
}

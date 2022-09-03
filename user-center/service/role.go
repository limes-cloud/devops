package service

import (
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
	"ums/consts"
	"ums/errors"
	"ums/model"
	"ums/tools"
	"ums/types"
)

func RoleDataScope(ctx *gin.Context) (map[string]string, error) {
	return map[string]string{
		consts.ALLTEAM:  "所有权限",
		consts.DOWNTEAM: "部门以下权限",
		consts.CURTEAM:  "当前部门权限",
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
	if tools.DataDup(role.Create(ctx)) {
		return errors.DulKeywordError
	}
	return nil
}

func UpdateRole(ctx *gin.Context, in *types.UpdateRoleRequest) error {
	if in.ID == 1 {
		return errors.SuperAdminEditError
	}
	role := model.Role{}
	if copier.Copy(&role, in) != nil {
		return errors.AssignError
	}
	if tools.DataDup(role.UpdateByID(ctx)) {
		return errors.DulKeywordError
	}
	return nil
}

func DeleteRole(ctx *gin.Context, in *types.DeleteRoleRequest) error {
	if in.ID == 1 {
		return errors.SuperAdminDelError
	}
	role := model.Role{}
	if role.OneById(ctx, in.ID) != nil {
		return errors.DBNotFoundError
	}
	ctx.Rbac().RemoveFilteredPolicy(0, role.Keyword)
	return role.DeleteByID(ctx, in.ID)
}

package handler

import (
	"github.com/limeschool/gin"
	"ums/errors"
	"ums/service"
	"ums/types"
)

func AllRole(ctx *gin.Context) {
	in := types.AllRoleRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if resp, err := service.AllRole(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func RoleDataScope(ctx *gin.Context) {
	if resp, err := service.RoleDataScope(ctx); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func AddRole(ctx *gin.Context) {
	in := types.AddRoleRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddRole(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateRole(ctx *gin.Context) {
	in := types.UpdateRoleRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.UpdateRole(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteRole(ctx *gin.Context) {
	in := types.DeleteRoleRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.DeleteRole(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

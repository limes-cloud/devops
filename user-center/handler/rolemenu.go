package handler

import (
	"github.com/limeschool/gin"
	"ums/errors"
	"ums/service"
	"ums/types"
)

func AddRoleMenu(ctx *gin.Context) {
	in := types.AddRoleMenuRequest{}
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddRoleMenu(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func RoleMenu(ctx *gin.Context) {
	if tree, err := service.RoleMenu(ctx); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(tree)
	}
}

func RoleMenuIds(ctx *gin.Context) {
	in := types.RoleMenuIdsRequest{}
	if err := ctx.ShouldBind(&in); err != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if ids, err := service.RoleMenuIds(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(ids)
	}
}

package handler

import (
	"github.com/limeschool/gin"
	"ums/errors"
	"ums/model"
	"ums/service"
	"ums/types"
)

func PageUser(ctx *gin.Context) {
	// 检验参数
	in := types.PageUserRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}

	if list, total, err := service.PageUser(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespList(in.Page, len(list), int(total), list)
	}
}

func UserTeamIds(ctx *gin.Context) {
	ids := model.CurUserTeamIds(ctx)
	ctx.RespData(ids)
}

func GetUser(ctx *gin.Context) {
	in := types.GetUserRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if data, err := service.GetUser(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(data)
	}
}

func CurUser(ctx *gin.Context) {
	if user, err := service.CurUser(ctx); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(user)
	}
}

func AddUser(ctx *gin.Context) {
	// 检验参数
	in := types.AddUserRequest{}
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.AddUser(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateUser(ctx *gin.Context) {
	// 检验参数
	in := types.UpdateUserRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.UpdateUser(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteUser(ctx *gin.Context) {
	// 检验参数
	in := types.DeleteUserRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.DeleteUser(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UserLogin(ctx *gin.Context) {
	// 检验参数
	in := types.UserLoginRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	// 调用实现
	if resp, err := service.UserLogin(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

package handler

import (
	"github.com/limeschool/gin"
	"ums/errors"
	"ums/service"
	"ums/types"
)

func AllTeam(ctx *gin.Context) {
	if resp, err := service.AllTeam(ctx); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func AddTeam(ctx *gin.Context) {
	in := types.AddTeamRequest{}
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddTeam(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateTeam(ctx *gin.Context) {
	in := types.UpdateTeamRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.UpdateTeam(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteTeam(ctx *gin.Context) {
	in := types.DeleteTeamRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.DeleteTeam(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

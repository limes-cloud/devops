package handler

import (
	"configure/errors"
	"configure/service"
	"configure/types"
	"github.com/limeschool/gin"
)

func AddServiceSystemField(ctx *gin.Context) {
	in := types.AddServiceSystemFieldRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddServiceSystemField(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func AllServiceSystemField(ctx *gin.Context) {
	in := types.AllServiceSystemFieldRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if resp, err := service.AllServiceSystemField(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func PageSystemField(ctx *gin.Context) {
	in := types.PageSystemFieldRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if resp, total, err := service.PageSystemField(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespList(in.Page, in.Count, int(total), resp)
	}
}

func AddSystemField(ctx *gin.Context) {
	in := types.AddSystemFieldRequest{}
	if err := ctx.ShouldBind(&in); err != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddSystemField(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateSystemField(ctx *gin.Context) {
	in := types.UpdateSystemFieldRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.UpdateSystemField(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteSystemField(ctx *gin.Context) {
	in := types.DeleteSystemFieldRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.DeleteSystemField(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

package handler

import (
	"configure/errors"
	"configure/service"
	"configure/types"
	"github.com/limeschool/gin"
)

func PageServiceField(ctx *gin.Context) {
	in := types.PageServiceFieldRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if resp, total, err := service.PageServiceField(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespList(in.Page, in.Count, int(total), resp)
	}
}

func AddServiceField(ctx *gin.Context) {
	in := types.AddServiceFieldRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddServiceField(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateServiceField(ctx *gin.Context) {
	in := types.UpdateServiceFieldRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.UpdateServiceField(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteServiceField(ctx *gin.Context) {
	in := types.DeleteServiceFieldRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.DeleteServiceField(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

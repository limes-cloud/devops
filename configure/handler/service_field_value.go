package handler

import (
	"configure/errors"
	"configure/service"
	"configure/types"
	"github.com/limeschool/gin"
)

func AllServiceFieldValue(ctx *gin.Context) {
	in := types.AllServiceFieldValueRequest{}
	if err := ctx.ShouldBind(&in); err != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if resp, err := service.AllServiceFieldValue(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func AddServiceFieldValue(ctx *gin.Context) {
	in := types.AddServiceFieldValueRequest{}
	if err := ctx.ShouldBind(&in); err != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddServiceFieldValue(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

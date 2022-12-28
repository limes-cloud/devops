package handler

import (
	"dc/errors"
	"dc/service"
	"dc/types"
	"github.com/limeschool/gin"
)

func AddNetwork(ctx *gin.Context) {
	in := types.AddNetworkRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddNetwork(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteNetwork(ctx *gin.Context) {
	in := types.DeleteNetworkRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.DeleteNetwork(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

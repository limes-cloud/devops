package handler

import (
	"github.com/limeschool/gin"
	"service/errors"
	"service/service"
	"service/types"
)

func AllImageRegistries(ctx *gin.Context) {
	if resp, err := service.AllImageRegistries(ctx); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func AllImageRegistryFilter(ctx *gin.Context) {
	if resp, err := service.AllImageRegistryFilter(ctx); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func AddImageRegistry(ctx *gin.Context) {
	in := types.AddImageRegistryRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddImageRegistry(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateImageRegistry(ctx *gin.Context) {
	in := types.UpdateImageRegistryRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.UpdateImageRegistry(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteImageRegistry(ctx *gin.Context) {
	in := types.DeleteImageRegistryRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.DeleteImageRegistry(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func ConnectImageRegistry(ctx *gin.Context) {
	in := types.ConnectImageRegistryRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}

	if err := service.ConnectImageRegistry(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

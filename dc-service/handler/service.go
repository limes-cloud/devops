package handler

import (
	"dc/errors"
	"dc/service"
	"dc/types"
	"github.com/limeschool/gin"
)

func GetServiceRelease(ctx *gin.Context) {
	in := types.GetServiceReleaseRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if data, err := service.GetServiceRelease(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(data)
	}
}

func GetServicePods(ctx *gin.Context) {
	in := types.GetServicePodsRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if data, err := service.GetServicePods(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(data)
	}
}

func AddService(ctx *gin.Context) {
	in := types.AddServiceRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddService(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteService(ctx *gin.Context) {
	in := types.DeleteServiceRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.DeleteService(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

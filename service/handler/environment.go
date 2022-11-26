package handler

import (
	"github.com/limeschool/gin"
	"service/errors"
	"service/service"
	"service/types"
)

func AllEnvironment(ctx *gin.Context) {
	in := types.AllEnvironmentRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if resp, err := service.AllEnvironment(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func AllEnvironmentFilter(ctx *gin.Context) {
	if resp, err := service.AllEnvironmentFilter(ctx); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func AddEnvironment(ctx *gin.Context) {
	in := types.AddEnvironmentRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddEnvironment(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateEnvironment(ctx *gin.Context) {
	in := types.UpdateEnvironmentRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.UpdateEnvironment(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteEnvironment(ctx *gin.Context) {
	in := types.DeleteEnvironmentRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.DeleteEnvironment(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateServiceEnv(ctx *gin.Context) {
	in := types.UpdateServiceEnvRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.UpdateServiceEnv(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

//
//func AllServiceEnv(ctx *gin.Context) {
//	in := types.AllServiceEnvRequest{}
//	if ctx.ShouldBind(&in) != nil {
//		ctx.RespError(errors.ParamsError)
//		return
//	}
//	if in.EnvId == 0 && in.SrvId == 0 {
//		ctx.RespError(errors.ParamsError)
//		return
//	}
//
//	if list, err := service.AllServiceEnv(ctx, &in); err != nil {
//		ctx.RespError(err)
//	} else {
//		ctx.RespData(list)
//	}
//}

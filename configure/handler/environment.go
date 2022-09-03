package handler

import (
	"configure/errors"
	"configure/service"
	"configure/types"
	"github.com/limeschool/gin"
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

func EnvironmentConnect(ctx *gin.Context) {
	in := types.EnvironmentConnectRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if _, err := service.EnvironmentConnect(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
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

func UpdateEnvService(ctx *gin.Context) {
	in := types.UpdateEnvServiceRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.UpdateEnvService(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func AllEnvService(ctx *gin.Context) {
	in := types.AllEnvServiceRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if in.EnvId == 0 && in.SrvId == 0 {
		ctx.RespError(errors.ParamsError)
		return
	}

	if list, err := service.AllEnvService(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(list)
	}
}

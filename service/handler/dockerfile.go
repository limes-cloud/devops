package handler

import (
	"github.com/limeschool/gin"
	"service/errors"
	"service/service"
	"service/types"
)

func PageDockerfile(ctx *gin.Context) {
	in := types.AllDockerfileRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if resp, err := service.AllDockerfile(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func AllDockerfileFilter(ctx *gin.Context) {
	if resp, err := service.AllDockerfileFilter(ctx); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func AddDockerfile(ctx *gin.Context) {
	in := types.AddDockerfileRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddDockerfile(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateDockerfile(ctx *gin.Context) {
	in := types.UpdateDockerfileRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.UpdateDockerfile(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteDockerfile(ctx *gin.Context) {
	in := types.DeleteDockerfileRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.DeleteDockerfile(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

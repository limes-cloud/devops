package handler

import (
	"github.com/limeschool/gin"
	"service/consts"
	"service/errors"
	"service/service"
	"service/types"
)

func AllReleaseTypes(ctx *gin.Context) {
	ctx.RespData([]string{consts.Dc, consts.K8s})
}

func AllReleaseStatus(ctx *gin.Context) {
	ctx.RespData(consts.ReleaseStatus)
}

func PageRelease(ctx *gin.Context) {
	in := types.AllReleaseRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if resp, err := service.AllRelease(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func AddRelease(ctx *gin.Context) {
	in := types.AddReleaseRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddRelease(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateRelease(ctx *gin.Context) {
	in := types.UpdateReleaseRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.UpdateRelease(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteRelease(ctx *gin.Context) {
	in := types.DeleteReleaseRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.DeleteRelease(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func AllReleaseImages(ctx *gin.Context) {
	in := types.AllReleaseImagesRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if resp, err := service.AllReleaseImages(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

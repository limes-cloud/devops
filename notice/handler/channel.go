package handler

import (
	"github.com/limeschool/gin"
	"notice/errors"
	"notice/service"
	"notice/types"
)

func AllChannel(ctx *gin.Context) {
	in := types.GetChannelRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if list, err := service.AllChannel(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(list)
	}
}

func AllChannelFilter(ctx *gin.Context) {
	in := types.GetChannelRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if list, err := service.AllChannelFilter(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(list)
	}
}

func AddChannel(ctx *gin.Context) {
	in := types.AddChannelRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddChannel(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateChannel(ctx *gin.Context) {
	in := types.UpdateChannelRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.UpdateChannel(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteChannel(ctx *gin.Context) {
	in := types.DeleteChannelRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.DeleteChannel(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

package handler

import (
	"github.com/limeschool/gin"
	"service/errors"
	"service/service"
	"service/types"
)

func PagePackLog(ctx *gin.Context) {
	in := types.PagePackRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if list, total, err := service.PagePackLog(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespList(in.Page, in.Count, int(total), list)
	}
}

func AddPack(ctx *gin.Context) {
	in := types.AddPackRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddPack(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

package handler

import (
	"github.com/limeschool/gin"
	"service/errors"
	"service/service"
	"service/types"
)

func PageReleaseLog(ctx *gin.Context) {
	in := types.PageReleaseLogRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if list, total, err := service.PageReleaseLog(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespList(in.Page, in.Count, int(total), list)
	}
}

func AddReleaseLog(ctx *gin.Context) {
	in := types.AddReleaseLogRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddReleaseLog(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

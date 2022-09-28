package handler

import (
	"github.com/limeschool/gin"
	"notice/errors"
	"notice/service"
	"notice/types"
)

func PageLog(ctx *gin.Context) {
	in := types.GetLogRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if list, total, err := service.PageLog(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespList(in.Page, in.Count, int(total), list)
	}
}

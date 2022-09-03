package handler

import (
	"configure/errors"
	"configure/service"
	"configure/types"
	"github.com/limeschool/gin"
)

func PageOperatorLog(ctx *gin.Context) {
	in := types.PageServiceFieldRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if resp, total, err := service.PageServiceField(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespList(in.Page, in.Count, int(total), resp)
	}
}

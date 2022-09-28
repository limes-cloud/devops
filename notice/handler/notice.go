package handler

import (
	"github.com/limeschool/gin"
	"notice/errors"
	"notice/service"
	"notice/types"
)

func PageNotice(ctx *gin.Context) {
	in := types.GetNoticeRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if list, total, err := service.PageNotice(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespList(in.Page, in.Count, int(total), list)
	}
}

func AddNotice(ctx *gin.Context) {
	in := types.AddNoticeRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddNotice(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateNotice(ctx *gin.Context) {
	in := types.UpdateNoticeRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.UpdateNotice(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteNotice(ctx *gin.Context) {
	in := types.DeleteNoticeRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.DeleteNotice(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

package handler

import (
	"configure/errors"
	"configure/service"
	"configure/types"
	"github.com/limeschool/gin"
)

func ParseTemplate(ctx *gin.Context) {
	in := types.ParseTemplateRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if resp, err := service.ParseTemplate(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func AllTemplate(ctx *gin.Context) {
	in := types.AllTemplateRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if resp, err := service.AllTemplate(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func GetTemplate(ctx *gin.Context) {
	in := types.GetTemplateRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if in.ID == 0 && in.Keyword == "" {
		ctx.RespError(errors.ParamsError)
		return
	}
	resp, _ := service.GetTemplate(ctx, &in)
	ctx.RespData(resp)
}

func AddTemplate(ctx *gin.Context) {
	in := types.AddTemplateRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddTemplate(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateTemplate(ctx *gin.Context) {
	in := types.UpdateTemplateRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.UpdateTemplate(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

package handler

import (
	"github.com/limeschool/gin"
	"go.uber.org/zap"
	"ums/errors"
	"ums/service"
	"ums/types"
)

func RefreshToken(ctx *gin.Context) {
	in := types.RefreshTokenRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	// 调用实现
	if resp, err := service.RefreshToken(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func Auth(ctx *gin.Context) {
	if status, err := service.Auth(ctx); err != nil {
		ctx.Log.Error("权限验证失败", zap.Error(err))
		ctx.AbortWithStatus(status)
	} else {
		ctx.RespSuccess()
	}
}

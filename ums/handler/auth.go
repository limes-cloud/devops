package handler

import (
	"fmt"
	"github.com/limeschool/gin"
	"ums/errors"
	"ums/meta"
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
		if custom, is := err.(*gin.CustomError); is && custom.Code == errors.TokenExpiredError.Code {
			err = errors.RefTokenExpiredError
		}
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func Auth(ctx *gin.Context) {
	if err := service.Auth(ctx); err != nil {
		errCode := 4000
		msg := err.Error()
		if customErr, ok := err.(*gin.CustomError); ok {
			errCode = customErr.Code
		}
		ctx.Writer.Header().Set(meta.ErrorHeader, msg)
		ctx.Writer.Header().Set(meta.ErrorCodeHeader, fmt.Sprint(errCode))
		ctx.AbortWithStatus(401)
	} else {
		ctx.RespSuccess()
	}
}

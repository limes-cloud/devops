package service

import (
	"github.com/limeschool/gin"
	"ums/model"
	"ums/tools/address"
	"ums/tools/ua"
	"ums/types"
)

func AddLoginLog(ctx *gin.Context, username string, err error) error {
	ip := ctx.Request.Header.Get("x-real-ip")
	userAgent := ctx.Request.Header.Get("User-Agent")
	info := ua.Parse(userAgent)
	desc := ""
	code := 0

	if err != nil {
		if customErr, is := err.(*gin.CustomError); is {
			code = customErr.Code
			desc = customErr.Msg
		} else {
			desc = err.Error()
		}
	}

	log := model.LoginLog{
		Username:    username,
		IP:          ip,
		Address:     address.GetAddress(ip),
		Browser:     info.Name,
		Status:      err == nil,
		Description: desc,
		Code:        code,
		Device:      info.OS + " " + info.OSVersion,
	}
	return log.Create(ctx)
}

func PageLoginLog(ctx *gin.Context, in *types.LoginLogRequest) ([]model.LoginLog, int64, error) {
	log := model.LoginLog{}
	return log.Page(ctx, in.Page, in.Count, in)
}

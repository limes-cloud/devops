package meta

import (
	"encoding/json"
	"github.com/limeschool/gin"
)

const (
	ApiPrefix       = "/api/v1"
	Token           = "Authorization"
	UserHeader      = "x-user"
	ErrorHeader     = "x-error"
	ErrorCodeHeader = "x-error-code"
	SuperAdmin      = "super_admin"
)

type Userinfo struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

func parseUserinfo(ctx *gin.Context) Userinfo {
	var userinfo Userinfo
	_ = json.Unmarshal([]byte(ctx.Request.Header.Get(UserHeader)), &userinfo)
	return userinfo
}

func UserID(ctx *gin.Context) int64 {
	user := parseUserinfo(ctx)
	return user.UserId
}

func UserName(ctx *gin.Context) string {
	user := parseUserinfo(ctx)
	return user.UserName
}

func User(ctx *gin.Context) Userinfo {
	return parseUserinfo(ctx)
}

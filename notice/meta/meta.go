package meta

import (
	"encoding/json"
	"github.com/limeschool/gin"
)

const (
	UserInfo = "x-user"
)

type Userinfo struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

func parseUserinfo(ctx *gin.Context) Userinfo {
	var userinfo Userinfo
	_ = json.Unmarshal([]byte(ctx.Request.Header.Get(UserInfo)), &userinfo)
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

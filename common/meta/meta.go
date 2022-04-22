package meta

import (
	"context"
	"devops/common/response"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

const (
	ApiPrefix      = "/api/v1"
	UserIDKey      = "user_id"
	UserNameKey    = "user_name"
	RoleNameKey    = "role_name"
	RoleIdKey      = "role_id"
	RoleKeywordKey = "role_keyword"
	SuperAdmin     = "super_admin"
)

func UserId(ctx context.Context) int64 {
	num, _ := ctx.Value(UserIDKey).(int64)
	return num
}

func UserName(ctx context.Context) string {
	name, _ := ctx.Value(UserNameKey).(string)
	return name
}

func UserRoleName(ctx context.Context) string {
	name, _ := ctx.Value(RoleNameKey).(string)
	return name
}

func UserRoleId(ctx context.Context) int64 {
	i, _ := ctx.Value(RoleIdKey).(int64)
	return i
}

func RoleKeyword(ctx context.Context) string {
	i, _ := ctx.Value(RoleKeywordKey).(string)
	return i
}

func IsSuper(ctx context.Context) bool {
	i, _ := ctx.Value(RoleKeywordKey).(string)
	return i == "super_admin" || i == "superadmin"
}

func ParsePwd(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CompareHashPwd(p1, p2 string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p1), []byte(p2)) == nil
}

type Userinfo struct {
	UserID      int64  `json:"user_id"`
	UserName    string `json:"user_name"`
	RoleName    string `json:"role_name"`
	RoleId      int64  `json:"role_id"`
	RoleKeyword string `json:"role_keyword"`
}

func SetUserIdHandle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if r.Method == http.MethodOptions {
			next(w, r.WithContext(ctx))
			return
		}
		if strings.Contains(r.RequestURI, "/auth/validate") || strings.Contains(r.RequestURI, "/user/login") {
			next(w, r.WithContext(ctx))
			return
		}
		userStr := r.Header.Get("X-User")
		if userStr == "" { //从redis 中获取白名单
			next(w, r.WithContext(ctx))
			return
		}
		userinfo, err := ParseUserInfo(userStr)
		if err != nil {
			httpx.OkJson(w, response.Resp{Code: 402, Msg: err.Error()})
			return
		}

		ctx = context.WithValue(ctx, UserIDKey, userinfo.UserID)
		ctx = context.WithValue(ctx, UserNameKey, userinfo.UserName)
		ctx = context.WithValue(ctx, RoleNameKey, userinfo.RoleName)
		ctx = context.WithValue(ctx, RoleIdKey, userinfo.RoleId)
		ctx = context.WithValue(ctx, RoleKeywordKey, userinfo.RoleKeyword)
		next(w, r.WithContext(ctx))
	}
}

func ParseUserInfo(data string) (*Userinfo, error) {
	userinfo := Userinfo{}
	if json.Unmarshal([]byte(data), &userinfo) != nil {
		return nil, errors.New("用户信息header解析失败")
	}
	return &userinfo, nil
}

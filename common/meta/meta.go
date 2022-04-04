package meta

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"strings"
)

const (
	UserIDKey   = "user_id"
	UserNameKey = "user_name"
)

func RedisUserKey(uid int64) string {
	return fmt.Sprintf("%v:%v", UserIDKey, uid)
}

func UserId(ctx context.Context) int64 {
	num, _ := ctx.Value(UserIDKey).(int64)
	return num
}

func UserName(ctx context.Context) string {
	name, _ := ctx.Value(UserNameKey).(string)
	return name
}

func UserRole(ctx context.Context) string {
	name, _ := ctx.Value(UserNameKey).(string)
	return name
}

func ParsePwd(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CompareHashPwd(p1, p2 string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p1), []byte(p2)) == nil
}

func SetUserIdHandle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if strings.Contains(r.RequestURI, "/auth/validate") || strings.Contains(r.RequestURI, "/user/login") {
			next(w, r.WithContext(ctx))
			return
		}
		userId, _ := strconv.ParseInt(r.Header.Get("X-User"), 10, 64)
		ctx = context.WithValue(ctx, UserIDKey, userId)
		next(w, r.WithContext(ctx))
	}
}

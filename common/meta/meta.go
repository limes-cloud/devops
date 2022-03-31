package meta

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	UserKey = "user_id"
)

func UserId(ctx context.Context) int64 {
	num, _ := ctx.Value(UserKey).(json.Number)
	uid, _ := num.Int64()
	return uid
}

func SetUserIdHandle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := strconv.ParseInt(r.Header.Get("X-User"), 10, 64)
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, userId)
		next(w, r.WithContext(ctx))
	}
}

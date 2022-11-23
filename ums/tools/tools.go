package tools

import (
	"github.com/limeschool/gin"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
	"ums/consts"
)

func DataDup(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "Duplicate")
}

func ParsePwd(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CompareHashPwd(p1, p2 string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p1), []byte(p2)) == nil
}

type ListType interface {
	~string | ~int | ~int64 | ~[]byte | ~rune | ~float64
}

func InList[ListType comparable](list []ListType, val ListType) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

func DelRedis(ctx *gin.Context, key string) {
	go func() {
		time.Sleep(1 * time.Second)
		ctx.Redis(consts.REDIS).Del(ctx, key)
	}()
}

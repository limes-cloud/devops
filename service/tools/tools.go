package tools

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"github.com/limeschool/gin"
	"service/consts"
	"strings"
	"time"
)

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

func GenToken() string {
	uid := uuid.New().String()
	return fmt.Sprintf("%x", md5.Sum([]byte(uid)))
}

func GenVersion() string {
	uid := uuid.New().String()
	data := fmt.Sprintf("%x", md5.Sum([]byte(uid)))[8:24]
	return strings.ToUpper(data)
}

func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

func Diff(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

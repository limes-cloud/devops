package service

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/limeschool/gin"
	"strings"
	"time"
	"ums/consts"
	"ums/errors"
	"ums/meta"
	"ums/model"
	"ums/types"
)

func Auth(ctx *gin.Context) (int, error) {
	path := ctx.Request.Header.Get("X-Original-Uri")
	method := ctx.Request.Header.Get("X-Original-Method")
	if strings.Contains(path, "?") {
		path = strings.Split(path, "?")[0]
	}
	if IsWhitePath(path, method) {
		return 200, nil
	}
	// token 解析
	token := ctx.Request.Header.Get(meta.Token)
	userId, err := ParseToken(ctx, "user_id", token)
	if err != nil {
		return 401, err
	}
	// 获取用户信息
	user := model.User{}
	if err = user.OneByID(ctx, userId); err != nil {
		return 401, err
	}
	// 判断当前token 和 redis存的是否一致
	if cacheToken, _ := ctx.Redis(consts.REDIS).Get(ctx, storeKey(userId)).Result(); cacheToken != "" {
		if token != cacheToken {
			return 401, errors.DulDeviceLoginError
		}
	}

	// 判断rbac权限
	if user.Role.Keyword != meta.SuperAdmin && !IsBaseApiPath(ctx, path, method) {
		if is, _ := ctx.Rbac().Enforce(user.Role.Keyword, path, method); !is {
			return 403, errors.NotResourcePower
		}
	}

	// 设置返回数据
	userinfo := gin.H{
		"user_id":   user.ID,
		"user_name": user.Name,
	}
	infoByte, _ := json.Marshal(userinfo)
	ctx.Writer.Header().Set(meta.UserInfo, string(infoByte))
	return 200, nil
}

func IsWhitePath(path, method string) bool {
	if strings.Contains(path, "healthy") {
		return true
	}
	key := strings.ToLower(method + ":" + path)
	return consts.WhitelistApi[key] == true
}

func IsBaseApiPath(ctx *gin.Context, path, method string) bool {
	redisKey := consts.RedisBaseApi
	var listMenu []*model.Menu
	if str, _ := ctx.Redis(consts.REDIS).Get(ctx, redisKey).Result(); str != "" {
		_ = json.Unmarshal([]byte(str), &listMenu)
	} else {
		menu := model.Menu{}
		listMenu, _ = menu.All(ctx, "permission = ? and type = 'A'", consts.BaseApi)
		byteData, _ := json.Marshal(listMenu)
		ctx.Redis(consts.REDIS).Set(ctx, redisKey, string(byteData), 1*time.Hour)
	}

	for _, item := range listMenu {
		if item.Path == path && item.Method == method {
			return true
		}
	}
	return false
}

func RefreshToken(ctx *gin.Context, in *types.RefreshTokenRequest) (*types.UserLoginResponse, error) {
	userId, err := ParseToken(ctx, "refresh_id", in.Token)
	if err != nil {
		return nil, err
	}
	return GenToken(ctx, userId)
}

func storeKey(id int64) string {
	return fmt.Sprintf("ums_user_token_%v", id)
}

func GenToken(ctx *gin.Context, userId int64) (*types.UserLoginResponse, error) {
	duration := int64(ctx.Config.GetDefaultInt("jwt.exp", 1800))
	// 生成token
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Unix() + duration
	claims["iat"] = time.Now().Unix()
	claims["user_id"] = userId
	tokenJwt := jwt.New(jwt.SigningMethodHS256)
	tokenJwt.Claims = claims
	token, err := tokenJwt.SignedString([]byte(ctx.Config.GetString("jwt.secret")))
	if err != nil {
		return nil, err
	}
	// 生成refreshToken
	var maxTime int64 = 3600 * 24 * 2
	refreshClaims := make(jwt.MapClaims)
	refreshClaims["exp"] = time.Now().Unix() + maxTime
	refreshClaims["iat"] = time.Now().Unix()
	refreshClaims["refresh_id"] = userId
	refreshJwt := jwt.New(jwt.SigningMethodHS256)
	refreshJwt.Claims = refreshClaims
	refreshToken, err := refreshJwt.SignedString([]byte(ctx.Config.GetString("jwt.secret")))
	if err != nil {
		return nil, err
	}
	_, err = ctx.Redis(consts.REDIS).Set(ctx, storeKey(userId), token, time.Duration(duration+maxTime)*time.Second).Result()
	return &types.UserLoginResponse{
		Token:        token,
		Duration:     int(duration),
		RefreshToken: refreshToken,
	}, err
}

func ParseToken(ctx *gin.Context, key, token string) (int64, error) {
	var m jwt.MapClaims
	secret := ctx.Config.GetString("jwt.secret")
	parser, err := jwt.ParseWithClaims(token, &m, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !parser.Valid {
		return -1, errors.TokenValidateError
	}
	id, _ := m[key].(float64)
	return int64(id), err
}

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

const (
	userKey    = "user_id"
	refreshKey = "refresh_id"
)

// Auth 鉴权接口
func Auth(ctx *gin.Context) error {

	// 获取方法以及path
	path := ctx.Request.Header.Get("X-Original-Uri")
	method := ctx.Request.Header.Get("X-Original-Method")
	if strings.Contains(path, "?") {
		path = strings.Split(path, "?")[0]
	}

	// 判断是否在白名单内
	if IsWhitePath(path, method) {
		return nil
	}

	// 获取token
	token := ctx.Request.Header.Get(meta.Token)
	if token == "" {
		return errors.TokenEmptyError
	}

	// 进行token解析
	userId, err := ParseToken(userKey, token)
	if err != nil {
		return err
	}

	// 判断当前token 和 redis存的是否一致
	if tokenInfo := loadToken(ctx, userId); tokenInfo == nil {
		return errors.RefTokenExpiredError
	} else if tokenInfo.Token != token {
		return errors.DulDeviceLoginError
	}

	// 获取用户信息
	user := model.User{}
	if err = user.OneByID(ctx, userId); err != nil {
		return errors.UserNameNotFoundError
	}

	// 判断rbac权限
	if user.Role.Keyword != meta.SuperAdmin && !IsBaseApiPath(ctx, path, method) {
		if is, _ := ctx.Rbac().Enforce(user.Role.Keyword, path, method); !is {
			return errors.NotResourcePower
		}
	}

	// 设置返回数据
	userinfo := gin.H{
		"user_id":   user.ID,
		"user_name": user.Name,
	}
	infoByte, _ := json.Marshal(userinfo)
	ctx.Writer.Header().Set(meta.UserHeader, string(infoByte))
	return nil
}

// IsWhitePath 判断是否为白名单
func IsWhitePath(path, method string) bool {
	if strings.Contains(path, "healthy") {
		return true
	}
	key := strings.ToLower(method + ":" + path)
	return consts.WhitelistApi[key] == true
}

// IsBaseApiPath 是否为基础白名单
func IsBaseApiPath(ctx *gin.Context, path, method string) bool {
	// 获取基础api列表
	menu := model.Menu{}
	list := menu.GetBaseApiPath(ctx)

	// 判断是否为基础api
	for _, item := range list {
		if item.Path == path && item.Method == method {
			return true
		}
	}
	return false
}

// RefreshToken 进行token刷新
func RefreshToken(ctx *gin.Context, in *types.RefreshTokenRequest) (*types.UserLoginResponse, error) {

	userId, err := ParseToken("refresh_id", in.Token)
	if err != nil {
		return nil, err
	}

	// 判断当前token是否还生效
	if tokenInfo := loadToken(ctx, userId); tokenInfo == nil || tokenInfo.RefreshToken != in.Token {
		return nil, errors.RefTokenExpiredError
	}

	return GenToken(ctx, userId)
}

// storeKey 获取缓存的key
func storeKey(id int64) string {
	return fmt.Sprintf("ums_user_token_%v", id)
}

// GenToken 生成token
func GenToken(ctx *gin.Context, userId int64) (*types.UserLoginResponse, error) {

	// 生成token携带信息
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Unix() + consts.Jwt.Expire
	claims["iat"] = time.Now().Unix()
	claims[userKey] = userId
	tokenJwt := jwt.New(jwt.SigningMethodHS256)
	tokenJwt.Claims = claims
	token, err := tokenJwt.SignedString([]byte(consts.Jwt.Secret))
	if err != nil {
		return nil, err
	}

	// 生成refreshToken
	refreshClaims := make(jwt.MapClaims)
	refreshClaims["exp"] = time.Now().Unix() + consts.Jwt.MaxExpire
	refreshClaims["iat"] = time.Now().Unix()
	refreshClaims[refreshKey] = userId
	refreshJwt := jwt.New(jwt.SigningMethodHS256)
	refreshJwt.Claims = refreshClaims
	refreshToken, err := refreshJwt.SignedString([]byte(consts.Jwt.Secret))
	if err != nil {
		return nil, err
	}

	// 进行token缓存
	resp := &types.UserLoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}
	return resp, storeToken(ctx, userId, resp)
}

// storeToken 存储token到缓存
func storeToken(ctx *gin.Context, uid int64, in *types.UserLoginResponse) error {
	cache := ctx.Redis(consts.REDIS)
	byteData, _ := json.Marshal(in)
	tokenExpire := time.Duration(consts.Jwt.MaxExpire) * time.Second
	return cache.Set(ctx, storeKey(uid), string(byteData), tokenExpire).Err()
}

// loadToken 从缓存中加载token
func loadToken(ctx *gin.Context, uid int64) *types.UserLoginResponse {
	str, err := ctx.Redis(consts.REDIS).Get(ctx, storeKey(uid)).Result()
	if err != nil {
		return nil
	}

	resp := &types.UserLoginResponse{}
	if json.Unmarshal([]byte(str), resp) != nil {
		return nil
	}
	return resp
}

// ParseToken 解析token
func ParseToken(key, token string) (int64, error) {
	var m jwt.MapClaims
	parser, err := jwt.ParseWithClaims(token, &m, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.Jwt.Secret), nil
	})

	if err != nil || !parser.Valid {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return -1, errors.TokenExpiredError
		}
		return -1, errors.TokenValidateError
	}

	id, _ := m[key].(float64)
	return int64(id), nil
}

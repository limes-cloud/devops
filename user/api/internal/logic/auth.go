package logic

import (
	"context"
	"devops/common/errorx"
	"devops/common/meta"
	"devops/common/tools"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/rest/token"
	"net/http"
	"strings"

	"devops/user/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthLogic) Auth(r *http.Request, w http.ResponseWriter) error {
	// 拿到真是的path 信息
	path := r.Header.Get("X-Original-Uri")
	method := r.Header.Get("X-Original-Method")
	if strings.Contains(path, "?") {
		path = strings.Split(path, "?")[0]
	}
	path = strings.Replace(path, meta.ApiPrefix, "", 1)
	// 判断是否是白名单
	if l.IsWhitePath(path) {
		logx.WithContext(l.ctx).Infof("url:%v是白名单", path)
		w.Header().Set("x-user", meta.WhitelistKey)
		return nil
	}

	// 判断token是否为空
	if r.Header.Get("Authorization") == "" {
		return errors.New(errorx.AuthErr)
	}

	// token 解析判断
	userinfo, err := l.ParseToken(r)
	if err != nil || userinfo == "" {
		logx.WithContext(l.ctx).Error(err)
		return errors.New(errorx.AuthErr)
	}

	// RBAC 权限控制
	if l.rbac(userinfo, path, method) != nil {
		return errors.New(errorx.RbacErr)
	}

	w.Header().Set("x-user", userinfo)
	return err
}

func (l *AuthLogic) rbac(info, path, method string) error {
	user, err := meta.ParseUserInfo(info)
	if err != nil {
		return err
	}
	if user.RoleKeyword == meta.SuperAdmin {
		return nil
	}
	if is, _ := l.svcCtx.Rbac.Enforce(user.RoleKeyword, path, method); !is {
		return errors.New(errorx.RbacErr)
	}
	return nil
}

func (l *AuthLogic) IsWhitePath(path string) bool {
	if strings.Contains(path, "check_health") {
		return true
	}
	return tools.InListStr(l.svcCtx.Config.GetStringSlice("whitelist"), path)
}

func (l *AuthLogic) ParseToken(r *http.Request) (string, error) {
	parser := token.NewTokenParser()
	tok, err := parser.ParseToken(r, l.svcCtx.Config.GetString("auth.access_secret"), "")
	if err != nil {
		return "", err
	}
	if tok.Valid {
		claims, _ := tok.Claims.(jwt.MapClaims) // 解析token中对内容
		if info, ok := claims["userinfo"].(string); ok {
			return info, nil
		}
	}
	return "", err
}

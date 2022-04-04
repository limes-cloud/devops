package logic

import (
	"context"
	"devops/common/meta"
	"devops/common/tools"
	"encoding/json"
	"errors"
	"fmt"
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
	if strings.Contains(path, "?") {
		path = strings.Split(path, "?")[0]
	}

	// 判断是否是白名单
	if l.IsWhitePath(path) {
		return nil
	}

	// 判断token是否为空
	if r.Header.Get("Authorization") == "" {
		return errors.New("token 不能为空")
	}

	// token 解析判断
	userId, err := l.ParseToken(r)
	if err != nil || userId == 0 {
		return errors.New("token 鉴权失败")
	}

	// RBAC 权限控制
	// todo this write rbac logic

	w.Header().Set("x-user", fmt.Sprintf("%d", userId))
	return err
}

func (l *AuthLogic) IsWhitePath(path string) bool {
	return tools.InListStr(l.svcCtx.Config.WhitePath, path)
}

func (l *AuthLogic) ParseToken(r *http.Request) (int64, error) {
	parser := token.NewTokenParser()
	tok, err := parser.ParseToken(r, l.svcCtx.Config.Auth.AccessSecret, "")
	if err != nil {
		return 0, err
	}
	if tok.Valid {
		claims, ok := tok.Claims.(jwt.MapClaims) // 解析token中对内容
		if ok {
			userId, _ := claims[meta.UserIDKey].(json.Number).Int64() // 获取userId 并且到后端redis校验是否过期
			return userId, nil
		}
	}
	return 0, err
}

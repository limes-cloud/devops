package logic

import (
	"context"
	"devops/common/meta"
	"devops/common/tools/rsa"
	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"
	"devops/user/models"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	user := models.User{}
	resp = new(types.LoginResponse)

	password, err := rsa.Decode(l.svcCtx.Config.Rsa.PrivateKey, req.Password)
	if err != nil {
		return nil, err
	}

	if user.OneByCall(func(db *gorm.DB) *gorm.DB {
		return db.Where("phone = ? or email = ?", req.UserName, req.UserName)
	}) != nil {
		return nil, errors.New("账号不存在")
	}

	if !*user.Status && user.ID != 1 {
		return nil, errors.New("当前用户已被禁用")
	}
	if !meta.CompareHashPwd(user.Password, password) {
		return nil, errors.New("密码错误")
	}

	resp.Token, err = l.NewToken(user)
	return resp, err
}

// NewToken 进行token生成
func (l *LoginLogic) NewToken(user models.User) (string, error) {
	auth := l.svcCtx.Config.Auth
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Unix() + auth.AccessExpire
	claims["iat"] = time.Now().Unix()
	b, _ := json.Marshal(map[string]interface{}{
		meta.UserIDKey:      user.ID,
		meta.UserNameKey:    user.Name,
		meta.RoleNameKey:    user.RoleName,
		meta.RoleIdKey:      user.RoleID,
		meta.RoleKeywordKey: user.Keyword,
	})

	claims["userinfo"] = string(b)

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(auth.AccessSecret))
}

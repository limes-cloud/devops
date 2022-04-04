package logic

import (
	"context"
	"devops/common/meta"
	"devops/common/tools/rsa"
	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"
	"devops/user/models"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
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

	db := l.svcCtx.Orm
	if db.Where("phone = ? or email = ?", req.UserName, req.UserName).First(&user).Error != nil {
		return nil, errors.New("账号不存在")
	}
	if !meta.CompareHashPwd(user.Password, password) {
		return nil, errors.New("密码错误")
	}

	resp.Token, err = l.NewToken(user.ID, user.Name)
	return resp, err
}

// NewToken 进行token生成
func (l *LoginLogic) NewToken(id int64, name string) (string, error) {
	auth := l.svcCtx.Config.Auth
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Unix() + auth.AccessExpire
	claims["iat"] = time.Now().Unix()
	claims[meta.UserIDKey] = id
	claims[meta.UserNameKey] = name
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(auth.AccessSecret))
}

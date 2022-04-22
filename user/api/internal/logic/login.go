package logic

import (
	"context"
	"devops/common/meta"
	"devops/common/model"
	"devops/common/tools/address"
	"devops/common/tools/rsa"
	"devops/common/tools/ua"
	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"
	"devops/user/models"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"net/http"
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

func (l *LoginLogic) Login(r *http.Request, req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	if err = l.ipLimit(r); err != nil {
		return
	}
	user := models.User{}
	resp = new(types.LoginResponse)
	password, err := rsa.Decode(l.svcCtx.Config.GetString("rsa.public_file"), req.Password)
	if err != nil {
		return nil, err
	}

	defer func() {
		l.NewLoginLog(r, req.UserName, err)
	}()
	if user.One(nil, func(db *gorm.DB) *gorm.DB {
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

func (l *LoginLogic) NewLoginLog(r *http.Request, username string, err error) {
	ip := r.Header.Get("x-real-ip")
	userAgent := r.Header.Get("User-Agent")
	info := ua.Parse(userAgent)
	errStr := ""
	if err != nil {
		errStr = err.Error()
	}
	log := models.LoginLog{
		Username:    username,
		IP:          ip,
		Address:     address.GetAddress(ip), //这里应该对接获取地址
		Browser:     info.Name,
		Device:      info.OS + " " + info.OSVersion,
		Status:      err == nil,
		Description: errStr,
	}
	log.Create()
}

// IpLimit 同一个ip，一天只能登陆错误10次
func (l *LoginLogic) ipLimit(r *http.Request) error {
	ip := r.Header.Get("x-real-ip")
	log := models.LoginLog{}
	count, err := log.Count(map[string]interface{}{"ip": ip, "status": false}, func(db *gorm.DB) *gorm.DB {
		return model.Today(db, "created_at")
	})
	if err != nil {
		return err
	}
	if count > 10 {
		return errors.New("当前登陆错误次数过多，已被封禁")
	}
	return nil
}

// NewToken 进行token生成
func (l *LoginLogic) NewToken(user models.User) (string, error) {

	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Unix() + l.svcCtx.Config.GetInt64("auth.access_expire")
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
	return token.SignedString([]byte(l.svcCtx.Config.GetString("auth.access_secret")))

}

package service

import (
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
	"gorm.io/gorm"
	"ums/consts"
	"ums/errors"
	"ums/model"
	"ums/tools"
	"ums/types"
)

func CurUser(ctx *gin.Context) (*model.User, error) {
	user := model.CurUser(ctx)
	if user.ID == 0 {
		return nil, errors.New("获取个人信息失败")
	}
	return &user, nil
}

func PageUser(ctx *gin.Context, in *types.PageUserRequest) ([]model.User, int64, error) {
	user := model.User{}
	return user.Page(ctx, in.Page, in.Count, in, func(db *gorm.DB) *gorm.DB {
		return db.Where("team_id in ?", model.CurUserTeamIds(ctx))
	})
}

func GetUser(ctx *gin.Context, in *types.GetUserRequest) (model.User, error) {
	user := model.User{}
	return user, user.OneByID(ctx, in.ID)
}

func AddUser(ctx *gin.Context, in *types.AddUserRequest) error {
	user := model.User{}
	if in.Nickname == "" {
		in.Nickname = in.Name
	}
	if !tools.InList(model.CurUserTeamIds(ctx), in.TeamID) {
		return errors.NotAddTeamUserError
	}
	if copier.Copy(&user, in) != nil {
		return errors.AssignError
	}

	if err := user.Create(ctx); err != nil {
		return err
	}
	return nil
}

func UpdateUser(ctx *gin.Context, in *types.UpdateUserRequest) error {
	if in.ID == 1 { //超级管理员不允许修改所在部门和角色
		in.RoleID = 0
		in.TeamID = 0
	}
	user := model.User{}
	if user.OneByID(ctx, in.ID) != nil {
		return errors.DBNotFoundError
	}

	if !tools.InList(model.CurUserTeamIds(ctx), user.TeamID) {
		return errors.NotEditTeamUserError
	}

	if copier.Copy(&user, in) != nil {
		return errors.AssignError
	}
	return user.Update(ctx)
}

func DeleteUser(ctx *gin.Context, in *types.DeleteUserRequest) error {
	if in.ID == 1 {
		return errors.SuperAdminDelError
	}
	user := model.User{}
	if user.OneByID(ctx, in.ID) != nil {
		return errors.DBNotFoundError
	}

	if !tools.InList(model.CurUserTeamIds(ctx), user.TeamID) {
		return errors.NotDelTeamUserError
	}

	return user.Delete(ctx)
}

func UserLogin(ctx *gin.Context, in *types.UserLoginRequest) (resp *types.UserLoginResponse, err error) {
	// 判断是否登陆次数过多
	if err = LoginIpLimit(ctx); err != nil {
		return nil, err
	}

	defer func() {
		_ = AddLoginLog(ctx, in.Phone, err)
		_ = DisableUser(ctx, in.Phone, err)
	}()

	user := model.User{}
	if in.Password, err = ctx.Rsa(consts.RsaPrivate).Decode(in.Password); err != nil {
		err = errors.RsaPasswordError
		return
	}
	if user.Scan(ctx, "phone = ?", in.Phone) != nil {
		err = errors.UserNameNotFoundError
		return
	}
	if !*user.Status {
		err = errors.UserNameDisableError
		return
	}
	if !tools.CompareHashPwd(user.Password, in.Password) {
		err = errors.PasswordError
		return
	}
	resp, err = GenToken(ctx, user.ID)
	return
}

// LoginIpLimit 判断是否登陆错误次数错误过多
func LoginIpLimit(ctx *gin.Context) error {
	if !consts.LoginLimit.Enable {
		return nil
	}
	ip := ctx.Request.Header.Get("x-real-ip")

	// 获取今日登陆次数
	log := model.LoginLog{}
	count, _ := log.Count(ctx, func(db *gorm.DB) *gorm.DB {
		db = db.Where("ip = ? and status = false", ip)
		return tools.Today(db, "created_at")
	})

	// 判断是否超过最大登陆错误次数
	if count >= consts.LoginLimit.IpLimit {
		return errors.IpLimitLoginError
	}
	return nil
}

// DisableUser 根据登陆日志进行用户禁用
func DisableUser(ctx *gin.Context, phone string, err error) error {

	data, _ := err.(*gin.CustomError)
	if data != nil && data.Code != errors.PasswordError.Code {
		return nil
	}

	// 获取用户最近的登陆日志
	log := model.LoginLog{}
	list, _, _ := log.Page(ctx, 1, int(consts.LoginLimit.PasswordErrorLimit), nil, func(db *gorm.DB) *gorm.DB {
		return db.Where("username = ?", phone).Order("created_at desc")
	})

	if len(list) < int(consts.LoginLimit.PasswordErrorLimit) {
		return nil
	}

	// 存在非账号密码错误则进行封禁
	for _, item := range list {
		if item.Code != errors.PasswordError.Code {
			return nil
		}
	}

	//进行账号封禁
	user := model.User{}
	return user.Disable(ctx, phone)
}

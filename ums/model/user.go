package model

import (
	"encoding/json"
	"fmt"
	"github.com/limeschool/gin"
	"gorm.io/gorm"
	"time"
	"ums/errors"
	"ums/meta"
	"ums/tools"
	"ums/tools/lock"
)

type User struct {
	gin.BaseModel
	TeamID     int64  `json:"team_id"`
	Team       Team   `json:"team" gorm:"->"`
	RoleID     int64  `json:"role_id"`
	Role       Role   `json:"role" gorm:"->"`
	Name       string `json:"name"`
	Sex        *bool  `json:"sex,omitempty"`
	Phone      string `json:"phone"`
	Password   string `json:"password,omitempty"  gorm:"->:false;<-:create,update"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email,omitempty"`
	Status     *bool  `json:"status,omitempty"`
	LastLogin  int64  `json:"last_login"`
	Operator   string `json:"operator"`
	OperatorID int64  `json:"operator_id"`
}

var userLockKey = "user_lock"

func (u User) Table() string {
	return "user"
}

func (u User) cacheKey(id int64) string {
	return fmt.Sprintf("ums_user_info_%v", id)
}

func (u *User) OneByCache(ctx *gin.Context, key string) (bool, error) {
	resByte, err := cache(ctx).Get(ctx, key).Bytes()
	if err != nil {
		return false, err
	}
	if len(resByte) == 0 {
		return false, nil
	}
	if err = json.Unmarshal(resByte, u); err != nil {
		return false, nil
	}

	if u.ID == 0 {
		return true, gorm.ErrRecordNotFound
	}

	return true, nil
}

// OneByID 通过id查询用户信息
func (u *User) OneByID(ctx *gin.Context, id int64) error {

	if is, err := u.OneByCache(ctx, u.cacheKey(id)); is {
		return transferErr(err)
	}

	// 加锁,防止缓存击穿
	rl := lock.NewLock(ctx, userLockKey)
	rl.Acquire()
	defer rl.Release()

	// 获取锁之后重新查询缓存
	if is, err := u.OneByCache(ctx, u.cacheKey(id)); is {
		return err
	}

	// 没有则从数据库获取
	db := database(ctx).Table(u.Table()).Preload("Role").Preload("Team")
	if err := db.First(u, "id=?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cache(ctx).Set(ctx, u.cacheKey(id), "{}", 5*time.Minute)
		}
		return transferErr(err)
	}

	// 进行数据缓存
	b, _ := json.Marshal(u)
	cache(ctx).Set(ctx, u.cacheKey(id), string(b), 2*time.Hour)
	return nil
}

// Scan 查询全部字段信息包括密码
func (u *User) Scan(ctx *gin.Context, conds ...interface{}) error {
	m := map[string]any{}
	if err := database(ctx).Table(u.Table()).First(u, conds...).Scan(&m).Error; err != nil {
		return transferErr(err)
	}
	u.Password, _ = m["password"].(string)
	return nil
}

// Page 查询分页数据
func (u *User) Page(ctx *gin.Context, page, count int, m interface{}, fs ...callback) ([]User, int64, error) {
	var list []User
	var total int64

	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	db = execCallback(db, fs...)

	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}

	db = db.Preload("Team").Preload("Role")
	return list, total, transferErr(db.Offset((page - 1) * count).Limit(count).Find(&list).Error)
}

// CurUser 查询当前用户信息
func CurUser(ctx *gin.Context) User {
	user := User{}
	_ = user.OneByID(ctx, meta.UserID(ctx))
	return user
}

// Create 创建用户信息
func (u *User) Create(ctx *gin.Context) error {

	// 操作者信息
	user := CurUser(ctx)
	u.OperatorID = user.ID
	u.Operator = user.Name

	// 加密密码
	u.Password, _ = tools.ParsePwd(u.Password)

	// 执行新增
	return transferErr(database(ctx).Table(u.Table()).Create(u).Error)
}

// Update 更新用户信息
func (u *User) Update(ctx *gin.Context) error {

	// 延迟双删
	delayDelCache(ctx, u.cacheKey(u.ID))

	// 操作者信息
	user := CurUser(ctx)
	u.Operator = user.Name
	u.OperatorID = user.ID

	// 是否更新密码
	if u.Password != "" {
		u.Password, _ = tools.ParsePwd(u.Password)
	}

	// 执行更新
	return transferErr(database(ctx).Table(u.Table()).Updates(&u).Error)
}

func (u *User) Disable(ctx *gin.Context, phone string) error {

	// 延迟双删
	delayDelCache(ctx, u.cacheKey(u.ID))

	// 进行账号禁用
	return transferErr(database(ctx).Table(u.Table()).Where("phone=?", phone).
		Update("status", false).Error)
}

func (u *User) Delete(ctx *gin.Context) error {

	// 延迟双删
	delayDelCache(ctx, u.cacheKey(u.ID))

	// 进行账号删除
	return transferErr(database(ctx).Table(u.Table()).Delete(u, "id=?", u.ID).Error)
}

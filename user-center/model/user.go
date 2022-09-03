package model

import (
	"encoding/json"
	"fmt"
	"github.com/limeschool/gin"
	"time"
	"ums/consts"
	"ums/meta"
	"ums/tools"
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
	Password   string `json:"password,omitempty"  gorm:"->:false;<-:create"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email,omitempty"`
	Status     *bool  `json:"status,omitempty"`
	LastLogin  int64  `json:"last_login"`
	Operator   string `json:"operator"`
	OperatorID int64  `json:"operator_id"`
}

func (u User) Table() string {
	return "user"
}

func (u User) cacheKey(id int64) string {
	return fmt.Sprintf("ums_user_info_%v", id)
}

func (u *User) OneByID(ctx *gin.Context, id int64) error {
	// 先从redis 中获取
	resByte, err := ctx.Redis(consts.REDIS).Get(ctx, u.cacheKey(id)).Bytes()
	if err == nil && string(resByte) != "" && json.Unmarshal(resByte, u) == nil {
		return nil
	}
	// 没有则从数据库获取
	db := database(ctx).Table(u.Table()).Preload("Role").Preload("Team")

	if err = db.First(u, "id = ?", id).Error; err != nil {
		return err
	}
	// 进行数据缓存
	b, _ := json.Marshal(u)
	ctx.Redis(consts.REDIS).Set(ctx, u.cacheKey(id), string(b), 2*time.Hour)
	return nil
}

func (u *User) Scan(ctx *gin.Context, conds ...interface{}) error {
	m := map[string]any{}
	if err := database(ctx).Table(u.Table()).First(u, conds...).Scan(&m).Error; err != nil {
		return err
	}
	u.Password, _ = m["password"].(string)
	return nil
}

func (u *User) One(ctx *gin.Context, conds ...interface{}) error {
	if err := database(ctx).Table(u.Table()).
		First(u, conds...).Error; err != nil {
		return err
	}
	b, _ := json.Marshal(u)
	ctx.Redis(consts.REDIS).Set(ctx, u.cacheKey(u.ID), string(b), 2*time.Hour)
	return nil
}

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
	return list, total, db.Offset((page - 1) * count).Limit(count).Find(&list).Error
}

func CurUser(ctx *gin.Context) User {
	user := User{}
	_ = user.OneByID(ctx, meta.UserID(ctx))
	return user
}

func (u *User) Create(ctx *gin.Context) error {
	user := CurUser(ctx)
	u.OperatorID = user.ID
	u.Operator = user.Name
	if u.Password != "" {
		u.Password, _ = tools.ParsePwd(u.Password)
	}
	return database(ctx).Table(u.Table()).Create(u).Error
}

func (u *User) UpdateByID(ctx *gin.Context) error {
	defer func() {
		tools.DelRedis(ctx, u.cacheKey(u.ID))
	}()
	ctx.Redis(consts.REDIS).Del(ctx, u.cacheKey(u.ID))

	user := CurUser(ctx)
	u.Operator = user.Name
	u.OperatorID = user.ID
	if u.Password != "" {
		u.Password, _ = tools.ParsePwd(u.Password)
	}
	return database(ctx).Table(u.Table()).Updates(&u).Error
}

func (u *User) Disable(ctx *gin.Context, phone string) error {
	defer func() {
		tools.DelRedis(ctx, u.cacheKey(u.ID))
	}()
	ctx.Redis(consts.REDIS).Del(ctx, u.cacheKey(u.ID))
	return database(ctx).Table(u.Table()).Where("phone = ?", phone).Update("status", false).Error
}

func (u *User) DeleteByID(ctx *gin.Context, id int64) error {
	defer func() {
		tools.DelRedis(ctx, u.cacheKey(id))
	}()
	ctx.Redis(consts.REDIS).Del(ctx, u.cacheKey(id))
	return database(ctx).Table(u.Table()).Delete(u, "id= ?", id).Error
}

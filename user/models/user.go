package models

import (
	"context"
	"devops/common/meta"
	"devops/common/model"
	"devops/common/tools"
	"time"
)

type User struct {
	RoleID     int    `json:"role_id"`
	TeamID     int    `json:"team_id"`
	RoleName   string `json:"role_name" gorm:"->"`
	Keyword    string `json:"keyword" gorm:"->"`
	TeamName   string `json:"team_name" gorm:"->"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	Password   string `json:"password,omitempty"`
	Status     *bool  `json:"status" `
	Operator   string `json:"operator"`
	OperatorID int64  `json:"operator_id"`
	model.BaseModel
}

// 如果没有自定义缓存。那就应该按照ID来进行缓存数据

func (u User) Table() string {
	return "user"
}

func (u *User) OneByID() error {
	return database().Table(u.Table()).
		Select("user.*,team.name team_name,role.name role_name,role.keyword").
		Joins("left join role on role.id = user.role_id").
		Joins("left join team on team.id = user.team_id").
		First(&u, u.ID).Error
}

func (u *User) One(query interface{}) error {
	db := database().Table(u.Table()).
		Select("user.*,team.name team_name,role.name role_name,role.keyword").
		Joins("left join role on role.id = user.role_id").
		Joins("left join team on team.id = user.team_id")
	db = model.SqlWhere(db, query)
	return db.First(&u).Error
}

func (u *User) OneByCall(f callback) error {
	db := database().Table(u.Table()).
		Select("user.*,team.name team_name,role.name role_name,role.keyword").
		Joins("left join role on role.id = user.role_id").
		Joins("left join team on team.id = user.team_id")
	db = f(db)
	return db.First(&u).Error
}

func (u *User) Create(ctx context.Context) error {
	u.OperatorID = meta.UserId(ctx)
	u.Operator = meta.UserName(ctx)
	return database().Table(u.Table()).Create(&u).Error
}

func (u *User) Page(query interface{}, page, count int64) ([]User, int64, error) {
	var list []User
	var total int64
	db := database().Table(u.Table()).
		Select("user.*,team.name team_name,role.name role_name,role.keyword").
		Joins("left join role on role.id = user.role_id").
		Joins("left join team on team.id = user.team_id")
	db = model.SqlWhere(db, query, "page", "count")
	db.Count(&total)
	db = db.Offset(int((page - 1) * count)).Limit(int(count))
	return list, total, db.Find(&list).Error
}

func (u *User) All(query interface{}) ([]User, int64, error) {
	var list []User
	var total int64
	db := database().Table(u.Table()).
		Select("user.*,team.name team_name,role.name role_name,role.keyword").
		Joins("left join role on role.id = user.role_id").
		Joins("left join team on team.id = user.team_id")
	db = model.SqlWhere(db, query)
	db.Count(&total)
	return list, total, db.Find(&list).Error
}

func (u *User) Update(ctx context.Context) error {
	u.OperatorID = meta.UserId(ctx)
	u.Operator = meta.UserName(ctx)
	return database().Table(u.Table()).Updates(u).Error
}

func (u *User) UpdateByFields(ctx context.Context, c interface{}, m interface{}) error {
	fields := tools.ToMap(m)
	db := database().Table(u.Table())
	db = model.SqlWhere(db, c)
	fields["created_at"] = time.Now().Unix()
	fields["updated_at"] = time.Now().Unix()
	fields["operator"] = meta.UserName(ctx)
	fields["operator_id"] = meta.UserId(ctx)
	return db.Updates(m).Error
}

func (u *User) UpdateByID(ctx context.Context, m interface{}) error {
	fields := tools.ToMap(m)
	fields["created_at"] = time.Now().Unix()
	fields["updated_at"] = time.Now().Unix()
	fields["operator"] = meta.UserName(ctx)
	fields["operator_id"] = meta.UserId(ctx)
	return database().Table(u.Table()).Where("id = ?", u.ID).Updates(m).Error
}

func (u *User) DeleteByID(ctx context.Context) error {
	u.OperatorID = meta.UserId(ctx)
	u.Operator = meta.UserName(ctx)
	return database().Table(u.Table()).Delete(&u).Error
}

func (u *User) Delete(ctx context.Context, m interface{}) error {
	u.OperatorID = meta.UserId(ctx)
	u.Operator = meta.UserName(ctx)
	db := database().Table(u.Table())
	db = model.SqlWhere(db, m)
	return db.Delete(&u).Error
}

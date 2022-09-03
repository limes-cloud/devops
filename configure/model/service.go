package model

import (
	"configure/meta"
	"github.com/limeschool/gin"
)

type Service struct {
	gin.BaseModel
	Keyword     string  `json:"keyword"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Operator    string  `json:"operator"`
	OperatorId  int64   `json:"operator_id"`
	EnvIds      []int64 `json:"env_ids" gorm:"-"`
}

func (u *Service) Table() string {
	return "service"
}

func (u *Service) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return database(ctx).Table(u.Table()).Create(&u).Error
}

func (u *Service) OneById(ctx *gin.Context, id int64) error {
	return database(ctx).Table(u.Table()).First(u, "id = ?", id).Error
}

func (u *Service) One(ctx *gin.Context, conds ...interface{}) error {
	return database(ctx).Table(u.Table()).First(u, conds...).Error
}

func (u *Service) All(ctx *gin.Context, m interface{}) ([]Service, error) {
	var list []Service
	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	return list, db.Find(&list).Error
}

func (u *Service) UpdateByID(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return database(ctx).Table(u.Table()).Updates(u).Error
}

func (u *Service) DeleteByID(ctx *gin.Context, id int64) error {
	return database(ctx).Table(u.Table()).Where("id = ?", id).Delete(&u).Error
}

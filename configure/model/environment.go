package model

import (
	"configure/meta"
	"github.com/limeschool/gin"
)

type Environment struct {
	gin.BaseModel
	Keyword     string  `json:"keyword,omitempty"`
	Name        string  `json:"name,omitempty"`
	Drive       string  `json:"drive,omitempty"`
	Config      string  `json:"config,omitempty"`
	Prefix      string  `json:"prefix,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *bool   `json:"status,omitempty"`
	Token       string  `json:"token,omitempty"`
	Operator    string  `json:"operator,omitempty"`
	OperatorId  int64   `json:"operator_id,omitempty"`
}

func (u *Environment) Table() string {
	return "environment"
}

func (u *Environment) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return database(ctx).Table(u.Table()).Create(&u).Error
}

func (u *Environment) OneById(ctx *gin.Context, id int64) error {
	return database(ctx).Table(u.Table()).First(u, "id = ?", id).Error
}

func (u *Environment) One(ctx *gin.Context, conds ...interface{}) error {
	return database(ctx).Table(u.Table()).First(u, conds...).Error
}

func (u *Environment) AllFilter(ctx *gin.Context, conds ...interface{}) ([]Environment, error) {
	var list []Environment
	db := database(ctx).Table(u.Table()).Select("id,keyword,name").Where("status = true")
	return list, db.Find(&list, conds...).Error
}

func (u *Environment) All(ctx *gin.Context, m interface{}) ([]Environment, error) {
	var list []Environment
	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	return list, db.Find(&list).Error
}

func (u *Environment) UpdateByID(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return database(ctx).Table(u.Table()).Updates(u).Error
}

func (u *Environment) DeleteByID(ctx *gin.Context, id int64) error {
	return database(ctx).Table(u.Table()).Where("id = ?", id).Delete(&u).Error
}

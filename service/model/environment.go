package model

import (
	"github.com/limeschool/gin"
	"service/meta"
)

type Environment struct {
	gin.BaseModel
	Keyword      string  `json:"keyword,omitempty"`
	Name         string  `json:"name,omitempty"`
	Description  *string `json:"description,omitempty"`
	DcHost       *string `json:"dc_host"`
	DcToken      *string `json:"dc_token"`
	K8sHost      *string `json:"k8s_host"`
	K8sToken     *string `json:"k8s_token"`
	K8sNamespace *string `json:"k8s_namespace"`
	Status       *bool   `json:"status,omitempty"`
	Operator     string  `json:"operator,omitempty"`
	OperatorId   int64   `json:"operator_id,omitempty"`
}

func (u *Environment) Table() string {
	return "environment"
}

func (u *Environment) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Create(&u).Error)
}

func (u *Environment) OneById(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).First(u, "id = ?", id).Error)
}

func (u *Environment) OneByKeyword(ctx *gin.Context, key string) error {
	return transferErr(database(ctx).Table(u.Table()).First(u, "keyword = ?", key).Error)
}

func (u *Environment) One(ctx *gin.Context, cond ...interface{}) error {
	return transferErr(database(ctx).Table(u.Table()).First(u, cond...).Error)
}

func (u *Environment) AllFilter(ctx *gin.Context, cond ...any) ([]Environment, error) {
	var list []Environment
	db := database(ctx).Table(u.Table()).Select("id,keyword,name").Where("status = true")
	return list, transferErr(db.Find(&list, cond...).Error)
}

func (u *Environment) All(ctx *gin.Context, m any) ([]Environment, error) {
	var list []Environment
	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	return list, transferErr(db.Find(&list).Error)
}

func (u *Environment) UpdateByID(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Updates(u).Error)
}

func (u *Environment) DeleteByID(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).Where("id = ?", id).Delete(&u).Error)
}

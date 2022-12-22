package model

import (
	"github.com/limeschool/gin"
	"service/meta"
)

type Network struct {
	gin.BaseModel
	ServiceName string `json:"service_name" gorm:"-"`
	EnvName     string `json:"env_name" gorm:"-"`
	SrvID       int64  `json:"srv_id"`
	EnvID       int64  `json:"env_id"`
	Host        string `json:"host"`
	Cert        string `json:"cert"`
	Key         string `json:"key"`
	Redirect    bool   `json:"redirect"`
	Operator    string `json:"operator,omitempty"`
	OperatorId  int64  `json:"operator_id,omitempty"`
}

func (u *Network) Table() string {
	return "network"
}

func (u *Network) OneById(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).First(u, "id = ?", id).Error)
}

func (u *Network) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Create(&u).Error)
}

func (u *Network) Page(ctx *gin.Context, page, count int, m interface{}, fs ...callback) ([]Network, int64, error) {
	var list []Network
	var total int64

	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	db = execCallback(db, fs...)
	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}

	return list, total, transferErr(db.Offset((page - 1) * count).Limit(count).Find(&list).Error)
}

func (u *Network) Count(ctx *gin.Context, m interface{}, fs ...callback) int64 {
	var total int64
	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	db = execCallback(db, fs...)
	db.Count(&total)
	return total
}

func (u *Network) UpdateByID(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Updates(u).Error)
}

func (u *Network) DeleteByID(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).Where("id = ?", id).Delete(&u).Error)
}

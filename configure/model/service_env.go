package model

import (
	"github.com/limeschool/gin"
)

type ServiceEnv struct {
	ID         int64  `json:"id"`
	EnvId      int64  `json:"env_id"`
	SrvId      int64  `json:"srv_id"`
	Operator   string `json:"operator"`
	OperatorId int64  `json:"operator_id"`
}

func (u *ServiceEnv) Table() string {
	return "service_env"
}

func (u *ServiceEnv) CreateAll(ctx *gin.Context, list []ServiceEnv) error {
	return transferErr(database(ctx).Table(u.Table()).Create(&list).Error)
}

func (u *ServiceEnv) Delete(ctx *gin.Context, conds ...interface{}) error {
	return transferErr(database(ctx).Table(u.Table()).Delete(&u, conds...).Error)
}

func (u *ServiceEnv) All(ctx *gin.Context, conds ...interface{}) ([]ServiceEnv, error) {
	var list []ServiceEnv
	db := database(ctx).Table(u.Table())
	return list, transferErr(db.Find(&list, conds...).Error)
}

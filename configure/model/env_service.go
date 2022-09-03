package model

import (
	"github.com/limeschool/gin"
)

type EnvService struct {
	ID         int64  `json:"id"`
	EnvId      int64  `json:"env_id"`
	SrvId      int64  `json:"srv_id"`
	Operator   string `json:"operator"`
	OperatorId int64  `json:"operator_id"`
}

func (u *EnvService) Table() string {
	return "env_service"
}

func (u *EnvService) CreateAll(ctx *gin.Context, list []EnvService) error {
	return database(ctx).Table(u.Table()).Create(&list).Error
}

func (u *EnvService) Delete(ctx *gin.Context, conds ...interface{}) error {
	return database(ctx).Table(u.Table()).Delete(&u, conds...).Error
}

func (u *EnvService) All(ctx *gin.Context, conds ...interface{}) ([]EnvService, error) {
	var list []EnvService
	db := database(ctx).Table(u.Table())
	return list, db.Find(&list, conds...).Error
}

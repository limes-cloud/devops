package model

import (
	"configure/meta"
	"github.com/limeschool/gin"
)

type TemplateLog struct {
	gin.CreateModel
	SrvId       int64  `json:"srv_id"`
	EnvId       int64  `json:"env_id"`
	Config      string `json:"config"`
	Description string `json:"description"`
	Operator    string `json:"operator"`
	OperatorId  int64  `json:"operator_id"`
}

func (u *TemplateLog) Table() string {
	return "template_log"
}

func (u *TemplateLog) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	err := database(ctx).Table(u.Table()).Create(&u).Error
	// 只保留15个版本
	temp := Template{}
	if database(ctx).Table(u.Table()).Offset(15).Limit(1).Order("id desc").First(&temp).Error == nil {
		database(ctx).Table(u.Table()).Where("id <= ?", temp.ID).Delete(&Template{})
	}
	return err
}

func (u *TemplateLog) All(ctx *gin.Context, m interface{}) ([]TemplateLog, error) {
	var list []TemplateLog
	db := database(ctx).Table(u.Table()).Order("created_at desc")
	db = gin.GormWhere(db, u.Table(), m)
	return list, db.Select("id,description,operator,operator_id,created_at").Find(&list).Error
}

func (u *TemplateLog) OneById(ctx *gin.Context, id int64) error {
	return database(ctx).Table(u.Table()).First(u, "id = ?", id).Error
}

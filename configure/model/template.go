package model

import (
	"configure/meta"
	"github.com/limeschool/gin"
	"gorm.io/gorm"
)

type Template struct {
	gin.BaseModel
	ServiceKeyword string  `json:"service_keyword"`
	Content        string  `json:"content"`
	Version        string  `json:"version"`
	IsUse          bool    `json:"is_use"`
	Description    *string `json:"description"`
	Operator       string  `json:"operator"`
	OperatorId     int64   `json:"operator_id"`
}

func (u *Template) Table() string {
	return "template"
}

func (u *Template) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	err := database(ctx).Table(u.Table()).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id != 0").Update("is_use", false).Error; err != nil {
			return err
		}
		return tx.Create(&u).Error
	})
	// 只保留15个版本
	temp := Template{}
	if database(ctx).Table(u.Table()).Offset(15).Limit(1).Order("id desc").First(&temp).Error == nil {
		database(ctx).Table(u.Table()).Where("id <= ?", temp.ID).Delete(&Template{})
	}
	return err
}

func (u *Template) OneById(ctx *gin.Context, id int64) error {
	return database(ctx).Table(u.Table()).First(u, "id = ?", id).Error
}

func (u *Template) One(ctx *gin.Context, conds ...interface{}) error {
	return database(ctx).Table(u.Table()).First(u, conds...).Error
}

func (u *Template) OneBy(ctx *gin.Context, conds ...interface{}) error {
	return database(ctx).Table(u.Table()).First(u, conds...).Error
}

func (u *Template) All(ctx *gin.Context, m interface{}) ([]Template, error) {
	var list []Template
	db := database(ctx).Table(u.Table()).Order("created_at desc")
	db = gin.GormWhere(db, u.Table(), m)
	return list, db.Select("id,service_keyword,version,is_use,description,operator,operator_id,created_at").Find(&list).Error
}

func (u *Template) UpdateVersionByID(ctx *gin.Context) error {
	user := meta.User(ctx)

	return database(ctx).Table(u.Table()).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id !=0").Update("is_use", false).Error; err != nil {
			return err
		}
		u.OperatorId = user.UserId
		u.Operator = user.UserName
		return tx.Where("id = ?", u.ID).Update("is_use", true).Error
	})

}

func (u *Template) DeleteByID(ctx *gin.Context, id int64) error {
	return database(ctx).Table(u.Table()).Where("id = ?", id).Delete(&u).Error
}

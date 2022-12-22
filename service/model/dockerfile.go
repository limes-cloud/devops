package model

import (
	"github.com/limeschool/gin"
	"service/meta"
)

type Dockerfile struct {
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Template   string `json:"template"`
	Operator   string `json:"operator,omitempty"`
	OperatorId int64  `json:"operator_id,omitempty"`
	gin.BaseModel
}

func (u *Dockerfile) Table() string {
	return "dockerfile_template"
}

func (u *Dockerfile) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Create(&u).Error)
}

func (u *Dockerfile) OneById(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).First(u, "id = ?", id).Error)
}

func (u *Dockerfile) Page(ctx *gin.Context, page, count int, m interface{}, fs ...callback) ([]Dockerfile, int64, error) {
	var list []Dockerfile
	var total int64

	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	db = execCallback(db, fs...).Order("created_at desc")
	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}

	return list, total, transferErr(db.Offset((page - 1) * count).Limit(count).Find(&list).Error)
}

func (u *Dockerfile) All(ctx *gin.Context, m any, fs ...callback) ([]Dockerfile, error) {
	var list []Dockerfile
	db := database(ctx).Table(u.Table())
	db = execCallback(db, fs...)
	db = gin.GormWhere(db, u.Table(), m)
	return list, transferErr(db.Find(&list).Error)
}

func (u *Dockerfile) UpdateByID(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Updates(u).Error)
}

func (u *Dockerfile) DeleteByID(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).Where("id = ?", id).Delete(&u).Error)
}

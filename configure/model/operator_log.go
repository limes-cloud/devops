package model

import (
	"configure/meta"
	"github.com/limeschool/gin"
)

type OperatorLog struct {
	ID             int64  `json:"id"`
	ServiceKeyword string `json:"service_keyword"`
	ServiceName    string `json:"service_name"`
	Title          string `json:"title"`
	Content        string `json:"content"`
	Operator       string `json:"operator"`
	OperatorId     int64  `json:"operator_id"`
}

func (u *OperatorLog) Table() string {
	return "service"
}

func (u *OperatorLog) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return database(ctx).Table(u.Table()).Create(&u).Error
}

func (u *OperatorLog) OneById(ctx *gin.Context, id int64) error {
	return database(ctx).Table(u.Table()).First(u, "id = ?", id).Error
}

func (u *OperatorLog) One(ctx *gin.Context, conds ...interface{}) error {
	return database(ctx).Table(u.Table()).First(u, conds...).Error
}

func (u *OperatorLog) Page(ctx *gin.Context, page, count int, m interface{}) ([]OperatorLog, int64, error) {
	var list []OperatorLog
	var total int64
	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}
	return list, total, db.Offset((page - 1) * count).Limit(count).Find(&list).Error
}

func (u *OperatorLog) UpdateByID(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return database(ctx).Table(u.Table()).Updates(u).Error
}

func (u *OperatorLog) DeleteByID(ctx *gin.Context, id int64) error {
	return database(ctx).Table(u.Table()).Where("id = ?", id).Delete(&u).Error
}

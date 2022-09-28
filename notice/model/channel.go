package model

import (
	"github.com/limeschool/gin"
	"notice/meta"
)

type Channel struct {
	Name       string `json:"name"`
	Config     string `json:"config,omitempty"`
	Status     *bool  `json:"status"`
	Operator   string `json:"operator"`
	OperatorID int64  `json:"operator_id"`
	gin.BaseModel
}

func (c *Channel) Table() string {
	return "channel"
}

func (c *Channel) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	c.OperatorID = user.UserId
	c.Operator = user.UserName
	return database(ctx).Table(c.Table()).Create(&c).Error
}

func (c *Channel) OneById(ctx *gin.Context, id int64) error {
	return database(ctx).Table(c.Table()).First(c, "id = ?", id).Error
}

func (c *Channel) All(ctx *gin.Context, m interface{}, fs ...callback) ([]Channel, error) {
	var list []Channel
	db := database(ctx).Table(c.Table())
	db = execCallback(db, fs...)
	db = gin.GormWhere(db, c.Table(), m)
	return list, db.Find(&list).Error
}

func (c *Channel) Update(ctx *gin.Context) error {
	user := meta.User(ctx)
	c.Operator = user.UserName
	c.OperatorID = user.UserId
	return database(ctx).Table(c.Table()).Updates(c).Error
}

func (c *Channel) DeleteByID(ctx *gin.Context, id int64) error {
	return database(ctx).Table(c.Table()).Where("id = ?", id).Delete(&c).Error
}

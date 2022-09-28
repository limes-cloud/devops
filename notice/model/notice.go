package model

import (
	"github.com/limeschool/gin"
	"notice/meta"
)

type Notice struct {
	gin.BaseModel
	Cid        string        `json:"cid"`
	Title      string        `json:"title"`
	Rule       string        `json:"rule"`
	Value      int64         `json:"value"`
	UserIds    string        `json:"user_ids"`
	Status     *bool         `json:"status"`
	Operator   string        `json:"operator"`
	OperatorID int64         `json:"operator_id"`
	Channels   []Channel     `json:"channels,omitempty" gorm:"-"`
	Users      []interface{} `json:"users,omitempty" gorm:"-"`
}

func (n *Notice) Table() string {
	return "notice"
}

func (n *Notice) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	n.OperatorID = user.UserId
	n.Operator = user.UserName
	return database(ctx).Table(n.Table()).Create(&n).Error
}

func (n *Notice) OneById(ctx *gin.Context, id int64) error {
	return database(ctx).Table(n.Table()).First(n, "id = ?", id).Error
}

func (n *Notice) Page(ctx *gin.Context, page, count int, m interface{}) ([]Notice, int64, error) {
	var list []Notice
	var total int64
	db := database(ctx).Table(n.Table())
	db = gin.GormWhere(db, n.Table(), m)
	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}
	return list, total, db.Offset((page - 1) * count).Limit(count).Find(&list).Error
}

func (n *Notice) Update(ctx *gin.Context) error {
	user := meta.User(ctx)
	n.Operator = user.UserName
	n.OperatorID = user.UserId
	return database(ctx).Table(n.Table()).Updates(n).Error
}

func (n *Notice) DeleteByID(ctx *gin.Context, id int64) error {
	return database(ctx).Table(n.Table()).Where("id = ?", id).Delete(&n).Error
}

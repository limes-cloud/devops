package model

import "github.com/limeschool/gin"

type Log struct {
	gin.CreateModel
	Cid      string `json:"cid"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Rule     string `json:"rule"`
	Users    string `json:"users"`
	Channels string `json:"channels"`
}

func (n *Log) Table() string {
	return "log"
}

func (n *Log) Create(ctx *gin.Context) error {
	return database(ctx).Table(n.Table()).Create(&n).Error
}

func (n *Log) Page(ctx *gin.Context, page, count int, m interface{}) ([]Log, int64, error) {
	var list []Log
	var total int64
	db := database(ctx).Table(n.Table())
	db = gin.GormWhere(db, n.Table(), m)
	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}
	return list, total, db.Offset((page - 1) * count).Limit(count).Find(&list).Error
}

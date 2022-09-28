package model

import "github.com/limeschool/gin"

type NoticeChannel struct {
	ID        int64 `json:"id"`
	ChannelID int64 `json:"channel_id"`
	NoticeID  int64 `json:"notice_id"`
}

func (c *NoticeChannel) Table() string {
	return "notice_channel"
}

func (c *NoticeChannel) CreateAll(ctx *gin.Context, list []NoticeChannel) error {
	return database(ctx).Table(c.Table()).Create(list).Error
}

func (c *NoticeChannel) Delete(ctx *gin.Context, conds ...interface{}) error {
	return database(ctx).Table(c.Table()).Delete(c, conds...).Error
}

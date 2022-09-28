package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
	"gorm.io/gorm"
	"notice/errors"
	"notice/model"
	"notice/types"
)

func PageNotice(ctx *gin.Context, in *types.GetNoticeRequest) ([]model.Notice, int64, error) {
	m := model.Notice{}
	list, total, err := m.Page(ctx, in.Page, in.Count, in)
	if err != nil {
		return nil, 0, errors.DBError
	}
	url := ctx.Config.GetString("ums-addr")
	umsToken := ctx.Config.GetString("ums-token")
	for key, item := range list {
		// 获取用户信息
		var users []interface{}
		var userIds []int
		_ = json.Unmarshal([]byte(item.UserIds), &userIds)
		for _, userId := range userIds {
			result := gin.Response{}
			if err = ctx.Http().Option(func(request *resty.Request) *resty.Request {
				return request.SetHeader("token", umsToken)
			}).Get(fmt.Sprintf("%v/user?id=%v", url, userId)).Result(&result); err == nil {
				if result.Code == 200 {
					users = append(users, result.Data)
				}
			}
		}

		//获取通道信息
		ch := model.Channel{}
		channels, _ := ch.All(ctx, nil, func(db *gorm.DB) *gorm.DB {
			return db.Select("id,name").
				Where("id in (select channel_id from notice_channel where notice_id = ?)", item.ID)
		})

		list[key].Channels = channels
		list[key].Users = users
	}
	return list, total, nil
}

func AddNotice(ctx *gin.Context, in *types.AddNoticeRequest) error {
	m := model.Notice{}
	if copier.Copy(&m, in) != nil {
		return errors.AssignError
	}
	if m.Create(ctx) != nil {
		return errors.DBError
	}

	mc := model.NoticeChannel{}
	var list []model.NoticeChannel
	for _, cid := range in.ChannelIds {
		list = append(list, model.NoticeChannel{
			ChannelID: cid,
			NoticeID:  m.ID,
		})
	}
	_ = mc.CreateAll(ctx, list)
	return nil
}

func UpdateNotice(ctx *gin.Context, in *types.UpdateNoticeRequest) error {
	m := model.Notice{}
	if copier.Copy(&m, in) != nil {
		return errors.AssignError
	}
	if m.Update(ctx) != nil {
		return errors.DBError
	}
	if len(in.ChannelIds) != 0 {
		mc := model.NoticeChannel{}
		var list []model.NoticeChannel
		for _, cid := range in.ChannelIds {
			list = append(list, model.NoticeChannel{
				ChannelID: cid,
				NoticeID:  m.ID,
			})
		}
		_ = mc.Delete(ctx, "notice_id = ?", in.ID)
		_ = mc.CreateAll(ctx, list)
	}
	return nil
}

func DeleteNotice(ctx *gin.Context, in *types.DeleteNoticeRequest) error {
	m := model.Notice{}
	if m.DeleteByID(ctx, in.ID) != nil {
		return errors.DBError
	}
	return nil
}

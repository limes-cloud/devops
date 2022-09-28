package service

import (
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
	"gorm.io/gorm"
	"notice/errors"
	"notice/model"
	"notice/types"
)

func AllChannel(ctx *gin.Context, in *types.GetChannelRequest) ([]model.Channel, error) {
	m := model.Channel{}
	list, err := m.All(ctx, in)
	if err != nil {
		return nil, errors.DBError
	}
	return list, nil
}

func AllChannelFilter(ctx *gin.Context, in *types.GetChannelRequest) ([]model.Channel, error) {
	m := model.Channel{}
	list, err := m.All(ctx, in, func(db *gorm.DB) *gorm.DB {
		return db.Select("id,name")
	})
	if err != nil {
		return nil, errors.DBError
	}
	return list, nil
}

func AddChannel(ctx *gin.Context, in *types.AddChannelRequest) error {
	m := model.Channel{}
	if copier.Copy(&m, in) != nil {
		return errors.AssignError
	}
	if m.Create(ctx) != nil {
		return errors.DBError
	}

	return nil
}

func UpdateChannel(ctx *gin.Context, in *types.UpdateChannelRequest) error {
	m := model.Channel{}
	if copier.Copy(&m, in) != nil {
		return errors.AssignError
	}
	if m.Update(ctx) != nil {
		return errors.DBError
	}
	return nil
}

func DeleteChannel(ctx *gin.Context, in *types.DeleteChannelRequest) error {
	m := model.Channel{}
	if m.DeleteByID(ctx, in.ID) != nil {
		return errors.DBError
	}
	return nil
}

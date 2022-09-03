package service

import (
	"configure/errors"
	"configure/meta"
	"configure/model"
	"configure/tools"
	"configure/types"
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
)

func AddServiceSystemField(ctx *gin.Context, in *types.AddServiceSystemFieldRequest) error {
	srv := model.ServiceSystemField{ServiceId: in.ServiceId}
	var list []model.ServiceSystemField

	user := meta.User(ctx)
	for _, systemId := range in.FieldIds {
		list = append(list, model.ServiceSystemField{
			ServiceId:     in.ServiceId,
			SystemFieldId: systemId,
			Operator:      user.UserName,
			OperatorId:    user.UserId,
		})
	}

	return srv.CreateAll(ctx, list)
}

func AllServiceSystemField(ctx *gin.Context, in *types.AllServiceSystemFieldRequest) ([]model.SystemField, error) {
	srv := model.ServiceSystemField{}
	list, err := srv.All(ctx, in)
	if err != nil {
		return nil, err
	}
	var fieldIds []int64
	for _, item := range list {
		fieldIds = append(fieldIds, item.SystemFieldId)
	}

	field := model.SystemField{}
	return field.All(ctx, "id in ?", fieldIds)
}

func PageSystemField(ctx *gin.Context, in *types.PageSystemFieldRequest) ([]model.SystemField, int64, error) {
	srv := model.SystemField{}
	return srv.Page(ctx, in.Page, in.Count, in)
}

func AddSystemField(ctx *gin.Context, in *types.AddSystemFieldRequest) error {
	srv := model.SystemField{}
	if copier.Copy(&srv, in) != nil {
		return errors.AssignError
	}
	if tools.DataDup(srv.Create(ctx)) {
		return errors.DulSrvKeywordError
	}
	return nil
}

func UpdateSystemField(ctx *gin.Context, in *types.UpdateSystemFieldRequest) error {
	srv := model.SystemField{}
	if copier.Copy(&srv, in) != nil {
		return errors.AssignError
	}
	err := srv.UpdateByID(ctx)
	if tools.DataDup(err) {
		return errors.DulSrvKeywordError
	}
	return err
}

func DeleteSystemField(ctx *gin.Context, in *types.DeleteSystemFieldRequest) error {
	srv := model.SystemField{}
	return srv.DeleteByID(ctx, in.ID)
}

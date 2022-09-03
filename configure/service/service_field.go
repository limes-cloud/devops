package service

import (
	"configure/errors"
	"configure/model"
	"configure/tools"
	"configure/types"
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
)

func PageServiceField(ctx *gin.Context, in *types.PageServiceFieldRequest) ([]model.ServiceField, int64, error) {
	srv := model.ServiceField{}
	return srv.Page(ctx, in.Page, in.Count, in)
}

func AddServiceField(ctx *gin.Context, in *types.AddServiceFieldRequest) error {
	srv := model.ServiceField{}
	if copier.Copy(&srv, in) != nil {
		return errors.AssignError
	}
	if tools.DataDup(srv.Create(ctx)) {
		return errors.DulSrvKeywordError
	}
	return nil
}

func UpdateServiceField(ctx *gin.Context, in *types.UpdateServiceFieldRequest) error {
	srv := model.ServiceField{}
	if copier.Copy(&srv, in) != nil {
		return errors.AssignError
	}
	err := srv.UpdateByID(ctx)
	if tools.DataDup(err) {
		return errors.DulSrvKeywordError
	}
	return err
}

func DeleteServiceField(ctx *gin.Context, in *types.DeleteServiceFieldRequest) error {
	srv := model.ServiceField{}
	return srv.DeleteByID(ctx, in.ID)
}

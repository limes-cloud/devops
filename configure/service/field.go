package service

import (
	"configure/errors"
	"configure/model"
	"configure/types"
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
)

func PageField(ctx *gin.Context, in *types.PageFieldRequest) ([]model.Field, int64, error) {
	srv := model.Field{}
	return srv.Page(ctx, in.Page, in.Count, in)
}

func AddField(ctx *gin.Context, in *types.AddFieldRequest) error {
	srv := model.Field{}
	if copier.Copy(&srv, in) != nil {
		return errors.AssignError
	}
	return srv.Create(ctx)
}

func UpdateField(ctx *gin.Context, in *types.UpdateFieldRequest) error {
	srv := model.Field{}
	if copier.Copy(&srv, in) != nil {
		return errors.AssignError
	}
	return srv.UpdateByID(ctx)
}

func DeleteField(ctx *gin.Context, in *types.DeleteFieldRequest) error {
	srv := model.Field{}
	return srv.DeleteByID(ctx, in.ID)
}

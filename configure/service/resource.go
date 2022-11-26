package service

import (
	"configure/errors"
	"configure/model"
	"configure/types"
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
)

func PageResource(ctx *gin.Context, in *types.PageResourceRequest) ([]model.Resource, int64, error) {
	srv := model.Resource{}
	return srv.Page(ctx, in.Page, in.Count, in)
}

func AddResource(ctx *gin.Context, in *types.AddResourceRequest) error {
	srv := model.Resource{}
	if copier.Copy(&srv, in) != nil {
		return errors.AssignError
	}
	return srv.Create(ctx)
}

// todo 判断是否存在有人使用，存在的情况下，不允许修改名称。
func UpdateResource(ctx *gin.Context, in *types.UpdateResourceRequest) error {
	srv := model.Resource{}
	if copier.Copy(&srv, in) != nil {
		return errors.AssignError
	}
	return srv.UpdateByID(ctx)
}

// todo 判断是否存在有人使用，存在的情况下，不允许删除。
func DeleteResource(ctx *gin.Context, in *types.DeleteResourceRequest) error {
	srv := model.Resource{}
	return srv.DeleteByID(ctx, in.ID)
}

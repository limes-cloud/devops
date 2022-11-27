package service

import (
	"configure/errors"
	"configure/meta"
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

func AddResourceService(ctx *gin.Context, in *types.AddResourceServiceRequest) error {
	srv := model.ServiceResource{ResourceID: in.ResourceID}
	var list []model.ServiceResource

	user := meta.User(ctx)
	for _, keyword := range in.ServiceKeywords {
		list = append(list, model.ServiceResource{
			ServiceKeyword: keyword,
			ResourceID:     in.ResourceID,
			Operator:       user.UserName,
			OperatorID:     user.UserId,
		})
	}

	return srv.CreateAll(ctx, list)
}

func AllResourceService(ctx *gin.Context, in *types.AllResourceServiceRequest) ([]model.ServiceResource, error) {
	srv := model.ServiceResource{}
	return srv.All(ctx, in)

}

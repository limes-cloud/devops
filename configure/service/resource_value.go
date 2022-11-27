package service

import (
	"configure/meta"
	"configure/model"
	"configure/types"
	"github.com/limeschool/gin"
)

func AddResourceValue(ctx *gin.Context, in *types.AddResourceValueRequest) error {
	srv := model.ResourceValue{ResourceId: in.ResourceID}
	var list []model.ResourceValue

	user := meta.User(ctx)
	for _, item := range in.Data {
		list = append(list, model.ResourceValue{
			EnvKeyword: item.EnvKeyword,
			ResourceId: in.ResourceID,
			Value:      item.Value,
			Operator:   user.UserName,
			OperatorId: user.UserId,
		})
	}
	return srv.CreateAll(ctx, list)
}

func AllResourceValue(ctx *gin.Context, in *types.AllResourceValueRequest) ([]model.ResourceValue, error) {
	srv := model.ResourceValue{}
	return srv.All(ctx, "resource_id = ?", in.ResourceID)
}

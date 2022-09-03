package service

import (
	"configure/meta"
	"configure/model"
	"configure/types"
	"github.com/limeschool/gin"
)

func AddSystemFieldValue(ctx *gin.Context, in *types.AddSystemFieldValueRequest) error {
	srv := model.SystemFieldValue{FieldId: in.FieldId}
	var list []model.SystemFieldValue

	user := meta.User(ctx)
	for _, item := range in.Data {
		list = append(list, model.SystemFieldValue{
			EnvId:      item.EnvId,
			FieldId:    in.FieldId,
			Value:      item.Value,
			Operator:   user.UserName,
			OperatorId: user.UserId,
		})
	}
	return srv.CreateAll(ctx, list)
}

func AllSystemFieldValue(ctx *gin.Context, in *types.AllSystemFieldValueRequest) ([]model.SystemFieldValue, error) {
	srv := model.SystemFieldValue{}
	return srv.All(ctx, "field_id = ?", in.FieldId)
}

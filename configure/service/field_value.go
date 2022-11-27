package service

import (
	"configure/meta"
	"configure/model"
	"configure/types"
	"github.com/limeschool/gin"
)

func AddFieldValue(ctx *gin.Context, in *types.AddFieldValueRequest) error {
	srv := model.FieldValue{FieldId: in.FieldId}
	var list []model.FieldValue

	user := meta.User(ctx)
	for _, item := range in.Data {
		list = append(list, model.FieldValue{
			EnvKeyword: item.EnvKeyword,
			FieldId:    in.FieldId,
			Value:      item.Value,
			Operator:   user.UserName,
			OperatorId: user.UserId,
		})
	}
	return srv.CreateAll(ctx, list)
}

func AllFieldValue(ctx *gin.Context, in *types.AllFieldValueRequest) ([]model.FieldValue, error) {
	srv := model.FieldValue{}
	return srv.All(ctx, "field_id = ?", in.FieldId)
}

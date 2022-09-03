package service

import (
	"configure/meta"
	"configure/model"
	"configure/types"
	"github.com/limeschool/gin"
)

func AddServiceFieldValue(ctx *gin.Context, in *types.AddServiceFieldValueRequest) error {
	srv := model.ServiceFieldValue{FieldId: in.FieldId}
	var list []model.ServiceFieldValue

	user := meta.User(ctx)
	for _, item := range in.Data {
		list = append(list, model.ServiceFieldValue{
			EnvId:      item.EnvId,
			FieldId:    in.FieldId,
			Value:      item.Value,
			Operator:   user.UserName,
			OperatorId: user.UserId,
		})
	}
	return srv.CreateAll(ctx, list)
}

func AllServiceFieldValue(ctx *gin.Context, in *types.AllServiceFieldValueRequest) ([]model.ServiceFieldValue, error) {
	srv := model.ServiceFieldValue{}
	return srv.All(ctx, "field_id = ?", in.FieldId)
}

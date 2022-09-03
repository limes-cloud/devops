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

func AllServiceEnv(ctx *gin.Context, in *types.AllServiceEnvRequest) ([]model.Environment, error) {
	env := model.Environment{}
	return env.AllFilter(ctx, "id in (select env_id from env_service where srv_id = ?)", in.SrvId)
}

func AllServiceField(ctx *gin.Context, in *types.AllServiceFieldRequest) (interface{}, error) {
	srvField := model.ServiceField{}
	srvFields, _ := srvField.All(ctx, "service_id = ?", in.SrvId)
	if srvFields == nil {
		srvFields = []model.ServiceField{}
	}
	sysFiled := model.SystemField{}
	sysFields, _ := sysFiled.All(ctx, "id in (select system_field_id from service_system_field where service_id = ?)", in.SrvId)
	if sysFields == nil {
		sysFields = []model.SystemField{}
	}
	return gin.H{
		"service":  srvFields,
		"resource": sysFields,
	}, nil
}

func AllService(ctx *gin.Context, in *types.AllServiceRequest) ([]model.Service, error) {
	srv := model.Service{}
	list, err := srv.All(ctx, in)
	if err != nil {
		return nil, err
	}
	for key, item := range list {
		envSrv := model.EnvService{}
		envSrvList, _ := envSrv.All(ctx, "srv_id = ?", item.BaseModel.ID)
		var ids []int64
		for _, ite := range envSrvList {
			ids = append(ids, ite.EnvId)
		}
		list[key].EnvIds = ids
	}
	return list, nil
}

func AddService(ctx *gin.Context, in *types.AddServiceRequest) error {
	srv := model.Service{}
	if copier.Copy(&srv, in) != nil {
		return errors.AssignError
	}
	if tools.DataDup(srv.Create(ctx)) {
		return errors.DulSrvKeywordError
	}
	return AddServiceEnv(ctx, srv.BaseModel.ID, in.EnvIds)
}

func AddServiceEnv(ctx *gin.Context, srvId int64, envIds []int64) error {
	srv := model.EnvService{}
	_ = srv.Delete(ctx, "srv_id = ?", srvId)
	var list []model.EnvService
	user := meta.User(ctx)
	for _, envId := range envIds {
		list = append(list, model.EnvService{
			EnvId:      envId,
			SrvId:      srvId,
			OperatorId: user.UserId,
			Operator:   user.UserName,
		})
	}
	return srv.CreateAll(ctx, list)
}

func UpdateService(ctx *gin.Context, in *types.UpdateServiceRequest) error {
	srv := model.Service{}
	if copier.Copy(&srv, in) != nil {
		return errors.AssignError
	}
	if tools.DataDup(srv.UpdateByID(ctx)) {
		return errors.DulSrvKeywordError
	}
	return AddServiceEnv(ctx, srv.BaseModel.ID, in.EnvIds)
}

func DeleteService(ctx *gin.Context, in *types.DeleteServiceRequest) error {
	srv := model.Service{}
	return srv.DeleteByID(ctx, in.ID)
}

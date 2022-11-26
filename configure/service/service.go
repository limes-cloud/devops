package service

import (
	"configure/errors"
	"configure/meta"
	"configure/model"
	"configure/types"
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
	"gorm.io/gorm"
)

func AddServiceResource(ctx *gin.Context, in *types.AddServiceResourceRequest) error {
	srv := model.ServiceResource{ServiceId: in.ServiceId}
	var list []model.ServiceResource

	user := meta.User(ctx)
	for _, systemId := range in.FieldIds {
		list = append(list, model.ServiceResource{
			ServiceId:  in.ServiceId,
			ResourceId: systemId,
			Operator:   user.UserName,
			OperatorId: user.UserId,
		})
	}

	return srv.CreateAll(ctx, list)
}

func AllServiceResource(ctx *gin.Context, in *types.AllServiceResourceRequest) ([]model.Resource, error) {
	srv := model.ServiceResource{}
	list, err := srv.All(ctx, in)
	if err != nil {
		return nil, err
	}
	var fieldIds []int64
	for _, item := range list {
		fieldIds = append(fieldIds, item.ResourceId)
	}

	field := model.Resource{}
	return field.AllByCallback(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("id in ?", fieldIds)
	})
}

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
	sysFiled := model.Resource{}
	sysFields, _ := sysFiled.AllByCallback(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (select system_field_id from service_system_field where service_id = ?)", in.SrvId)
	})

	if sysFields == nil {
		sysFields = []model.Resource{}
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
		envSrv := model.ServiceEnv{}
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
	if err := srv.Create(ctx); err != nil {
		return err
	}
	return AddServiceEnv(ctx, srv.BaseModel.ID, in.EnvIds)
}

func AddServiceEnv(ctx *gin.Context, srvId int64, envIds []int64) error {
	srv := model.ServiceEnv{}
	_ = srv.Delete(ctx, "srv_id = ?", srvId)
	var list []model.ServiceEnv
	user := meta.User(ctx)
	for _, envId := range envIds {
		list = append(list, model.ServiceEnv{
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
	if err := srv.UpdateByID(ctx); err != nil {
		return err
	}
	return AddServiceEnv(ctx, srv.BaseModel.ID, in.EnvIds)
}

func DeleteService(ctx *gin.Context, in *types.DeleteServiceRequest) error {
	srv := model.Service{}
	return srv.DeleteByID(ctx, in.ID)
}

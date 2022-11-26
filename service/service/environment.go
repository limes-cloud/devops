package service

import (
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
	"service/errors"
	"service/meta"
	"service/model"
	"service/types"
)

func AllEnvironment(ctx *gin.Context, in *types.AllEnvironmentRequest) ([]model.Environment, error) {
	env := model.Environment{}
	return env.All(ctx, in)
}

func AllEnvironmentFilter(ctx *gin.Context) ([]model.Environment, error) {
	env := model.Environment{}
	return env.AllFilter(ctx)
}

func AddEnvironment(ctx *gin.Context, in *types.AddEnvironmentRequest) error {
	env := model.Environment{}
	if copier.Copy(&env, in) != nil {
		return errors.AssignError
	}
	return env.Create(ctx)
}

func UpdateEnvironment(ctx *gin.Context, in *types.UpdateEnvironmentRequest) error {
	env := model.Environment{}
	if copier.Copy(&env, in) != nil {
		return errors.AssignError
	}
	return env.UpdateByID(ctx)
}

func DeleteEnvironment(ctx *gin.Context, in *types.DeleteEnvironmentRequest) error {
	env := model.Environment{}
	return env.DeleteByID(ctx, in.ID)
}

func UpdateServiceEnv(ctx *gin.Context, in *types.UpdateServiceEnvRequest) error {
	m := model.ServiceEnv{}
	_ = m.Delete(ctx, "env_id = ?", in.ID)
	user := meta.User(ctx)
	var list []model.ServiceEnv
	for _, srvId := range in.SrvIds {
		list = append(list, model.ServiceEnv{
			EnvId:      in.ID,
			SrvId:      srvId,
			Operator:   user.UserName,
			OperatorId: user.UserId,
		})
	}
	return m.CreateAll(ctx, list)
}

//func AllServiceEnv(ctx *gin.Context, in *types.AllServiceEnvRequest) ([]model.ServiceEnv, error) {
//	m := model.ServiceEnv{}
//	if in.EnvId != 0 {
//		return m.All(ctx, "env_id = ?", in.EnvId)
//	} else {
//		return m.All(ctx, "srv_id = ?", in.EnvId)
//	}
//}

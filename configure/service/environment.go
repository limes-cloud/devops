package service

import (
	"configure/errors"
	"configure/meta"
	"configure/model"
	"configure/tools"
	"configure/types"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
	"github.com/limeschool/gin/config_drive"
)

func AllEnvironment(ctx *gin.Context, in *types.AllEnvironmentRequest) ([]model.Environment, error) {
	env := model.Environment{}
	return env.All(ctx, in)
}

func AllEnvironmentFilter(ctx *gin.Context) ([]model.Environment, error) {
	env := model.Environment{}
	return env.AllFilter(ctx)
}

func EnvironmentConnect(ctx *gin.Context, in *types.EnvironmentConnectRequest) (config_drive.ConfigService, error) {
	env := model.Environment{}
	if env.OneById(ctx, in.ID) != nil {
		return nil, errors.New("不存在此环境数据")
	}
	config := &config_drive.Config{
		Drive: env.Drive,
		Path:  env.Prefix,
	}
	_ = json.Unmarshal([]byte(env.Config), &config)
	var err error
	var cs config_drive.ConfigService
	switch config.Drive {
	case "etcd":
		cs, err = config_drive.NewEtcd(config)
	case "zk":
		cs, err = config_drive.NewZK(config)
	case "consul":
		cs, err = config_drive.NewConsul(config)
	default:
		err = errors.New("不支持的中间件配置")
	}
	if err != nil {
		return nil, errors.New("连接失败：" + err.Error())
	}
	return cs, nil
}

func AddEnvironment(ctx *gin.Context, in *types.AddEnvironmentRequest) error {
	env := model.Environment{}
	if copier.Copy(&env, in) != nil {
		return errors.AssignError
	}
	env.Token = tools.GenToken()
	if tools.DataDup(env.Create(ctx)) {
		return errors.DulEnvKeywordError
	}
	return nil
}

func UpdateEnvironment(ctx *gin.Context, in *types.UpdateEnvironmentRequest) error {
	env := model.Environment{}
	if copier.Copy(&env, in) != nil {
		return errors.AssignError
	}
	err := env.UpdateByID(ctx)
	if tools.DataDup(err) {
		return errors.DulEnvKeywordError
	}
	return err
}

func DeleteEnvironment(ctx *gin.Context, in *types.DeleteEnvironmentRequest) error {
	env := model.Environment{}
	return env.DeleteByID(ctx, in.ID)
}

func UpdateEnvService(ctx *gin.Context, in *types.UpdateEnvServiceRequest) error {
	m := model.EnvService{}
	_ = m.Delete(ctx, "env_id = ?", in.ID)
	user := meta.User(ctx)
	var list []model.EnvService
	for _, srvId := range in.SrvIds {
		list = append(list, model.EnvService{
			EnvId:      in.ID,
			SrvId:      srvId,
			Operator:   user.UserName,
			OperatorId: user.UserId,
		})
	}
	return m.CreateAll(ctx, list)
}

func AllEnvService(ctx *gin.Context, in *types.AllEnvServiceRequest) ([]model.EnvService, error) {
	m := model.EnvService{}
	if in.EnvId != 0 {
		return m.All(ctx, "env_id = ?", in.EnvId)
	} else {
		return m.All(ctx, "srv_id = ?", in.EnvId)
	}
}

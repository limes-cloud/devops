package service

import (
	"configure/errors"
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

func EnvironmentConnect(ctx *gin.Context, in *types.EnvironmentConnectRequest) (config_drive.ConfigService, error) {
	env := model.Environment{}
	if env.One(ctx, "env_keyword=?", in.Keyword) != nil {
		return nil, errors.New("不存在此环境中间件配置")
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

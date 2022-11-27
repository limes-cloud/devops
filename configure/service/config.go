package service

import (
	"configure/errors"
	"configure/model"
	"configure/tools"
	"configure/types"
	"encoding/json"
	"fmt"
	"github.com/limeschool/gin"
	"github.com/limeschool/gin/config_drive"
)

func RollbackConfig(ctx *gin.Context, in *types.RollbackConfigRequest) error {
	log := model.TemplateLog{}
	if err := log.OneById(ctx, in.ID); err != nil {
		return err
	}

	// 连接配置中心
	client, err := EnvironmentConnect(ctx, &types.EnvironmentConnectRequest{Keyword: log.EnvKeyword})
	if err != nil {
		return err
	}
	return client.Set(log.Config)
}

func DriverConfig(ctx *gin.Context, in *types.DriverConfigRequest) (string, error) {
	// 连接配置中心
	client, err := EnvironmentConnect(ctx, &types.EnvironmentConnectRequest{Keyword: in.EnvKeyword})
	if err != nil {
		return "", err
	}
	client.SetPath("/" + in.SrvKeyword)
	data, _ := client.Get()
	return string(data), nil
}

func Config(ctx *gin.Context, in *types.ConfigRequest) (*config_drive.Config, error) {
	env := model.Environment{}
	if env.One(ctx, "token = ? ", in.Token) != nil {
		return nil, errors.New("env value error")
	}
	config := &config_drive.Config{
		Drive: env.Drive,
		Path:  env.Prefix + "/" + in.Service,
		Type:  "json",
	}
	return config, json.Unmarshal([]byte(env.Config), &config)
}

func AllConfigLog(ctx *gin.Context, in *types.AllConfigLogRequest) ([]model.TemplateLog, error) {
	log := model.TemplateLog{}
	return log.All(ctx, in)
}

func ConfigLog(ctx *gin.Context, in *types.ConfigLogRequest) (model.TemplateLog, error) {
	log := model.TemplateLog{}
	return log, log.OneById(ctx, in.ID)
}

func CompareConfig(ctx *gin.Context, in *types.CompareConfigRequest) ([]interface{}, error) {
	config, err := ParseTemplate(ctx, &types.ParseTemplateRequest{SrvKeyword: in.SrvKeyword, EnvKeyword: in.EnvKeyword})
	if err != nil {
		return nil, err
	}
	newConf := map[string]interface{}{}
	_ = json.Unmarshal([]byte(config), &newConf)
	var newKeys []string
	for key, _ := range newConf {
		newKeys = append(newKeys, key)
	}

	// 连接配置中心
	client, err := EnvironmentConnect(ctx, &types.EnvironmentConnectRequest{Keyword: in.EnvKeyword})
	if err != nil {
		return nil, err
	}
	client.SetPath("/" + in.SrvKeyword)
	data, _ := client.Get()
	oldConf := map[string]interface{}{}
	_ = json.Unmarshal(data, &oldConf)
	var oldKeys []string
	for key, _ := range oldConf {
		oldKeys = append(oldKeys, key)
	}
	// 进行对比
	var resp []interface{}
	adds := tools.Diff(newKeys, oldKeys)
	for _, val := range adds {
		resp = append(resp, gin.H{"type": "add", "key": val, "old": nil, "val": newConf[val]})
	}
	dels := tools.Diff(oldKeys, newKeys)
	for _, val := range dels {
		resp = append(resp, gin.H{"type": "delete", "key": val, "old": newConf[val], "val": nil})
	}
	for key, val := range oldConf {
		if v := newConf[key]; v != nil {
			if fmt.Sprint(val) != fmt.Sprint(v) {
				resp = append(resp, gin.H{"type": "update", "key": val, "old": val, "val": v})
			}
		}
	}
	return resp, nil
}

func SyncConfig(ctx *gin.Context, in *types.SyncConfigRequest) error {
	// 解析配置
	config, err := ParseTemplate(ctx, &types.ParseTemplateRequest{SrvKeyword: in.SrvKeyword, EnvKeyword: in.EnvKeyword})
	if err != nil {
		return err
	}

	// 对比
	resp, err := CompareConfig(ctx, &types.CompareConfigRequest{SrvKeyword: in.SrvKeyword, EnvKeyword: in.EnvKeyword})
	if err != nil {
		return err
	}
	respStr, _ := json.Marshal(resp)

	// 连接配置中心
	client, err := EnvironmentConnect(ctx, &types.EnvironmentConnectRequest{Keyword: in.EnvKeyword})
	if err != nil {
		return err
	}
	client.SetPath("/" + in.SrvKeyword)
	if err = client.Set(config); err != nil {
		return err
	}
	log := model.TemplateLog{
		ServiceKeyword: in.SrvKeyword,
		EnvKeyword:     in.EnvKeyword,
		Config:         config,
		Description:    string(respStr),
	}
	return log.Create(ctx)
}

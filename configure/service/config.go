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
	"gorm.io/gorm"
	"regexp"
	"strings"
)

func RollbackConfig(ctx *gin.Context, in *types.RollbackConfigRequest) error {
	log := model.TemplateLog{}
	if err := log.OneById(ctx, in.ID); err != nil {
		return err
	}

	// 获取服务信息
	server := model.Service{}
	if server.OneById(ctx, log.SrvId) != nil {
		return errors.New("未找到对应的服务信息")
	}
	// 连接配置中心
	client, err := EnvironmentConnect(ctx, &types.EnvironmentConnectRequest{ID: log.EnvId})
	if err != nil {
		return err
	}
	return client.Set(log.Config)
}

func DriverConfig(ctx *gin.Context, in *types.DriverConfigRequest) (string, error) {
	// 获取服务信息
	server := model.Service{}
	if server.OneById(ctx, in.SrvId) != nil {
		return "", errors.New("未找到对应的服务信息")
	}
	// 连接配置中心
	client, err := EnvironmentConnect(ctx, &types.EnvironmentConnectRequest{ID: in.EnvId})
	if err != nil {
		return "", err
	}
	client.SetPath("/" + server.Keyword)
	data, _ := client.Get()
	return string(data), nil
}

func Config(ctx *gin.Context, in *types.ConfigRequest) (*config_drive.Config, error) {
	// 获取服务信息
	server := model.Service{}
	if server.OneByKeyword(ctx, in.Service) != nil {
		return nil, errors.New("token value error")
	}
	// 获取环境信息
	env := model.Environment{}
	if env.One(ctx, "token = ? ", in.Token) != nil {
		return nil, errors.New("env value error")
	}
	config := &config_drive.Config{
		Drive: env.Drive,
		Path:  env.Prefix + "/" + server.Keyword,
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
	config, err := ParseTemplate(ctx, &types.ParseTemplateRequest{SrvId: in.SrvId, EnvId: in.EnvId})
	if err != nil {
		return nil, err
	}
	newConf := map[string]interface{}{}
	_ = json.Unmarshal([]byte(config), &newConf)
	var newKeys []string
	for key, _ := range newConf {
		newKeys = append(newKeys, key)
	}
	// 获取服务信息
	server := model.Service{}
	if server.OneById(ctx, in.SrvId) != nil {
		return nil, errors.New("未找到对应的服务信息")
	}
	// 连接配置中心
	client, err := EnvironmentConnect(ctx, &types.EnvironmentConnectRequest{ID: in.EnvId})
	if err != nil {
		return nil, err
	}
	client.SetPath("/" + server.Keyword)
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
	config, err := ParseTemplate(ctx, &types.ParseTemplateRequest{SrvId: in.SrvId, EnvId: in.EnvId})
	if err != nil {
		return err
	}
	resp, err := CompareConfig(ctx, &types.CompareConfigRequest{SrvId: in.SrvId, EnvId: in.EnvId})
	if err != nil {
		return err
	}
	respStr, _ := json.Marshal(resp)
	// 获取服务信息
	server := model.Service{}
	if server.OneById(ctx, in.SrvId) != nil {
		return errors.New("未找到对应的服务信息")
	}
	// 连接配置中心
	client, err := EnvironmentConnect(ctx, &types.EnvironmentConnectRequest{ID: in.EnvId})
	if err != nil {
		return err
	}
	client.SetPath("/" + server.Keyword)
	if err = client.Set(config); err != nil {
		return err
	}
	log := model.TemplateLog{
		SrvId:       in.SrvId,
		EnvId:       in.EnvId,
		Config:      config,
		Description: string(respStr),
	}
	return log.Create(ctx)
}

func ParseTemplate(ctx *gin.Context, in *types.ParseTemplateRequest) (string, error) {
	// 当前的配置模板
	template, err := GetTemplate(ctx, &types.GetTemplateRequest{SrvId: in.SrvId})
	if err != nil {
		return "", err
	}

	//获取服务字段以及值配置
	sf := model.ServiceField{}
	srvSf, _ := sf.AllByCallback(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Select("field,sv.value field_value").
			Joins("left join service_field_value sv on service_field.id = sv.field_id and sv.env_id = ?", in.EnvId).
			Where("service_field.service_id = ?", in.SrvId)
	})

	rf := model.Resource{}
	sysFs, _ := rf.AllByCallback(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Select("field,sv.value field_value").
			Joins("left join system_field_value sv on system_field.id = sv.field_id and sv.env_id = ?", in.EnvId).
			Where("system_field.id in (select system_field_id from service_system_field where service_id = ?)", in.SrvId)
	})

	//组合两边的key
	keys := map[string]value{}
	for _, item := range srvSf {
		keys[fmt.Sprintf("{{%v}}", item.Field)] = parseValue(item.FieldValue)
	}
	for _, item := range sysFs {
		config := map[string]string{}
		_ = json.Unmarshal([]byte(item.FieldValue), &config)
		for key, val := range config {
			k := fmt.Sprintf("%v.%v", item.Field, key)
			keys[fmt.Sprintf("{{%v}}", k)] = parseValue(val)
		}
	}

	//进行增则匹配
	reg := regexp.MustCompile(`\{\{(\w|\.)+}}`)
	tempKeys := reg.FindAllString(template.Content, -1)
	// 进行参数判断
	for _, key := range tempKeys {
		if val, ok := keys[key]; !ok {
			return "", fmt.Errorf("非法字段：%v", key)
		} else {
			template.Content = replace(template.Content, key, val)
		}
	}
	return template.Content, nil
}

type value struct {
	value   interface{}
	exclude bool
}

func replace(template, key string, val value) string {
	if strings.Contains(template, fmt.Sprintf(`"%v"`, key)) && val.exclude {
		template = strings.Replace(template, fmt.Sprintf(`"%v"`, key), fmt.Sprintf("%v", val.value), 1)
	} else {
		template = strings.Replace(template, key, fmt.Sprintf("%v", val.value), 1)
	}
	return template
}

func parseValue(v interface{}) value {
	switch v.(type) {
	case string:
		// 判断能否转成json
		var data interface{}
		if json.Unmarshal([]byte(v.(string)), &data) != nil && data != nil {
			return parseValue(data)
		}
		return value{value: v, exclude: false}
	case map[string]interface{}, []interface{}:
		b, _ := json.Marshal(v)
		return value{value: string(b), exclude: true}
	default:
		return value{value: fmt.Sprintf("%v", v), exclude: true}
	}
}

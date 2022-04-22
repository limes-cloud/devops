package common

import (
	"devops/configrue/models"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

type value struct {
	value   interface{}
	exclude bool
}

func ParseTemplate(serviceId int64, env, template string) (string, error) {
	//获取服务关键字
	service := models.Service{}
	service.ID = serviceId
	if service.OneByID() != nil {
		return "", errors.New("未找到对应的服务信息")
	}

	//获取指定服务的全部字段
	sf := models.ServiceField{}
	serviceFields, _, _ := sf.All(nil, func(db *gorm.DB) *gorm.DB {
		return db.Select("field,config").Where("service_id = ?", serviceId)
	})
	//获取全部公共服务字段
	//这里应该要获取到config 判断配置里面的字段是痘存在
	cf := models.ConfigureField{}
	configureFields, _, _ := cf.All(nil, func(db *gorm.DB) *gorm.DB {
		return db.Select("field,config")
	})

	//组合两边的key
	serviceConfig := map[string]interface{}{}
	keys := map[string]value{}
	for _, item := range serviceFields {
		if json.Unmarshal([]byte(item.Config), &serviceConfig) != nil {
			return "", errors.New(item.Field + "字段数据格式错误")
		}
		keys[fmt.Sprintf("{{%v}}", item.Field)] = parseValue(serviceConfig[env])
	}
	config := map[string]string{}
	for _, item := range configureFields {
		if json.Unmarshal([]byte(item.Config), &config) != nil {
			return "", errors.New(item.Field + "字段数据格式错误")
		}
		for _, val := range config {
			temp := map[string]interface{}{}
			if json.Unmarshal([]byte(val), &temp) != nil {
				return "", errors.New("全局字段" + item.Field + "配置出错")
			}
			for k, v := range temp {
				keys[fmt.Sprintf("{{%v}}", item.Field+"."+k)] = parseValue(v)
			}
		}
	}
	//进行增则匹配
	reg := regexp.MustCompile(`\{\{(\w|\.)+}}`)
	tempKeys := reg.FindAllString(template, -1)
	// 进行参数判断
	for _, key := range tempKeys {
		if val, ok := keys[key]; !ok {
			return "", fmt.Errorf("非法字段：%v", key)
		} else {
			template = replace(template, key, val)

		}
	}
	return template, nil
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
	if v == "" {
		return value{
			value:   "",
			exclude: false,
		}
	}
	var data interface{}
	if str, ok := v.(string); ok {
		if json.Unmarshal([]byte(str), &data) != nil {
			return value{
				value:   str,
				exclude: true,
			}
		}
	} else {
		data = v
	}

	switch data.(type) {
	case string:
		return value{
			value:   data,
			exclude: false,
		}
	case map[string]interface{}, []interface{}:
		b, _ := json.Marshal(data)
		return value{
			value:   string(b),
			exclude: true,
		}
	default:
		return value{
			value:   fmt.Sprintf("%v", data),
			exclude: true,
		}
	}
}

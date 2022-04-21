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
	config := map[string]interface{}{}
	keys := map[string]interface{}{}
	for _, item := range serviceFields {
		if json.Unmarshal([]byte(item.Config), &config) != nil {
			return "", errors.New(item.Field + "字段数据格式错误")
		}
		keys[fmt.Sprintf("{{%v}}", item.Field)] = config[env]
	}

	for _, item := range configureFields {
		if json.Unmarshal([]byte(item.Config), &config) != nil {
			return "", errors.New(item.Field + "字段数据格式错误")
		}
		for key, val := range config {
			keys[fmt.Sprintf("{{%v}}", item.Field+"."+key)] = val
		}
	}
	//进行增则匹配
	reg := regexp.MustCompile(`\{\{(\w|\.)+}}`)
	tempKeys := reg.FindAllString(template, -1)
	// 进行参数判断
	for _, key := range tempKeys {
		if conf, ok := keys[key]; !ok {
			return "", fmt.Errorf("非法字段：%v", key)
		} else {
			template = strings.Replace(template, key, fmt.Sprintf("%v", conf), 1)
		}
	}
	return template, nil
}

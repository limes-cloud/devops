package service

import (
	"configure/model"
	"configure/types"
	"encoding/json"
	"fmt"
	"github.com/limeschool/gin"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

func CheckTemplate(ctx *gin.Context, keyword string, template string) error {

	//获取服务关键字
	sf := model.Field{}
	srvFs, _ := sf.All(ctx, "service_keyword = ?", keyword)

	// 获取资源字段
	rf := model.Resource{}
	sysFs, _ := rf.AllByCallback(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (select id from service_resource where service_keyword=?)", keyword)
	})

	//组合两边的key
	keys := map[string]bool{}
	for _, item := range sysFs {
		var fields []string
		_ = json.Unmarshal([]byte(item.ChildField), &fields)
		for _, val := range fields {
			keys[filedKey(item.Field+"."+val)] = true
		}
	}

	for _, item := range srvFs {
		keys[filedKey(item.Field)] = true
	}

	//进行增则匹配
	reg := regexp.MustCompile(`\{\{(\w|\.)+}}`)
	tempKeys := reg.FindAllString(template, -1)
	// 进行参数判断
	for _, key := range tempKeys {
		if !keys[key] {
			return fmt.Errorf("非法字段：%v", key)
		}
	}
	return nil
}

func filedKey(val string) string {
	return fmt.Sprintf("{{%v}}", val)
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

func ParseTemplate(ctx *gin.Context, in *types.ParseTemplateRequest) (string, error) {
	// 当前的配置模板
	template, err := GetTemplate(ctx, &types.GetTemplateRequest{Keyword: in.SrvKeyword})
	if err != nil {
		return "", err
	}

	//获取服务字段以及值配置
	sf := model.Field{}
	srvSf, _ := sf.AllByCallback(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Select("field,sv.value field_value").
			Joins("left join field_value sv on field.id = sv.field_id and sv.env_keyword = ?", in.EnvKeyword).
			Where("field.service_keyword = ?", in.SrvKeyword)
	})

	rf := model.Resource{}
	sysFs, _ := rf.AllByCallback(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Select("field,sv.value field_value").
			Joins("left join resource_value sv on resource.id = sv.resource_id and sv.env_keyword = ?", in.EnvKeyword).
			Where("resource.id in (select resource from service_resource where service_keyword = ?)", in.SrvKeyword)
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

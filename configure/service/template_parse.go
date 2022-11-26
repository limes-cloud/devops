package service

import (
	"configure/model"
	"encoding/json"
	"fmt"
	"github.com/limeschool/gin"
	"gorm.io/gorm"
	"regexp"
)

func CheckTemplate(ctx *gin.Context, serviceId int64, template string) error {
	//获取服务关键字
	sf := model.ServiceField{}
	//获取指定服务的全部字段
	srvFs, _ := sf.All(ctx, "service_id = ?", serviceId)

	rf := model.Resource{}
	sysFs, _ := rf.AllByCallback(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (select id from service_system_field where service_id = ?)", serviceId)
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

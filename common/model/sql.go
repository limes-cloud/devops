package model

import (
	"devops/common/tools"
	"devops/common/tools/jsonx"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

func SqlWhere(db *gorm.DB, query interface{}, exclude ...string) *gorm.DB {
	m := make(map[string]interface{})
	switch query.(type) {
	case map[string]string, map[string]interface{}, map[string]int64, map[string]int:
		tools.Transform(query, &m)
	default:
		m = toMap(db, query)
	}
	for key, item := range m {
		key = jsonx.Camel2Case(key)
		if inListStr(exclude, key) {
			continue
		}
		db = db.Where(key, item)
	}
	return db
}

func inListStr(list []string, val string) bool {
	val = val[strings.Index(val, ".")+1:]
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

func toMap(db *gorm.DB, src interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	value := reflect.ValueOf(src)
	tp := reflect.TypeOf(src)
	for value.Kind().String() == "ptr" {
		value = value.Elem()
		tp = tp.Elem()
	}

	if value.Kind().String() == "struct" {
		num := value.NumField()
		for i := 0; i < num; i++ {
			jsonTag := tp.Field(i).Tag.Get("json")
			if strings.Contains(jsonTag, "optional") || strings.Contains(jsonTag, "omitempty") {
				if isBlank(value.Field(i)) {
					continue
				}
			}

			formTag := tp.Field(i).Tag.Get("form")
			if strings.Contains(formTag, "optional") {
				if isBlank(value.Field(i)) {
					continue
				}
			}
			m[db.Statement.Table+"."+tp.Field(i).Name] = value.Field(i).Interface()
		}
	}
	return m
}

func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

package tools

import (
	"encoding/json"
)

func Transform(src, dst interface{}) {
	b, _ := json.Marshal(src)
	_ = json.Unmarshal(b, dst)
}

func ToMap(src interface{}) map[string]interface{} {
	var m = make(map[string]interface{})
	Transform(src, &m)
	return m
}

func InListStr(list []string, val string) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

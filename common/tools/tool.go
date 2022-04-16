package tools

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func CheckIP(ip string) bool {
	addr := strings.Trim(ip, " ")
	regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	if match, _ := regexp.MatchString(regStr, addr); match {
		return true
	}
	return false
}

func Transform(src, dst interface{}) {
	b, _ := json.Marshal(src)
	_ = json.Unmarshal(b, dst)
}

func ToMap(m interface{}) map[string]interface{} {
	var dest = make(map[string]interface{})
	switch m.(type) {
	case map[string]interface{}:
		dest = m.(map[string]interface{})
	default:
		Transform(m, &dest)
	}

	Transform(m, &dest)
	return dest
}

func InListStr(list []string, val string) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

func Exclude(m interface{}, keys ...string) map[string]interface{} {
	dest := make(map[string]interface{})
	switch m.(type) {
	case map[string]interface{}:
		dest = m.(map[string]interface{})
	default:
		Transform(m, &dest)
	}

	for _, key := range keys {
		delete(dest, key)
	}
	return dest
}

// todo 这里有一个小bug 就是exlud 过来的map false
func ToHasValueMap(m interface{}, keys ...string) map[string]interface{} {
	dest := make(map[string]interface{})
	switch m.(type) {
	case map[string]interface{}:
		dest = m.(map[string]interface{})
	default:
		Transform(m, &dest)
	}
	for key, item := range dest {
		if !HasValue(item) {
			delete(dest, key)
		}
	}
	return dest
}

func HasValue(val interface{}) bool {
	switch val.(type) {
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64:
		return fmt.Sprintf("%v", val) != "0"
	case bool:
		return val.(bool)
	case string:
		return val.(string) != ""
	default:
		return val != nil
	}
}

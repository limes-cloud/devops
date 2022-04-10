package jsonx

import (
	"bytes"
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"unicode"
)

type JsonSnakeCase struct {
	Value interface{}
}

func (c JsonSnakeCase) MarshalJSON() ([]byte, error) {
	// Regexp definitions
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)
	marshalled, err := json.Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			return bytes.ToLower(wordBarrierRegex.ReplaceAll(
				match,
				[]byte(`${1}_${2}`),
			))
		},
	)
	return converted, err
}

type JsonCamelCase struct {
	Value interface{}
}

//
//func (c JsonCamelCase) MarshalJSON() ([]byte, error) {
//	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
//	marshalled, err := json.Marshal(c.Value)
//	converted := keyMatchRegex.ReplaceAllFunc(
//		marshalled,
//		func(match []byte) []byte {
//			matchStr := string(match)
//			key := matchStr[1 : len(matchStr)-2]
//			resKey := Lcfirst(Case2Camel(key))
//			return []byte(`"` + resKey + `":`)
//		},
//	)
//	return converted, err
//}

func Camel2Case(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 && !unicode.IsUpper(rune(name[i-1])) && name[i-1] != '.' {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

//
//func Case2Camel(name string) string {
//	name = strings.Replace(name, "_", " ", -1)
//	name = strings.Title(name)
//	return strings.Replace(name, " ", "", -1)
//}

//func Ucfirst(str string) string {
//	for i, v := range str {
//		return string(unicode.ToUpper(v)) + str[i+1:]
//	}
//	return ""
//}

//func Lcfirst(str string) string {
//	for i, v := range str {
//		return string(unicode.ToLower(v)) + str[i+1:]
//	}
//	return ""
//}

type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()
	b.WriteString(s)
	return b
}

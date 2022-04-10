package response

import (
	"github.com/go-sql-driver/mysql"
	"strings"
)

const defaultCode = 400

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

// error 需要实现Error() 这个方法
func (e *CodeError) Error() string {
	return e.Msg
}

// Data 返回自定义的结构体
func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

// NewDefaultError 处理默认返回数据
func NewDefaultError(msg string) *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: defaultCode,
		Msg:  msg,
	}
}

// HandlerError 处理错误参数返回格式
func HandlerError(err error) interface{} {
	switch e := err.(type) {
	case *CodeError:
		return e.Data()
	case *mysql.MySQLError:
		return NewDefaultError(handleMysqlError(e.Error()))
	default:
		return NewDefaultError(err.Error())
	}
}

func handleMysqlError(e string) string {
	if strings.Contains(e, "Duplicate") {
		info := strings.Split(e, "'")
		if len(info) > 1 {
			return "系统已存在数据：" + info[1]
		}
	}
	return e
}

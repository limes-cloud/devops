package errors

import "github.com/limeschool/gin"

var (
	ParamsError = &gin.CustomError{Code: 210001, Msg: "参数验证失败"}
	AssignError = &gin.CustomError{Code: 210002, Msg: "数据赋值失败"}
)

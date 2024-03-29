package errors

import "github.com/limeschool/gin"

var (
	ParamsError     = &gin.CustomError{Code: 110002, Msg: "参数验证失败"}
	AssignError     = &gin.CustomError{Code: 110003, Msg: "数据赋值失败"}
	DBError         = &gin.CustomError{Code: 110004, Msg: "数据库操作失败"}
	DBNotFoundError = &gin.CustomError{Code: 110005, Msg: "未查询到指定数据"}
)

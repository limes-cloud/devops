package errors

import (
	"errors"
	"fmt"
	"github.com/limeschool/gin"
)

const (
	defaultCode = 210000 //config通用错误吗
)

var (
	New = func(msg string) error {
		return &gin.CustomError{
			Code: defaultCode,
			Msg:  msg,
		}
	}

	NewF = func(msg string, arg ...interface{}) error {
		return &gin.CustomError{
			Code: defaultCode,
			Msg:  fmt.Sprintf(msg, arg...),
		}
	}

	Is = func(err, tar error) bool {
		return errors.Is(err, tar)
	}
)

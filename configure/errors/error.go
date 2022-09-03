package errors

import (
	"errors"
	"fmt"
	"github.com/limeschool/gin"
	"gorm.io/gorm"
)

const (
	defaultCode = 110001 //config通用错误吗
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
		return errors.Is(err, gorm.ErrRecordNotFound)
	}
)

package biz

import (
	"errors"
	"fmt"
)

type Error struct {
	Err  error             // 原始错误
	Code string            // 错误码
	Args map[string]string // 错误参数
}

func (Self *Error) Error() string {
	return Self.Err.Error()
}

func WithError(err error) *Error {
	return &Error{
		Err:  err,
		Code: CodeServerError,
	}
}

func NewValidError(message string) *Error {
	return &Error{
		Code: CodeValidError,
		Err:  errors.New(message),
		Args: map[string]string{
			"message": message,
		},
	}
}

func NewParamError(fieldName, fieldValue string) *Error {
	return &Error{
		Code: CodeParamError,
		Err:  fmt.Errorf("参数验证失败.fieldName=%s,fieldValue=%s", fieldName, fieldValue),
		Args: map[string]string{
			"fieldName":  fieldName,
			"fieldValue": fieldValue,
		},
	}
}

//-----------------------------------------------------------------------------------------
//						错误转换处理器
//-----------------------------------------------------------------------------------------

type ErrorHandler func(err error) *Error

var errorHandlers = []ErrorHandler{
	bizErrorHandler(),
}

func RegisterErrorHandler(handler ErrorHandler) {
	errorHandlers = append(errorHandlers, handler)
}

func bizErrorHandler() ErrorHandler {
	return func(err error) *Error {
		var ex *Error
		if errors.As(err, &ex) {
			return ex
		}
		return nil
	}
}

func ConverterError(err error) *Error {
	for index := range errorHandlers {
		if ex := errorHandlers[index](err); ex != nil {
			return ex
		}
	}
	return nil
}

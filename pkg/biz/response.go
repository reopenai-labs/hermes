package biz

import (
	"qiyibyte.com/hermes/pkg/i18n"
)

type ApiResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(data any) *ApiResponse {
	return &ApiResponse{
		Code:    CodeOk,
		Message: "success",
		Data:    data,
	}
}

func Failure(code, message string) *ApiResponse {
	return &ApiResponse{
		Code:    code,
		Message: message,
	}
}

func FailureWithLocale(locale any, code string, args map[string]string) *ApiResponse {
	return &ApiResponse{
		Code:    code,
		Message: i18n.GetMessageWithArgs(locale, code, args),
	}
}

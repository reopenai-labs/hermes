package biz

const (
	CodeOk             = "200"
	CodeNotFound       = "404"
	CodeMethodNotAllow = "405"
	CodeServerError    = "500"
)

const (
	CodeValidError = "4001" // 通用验证失败的错误.会将原始的参数错误返回给客户端
	CodeParamError = "4002" // 参数验证失败的错误.接受参数: fieldName: 字段名, fieldValue: 字段值
)

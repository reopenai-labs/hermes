package hertzx

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"qiyibyte.com/hermes/pkg/biz"
	"qiyibyte.com/hermes/pkg/i18n"
	"qiyibyte.com/hermes/pkg/logger"
)

type errorHandler func(c context.Context, ctx *app.RequestContext, err error)

var errorHandlers = make([]errorHandler, 0)

func init() {
	// 业务异常处理器
	AddErrorHandler(func(c context.Context, ctx *app.RequestContext, err error) {
		var ex *biz.Error
		if errors.As(err, &ex) {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				biz.FailureWithLocale(c, ex.Code, ex.Args),
			)
		}
	})
	// JSON异常处理器
	AddErrorHandler(func(c context.Context, ctx *app.RequestContext, err error) {
		var unmarshalTypeError *json.UnmarshalTypeError
		if errors.As(err, &unmarshalTypeError) {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				biz.Failure(biz.CodeParamError, i18n.GetMessageWithArgs(c, biz.CodeParamError, map[string]string{
					"fieldName":  unmarshalTypeError.Field,
					"fieldValue": unmarshalTypeError.Value,
				})),
			)
		}
	})
}

func globalErrorHandler(c context.Context, ctx *app.RequestContext, err error) {
	for i := range errorHandlers {
		errorHandlers[i](c, ctx, err)
		if ctx.IsAborted() {
			return
		}
	}
	logger.ErrorfWithCtx(c, "[Error]未知错误:%s", err.Error())
	ctx.JSON(
		http.StatusInternalServerError,
		biz.Failure(biz.CodeServerError, i18n.GetMessage(c, biz.CodeServerError)),
	)
}

func AddErrorHandler(handle func(c context.Context, ctx *app.RequestContext, err error)) {
	errorHandlers = append(errorHandlers, handle)
}

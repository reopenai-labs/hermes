package hertzx

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"qiyibyte.com/hermes/pkg/biz"
)

func RequestBaseHandler(handler func(c context.Context, ctx *app.RequestContext) (any, error)) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		response, err := handler(c, ctx)
		if err != nil {
			_ = ctx.Error(err)
		} else {
			ctx.JSON(http.StatusOK, biz.Success(response))
		}
	}
}

func RequestBodyHandler[T any](handler func(c context.Context, ctx *app.RequestContext, request *T) (any, error)) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var request T
		if err := ctx.BindJSON(&request); err != nil {
			_ = ctx.Error(err)
			return
		}
		if err := ctx.Validate(&request); err != nil {
			_ = ctx.Error(err)
			return
		}
		response, err := handler(c, ctx, &request)
		if err != nil {
			_ = ctx.Error(err)
		} else {
			ctx.JSON(http.StatusOK, biz.Success(response))
		}
	}
}

func RequestParamsHandler[T any](handler func(c context.Context, ctx *app.RequestContext, request *T) (any, error)) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var request T
		if err := ctx.BindQuery(&request); err != nil {
			_ = ctx.Error(err)
			return
		}
		if err := ctx.Validate(&request); err != nil {
			_ = ctx.Error(err)
			return
		}
		response, err := handler(c, ctx, &request)
		if err != nil {
			_ = ctx.Error(err)
		} else {
			ctx.JSON(http.StatusOK, biz.Success(response))
		}
	}
}

func RequestFormHandler[T any](handler func(c context.Context, ctx *app.RequestContext, request *T) (any, error)) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var request T
		if err := ctx.BindForm(&request); err != nil {
			_ = ctx.Error(err)
			return
		}
		if err := ctx.Validate(&request); err != nil {
			_ = ctx.Error(err)
			return
		}
		response, err := handler(c, ctx, &request)
		if err != nil {
			_ = ctx.Error(err)
		} else {
			ctx.JSON(http.StatusOK, biz.Success(response))
		}
	}
}

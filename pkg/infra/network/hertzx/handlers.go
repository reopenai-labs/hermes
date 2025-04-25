package hertzx

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"net/http"
	"qiyibyte.com/hermes/pkg/biz"
	"qiyibyte.com/hermes/pkg/builtin/constx"
	"qiyibyte.com/hermes/pkg/i18n"
	"qiyibyte.com/hermes/pkg/infra/metrics/bench"
	"strings"
)

func handNotFund() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(http.StatusNotFound, biz.ApiResponse{
			Code: biz.CodeNotFound,
			Message: i18n.GetMessageWithArgs(c, biz.CodeNotFound, map[string]string{
				"url": string(ctx.Request.Path()),
			}),
		})
	}
}

func handleAcceptLanguage() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		lang := ctx.Request.Header.Get(consts.HeaderAcceptLanguage)
		if lang == "" {
			lang = ctx.Request.Header.Get("lang")
			if lang == "" {
				lang = ctx.Request.Header.Get("locale")
				if lang == "" {
					lang = i18n.GetDefaultLanguage().String()
				}
			}
		}
		ctx.Next(context.WithValue(c, constx.ContextKeyLocale, lang))
	}
}

func handleMethodNotAllow() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(http.StatusMethodNotAllowed, biz.ApiResponse{
			Code: biz.CodeMethodNotAllow,
			Message: i18n.GetMessageWithArgs(c, biz.CodeMethodNotAllow, map[string]string{
				"method": string(ctx.Request.Method()),
			}),
		})
	}
}

func handleRecovery() app.HandlerFunc {
	return recovery.Recovery(recovery.WithRecoveryHandler(
		func(c context.Context, ctx *app.RequestContext, err any, stack []byte) {
			globalErrorHandler(c, ctx, ctx.Errors.Last())
		}),
	)
}

func handleErrors() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		ctx.Next(c)
		if err := ctx.Errors.Last(); err != nil {
			globalErrorHandler(c, ctx, err)
			ctx.Abort()
		}
	}
}

func handleBenchMarker() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		marker := bench.NewMarker()
		defer func() {
			statusCode := ctx.Response.StatusCode()
			marker.Mark("outer")
			if statusCode != http.StatusOK || marker.Total() > 200 {
				var builder strings.Builder
				builder.WriteString("HTTP: [startTime=")
				builder.WriteString(fmt.Sprintf("%d", marker.StartTime))
				builder.WriteString("][requestId=")
				builder.WriteString(c.Value(constx.ContextKeyRequestId).(string))
				builder.WriteString("][code=")
				builder.WriteString(fmt.Sprintf("%d", statusCode))
				builder.WriteString("][url=")
				builder.Write(ctx.Request.Method())
				builder.WriteRune(' ')
				builder.Write(ctx.Request.Path())
				builder.WriteRune(']')
				builder.WriteString(marker.String())
				hlog.CtxInfof(c, builder.String())
			}
		}()
		marker.Mark("enter")
		ctx.Next(context.WithValue(c, constx.ContextKeyBenchMarker, marker))
	}
}

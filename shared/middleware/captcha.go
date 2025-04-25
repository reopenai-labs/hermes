package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"qiyibyte.com/hermes/apps/common/facade"
	"qiyibyte.com/hermes/apps/common/facade/captchaapi"
	"qiyibyte.com/hermes/pkg/biz"
)

func CaptchaWebMiddleware() app.HandlerFunc {
	var captchaService = facade.Create(captchaapi.NewValidationClient)
	return func(c context.Context, ctx *app.RequestContext) {
		version := ctx.GetHeader("X-CAPTCHA-VERSION")
		if version == nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, biz.NewValidError("Missing X-CAPTCHA-VERSION in Request Header"))
			return
		}
		payload := ctx.GetHeader("X-CAPTCHA-CF-PAYLOAD")
		if payload == nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, biz.NewValidError("Missing X-CAPTCHA-CF-PAYLOAD in Request Header"))
			return
		}
		reply, err := captchaService.Validate(c, &captchaapi.ValidationRequest{
			Version: string(version),
			Payload: map[string]string{
				"response": string(payload),
			},
		})
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		if !reply.Success {

		}
		ctx.Next(c)
	}
}

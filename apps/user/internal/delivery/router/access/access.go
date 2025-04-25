package access

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"qiyibyte.com/hermes/apps/user/internal/domain/po"
	"qiyibyte.com/hermes/apps/user/internal/service"
	"qiyibyte.com/hermes/pkg/infra/network/hertzx"
	"qiyibyte.com/hermes/shared/middleware"
)

func Register(hz *server.Hertz) {
	hz.POST("/public/login", middleware.CaptchaWebMiddleware(), hertzx.RequestBodyHandler(access))
}

func access(c context.Context, ctx *app.RequestContext, request *po.LoginRequest) (any, error) {
	return service.Access().Login(c, request)
}

package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"time"
)

func Register(hz *server.Hertz) {
	hz.GET("hello", func(c context.Context, ctx *app.RequestContext) {
		time.Sleep(1 * time.Second)
		ctx.JSON(200, "hello world")
	})
}

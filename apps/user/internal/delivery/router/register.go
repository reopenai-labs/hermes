package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"qiyibyte.com/hermes/apps/user/internal/delivery/router/access"
)

func Register(hz *server.Hertz) {
	access.Register(hz)
}

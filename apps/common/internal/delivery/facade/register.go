package facade

import (
	"google.golang.org/grpc"
	"qiyibyte.com/hermes/apps/common/facade/captchaapi"
	"qiyibyte.com/hermes/apps/common/internal/delivery/facade/captcha"
)

func Register(server *grpc.Server) {
	captchaapi.RegisterValidationServer(server, &captcha.ValidatorFacade{})
}

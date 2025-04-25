package facade

import (
	"google.golang.org/grpc"
	"qiyibyte.com/hermes/apps/common/facade/captchaapi"
)

func Register(server *grpc.Server) {
	captchaapi.RegisterValidationServer(server, &ValidatorFacade{})
}

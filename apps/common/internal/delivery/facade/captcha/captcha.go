package captcha

import (
	"context"
	"qiyibyte.com/hermes/apps/common/facade/captchaapi"
)

type ValidatorFacade struct {
	captchaapi.UnimplementedValidationServer
}

func (Self *ValidatorFacade) Validate(ctx context.Context, request *captchaapi.ValidationRequest) (*captchaapi.ValidationReply, error) {
	return nil, nil
}

package hertzx

import (
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"qiyibyte.com/hermes/pkg/biz"
)

func newValidation() *binding.ValidateConfig {
	validateConfig := &binding.ValidateConfig{}
	validateConfig.SetValidatorErrorFactory(func(failField, msg string) error {
		return biz.NewParamError(failField, msg)
	})
	return validateConfig
}

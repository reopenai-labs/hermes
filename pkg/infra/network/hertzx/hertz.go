package hertzx

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzZerolog "github.com/hertz-contrib/logger/zerolog"
	"github.com/hertz-contrib/requestid"
	"github.com/rs/zerolog/log"
	"qiyibyte.com/hermes/pkg/builtin/constx"
	"qiyibyte.com/hermes/pkg/configs"
	"time"
)

func init() {
	configs.SetDefault("server.port", "8000")
	hlog.SetLogger(hertzZerolog.From(log.Logger))
}

func NewServer() *server.Hertz {
	hz := server.New(
		server.WithKeepAlive(true),
		server.WithKeepAliveTimeout(time.Minute*2),
		server.WithValidateConfig(newValidation()),
		server.WithHostPorts(":"+configs.GetString("server.port")),
	)
	hz.NoRoute(handNotFund())
	hz.NoMethod(handleMethodNotAllow())
	hz.Use(requestid.New(requestid.WithCustomHeaderStrKey(constx.ContextKeyRequestId)))
	hz.Use(handleBenchMarker())
	hz.Use(handleRecovery())
	hz.Use(handleAcceptLanguage())
	hz.Use(handleErrors())
	return hz
}

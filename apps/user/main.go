package main

import (
	"qiyibyte.com/hermes/apps/user/internal/delivery/facade"
	"qiyibyte.com/hermes/apps/user/internal/delivery/router"
	"qiyibyte.com/hermes/pkg/infra/network/grpcx"
	"qiyibyte.com/hermes/pkg/infra/network/hertzx"
)

func main() {
	initGrpcServer()
	initWebServer()
}

func initGrpcServer() {
	server := grpcx.NewServer()
	facade.Register(server)
	grpcx.Span(server)
}

func initWebServer() {
	hz := hertzx.NewServer()
	router.Register(hz)
	hz.Spin()
}

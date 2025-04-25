package grpcx

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"qiyibyte.com/hermes/pkg/builtin/eventbus"
	"qiyibyte.com/hermes/pkg/configs"
	"qiyibyte.com/hermes/pkg/logger"
)

func init() {
	configs.SetDefault("grpc.server.port", "8001")
}

func NewServer(opts ...grpc.ServerOption) *grpc.Server {
	options := []grpc.ServerOption{
		grpc.UnaryInterceptor(serverContextInterceptor()),
	}
	server := grpc.NewServer(append(options, opts...)...)
	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(server, healthcheck)
	return server
}

func Span(server *grpc.Server) {
	go func() {
		listen, err := net.Listen("tcp", ":"+configs.GetString("grpc.server.port"))
		if err != nil {
			logger.Fatalf("gRPC: 创建Listen失败: %v", err)
		}
		if err = server.Serve(listen); err != nil {
			logger.Fatalf("gRPC: 服务启动失败.error: %v", err)
		}
	}()
	eventbus.Subscribe(eventbus.TopicShutdown, func(args ...any) {
		logger.Infof("gRPC: 正在优雅停机...")
		server.GracefulStop()
		logger.Infof("gRPC: 已完成优雅停机")
	})
}

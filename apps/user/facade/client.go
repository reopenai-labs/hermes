package facade

import (
	"google.golang.org/grpc"
	"log"
	"qiyibyte.com/hermes/pkg/infra/network/grpcx"
)

var conn *grpc.ClientConn

func init() {
	client, err := grpcx.NewClient("localhost:8001")
	if err != nil {
		log.Fatalf("[gRPC][User]创建客户端连接失败.error: %v", err)
	}
	conn = client
}

func Create[T any](fn func(cc grpc.ClientConnInterface) T) T {
	return fn(conn)
}

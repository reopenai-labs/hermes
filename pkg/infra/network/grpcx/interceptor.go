package grpcx

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"qiyibyte.com/hermes/pkg/builtin/constx"
)

func serverContextInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			if locale := md.Get(constx.ContextKeyLocale); len(locale) > 0 {
				ctx = context.WithValue(ctx, constx.ContextKeyLocale, locale[0])
			}
			if requestId := md.Get(constx.ContextKeyRequestId); len(requestId) > 0 {
				ctx = context.WithValue(ctx, constx.ContextKeyRequestId, requestId[0])
			}
		}
		return handler(ctx, req)
	}
}

func clientContextInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		md := metadata.Pairs()
		if locale := ctx.Value(constx.ContextKeyLocale); locale != nil {
			if value, ok := locale.(string); ok {
				md.Set(constx.ContextKeyLocale, value)
			}
		}
		if requestId := ctx.Value(constx.ContextKeyRequestId); requestId != nil {
			if value, ok := requestId.(string); ok {
				md.Set(constx.ContextKeyRequestId, value)
			}
		}
		ctx = metadata.NewOutgoingContext(ctx, md)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

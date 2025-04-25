package redisx

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

//-------------------------
// 通用 API
//-------------------------

func ExpireTime(ctx context.Context, key string) *redis.DurationCmd {
	return getInstance().ExpireTime(ctx, key)
}

func Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return getInstance().Expire(ctx, key, expiration)
}

func ExpireNX(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return getInstance().ExpireNX(ctx, key, expiration)
}

func ExpireXX(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return getInstance().ExpireXX(ctx, key, expiration)
}

func ExpireGT(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return getInstance().ExpireLT(ctx, key, expiration)
}

func ExpireLT(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return getInstance().ExpireLT(ctx, key, expiration)
}

package redisx

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type StringCommand struct {
	client redis.Cmdable
}

func (Self *StringCommand) Append(ctx context.Context, key, value string) *redis.IntCmd {
	return Self.client.Append(ctx, key, value)
}

func (Self *StringCommand) Decr(ctx context.Context, key string) *redis.IntCmd {
	return Self.client.Decr(ctx, key)
}

func (Self *StringCommand) DecrBy(ctx context.Context, key string, decrement int64) *redis.IntCmd {
	return Self.client.DecrBy(ctx, key, decrement)
}

func (Self *StringCommand) Get(ctx context.Context, key string) *redis.StringCmd {
	return Self.client.Get(ctx, key)
}

func (Self *StringCommand) GetRange(ctx context.Context, key string, start, end int64) *redis.StringCmd {
	return Self.client.GetRange(ctx, key, start, end)
}

func (Self *StringCommand) GetSet(ctx context.Context, key string, value any) *redis.StringCmd {
	return Self.client.GetSet(ctx, key, value)
}

func (Self *StringCommand) GetEx(ctx context.Context, key string, expiration time.Duration) *redis.StringCmd {
	return Self.client.GetEx(ctx, key, expiration)
}

func (Self *StringCommand) GetDel(ctx context.Context, key string) *redis.StringCmd {
	return Self.client.GetDel(ctx, key)
}

func (Self *StringCommand) Incr(ctx context.Context, key string) *redis.IntCmd {
	return Self.client.Incr(ctx, key)
}

func (Self *StringCommand) IncrBy(ctx context.Context, key string, value int64) *redis.IntCmd {
	return Self.client.IncrBy(ctx, key, value)
}

func (Self *StringCommand) IncrByFloat(ctx context.Context, key string, value float64) *redis.FloatCmd {
	return Self.client.IncrByFloat(ctx, key, value)
}

func (Self *StringCommand) LCS(ctx context.Context, q *redis.LCSQuery) *redis.LCSCmd {
	return Self.client.LCS(ctx, q)
}

func (Self *StringCommand) MGet(ctx context.Context, keys ...string) *redis.SliceCmd {
	return Self.client.MGet(ctx, keys...)
}

func (Self *StringCommand) MSet(ctx context.Context, values ...any) *redis.StatusCmd {
	return Self.client.MSet(ctx, values...)
}

func (Self *StringCommand) MSetNX(ctx context.Context, values ...any) *redis.BoolCmd {
	return Self.client.MSetNX(ctx, values...)
}

func (Self *StringCommand) Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd {
	return Self.client.Set(ctx, key, value, expiration)
}

func (Self *StringCommand) SetArgs(ctx context.Context, key string, value any, a redis.SetArgs) *redis.StatusCmd {
	return Self.client.SetArgs(ctx, key, value, a)
}

func (Self *StringCommand) SetEx(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd {
	return Self.client.SetEx(ctx, key, value, expiration)
}

func (Self *StringCommand) SetNX(ctx context.Context, key string, value any, expiration time.Duration) *redis.BoolCmd {
	return Self.client.SetNX(ctx, key, value, expiration)
}

func (Self *StringCommand) SetXX(ctx context.Context, key string, value any, expiration time.Duration) *redis.BoolCmd {
	return Self.client.SetXX(ctx, key, value, expiration)
}

func (Self *StringCommand) SetRange(ctx context.Context, key string, offset int64, value string) *redis.IntCmd {
	return Self.client.SetRange(ctx, key, offset, value)
}

func (Self *StringCommand) StrLen(ctx context.Context, key string) *redis.IntCmd {
	return Self.client.StrLen(ctx, key)
}

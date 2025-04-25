package redisx

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type HashCommand struct {
	client redis.Cmdable
}

//-------------------------
// Hash API
//-------------------------

func (Self *HashCommand) GetAll(ctx context.Context, key string) *redis.MapStringStringCmd {
	return Self.client.HGetAll(ctx, key)
}

func (Self *HashCommand) Del(ctx context.Context, key string, fields ...string) *redis.IntCmd {
	return Self.client.HDel(ctx, key, fields...)
}

func (Self *HashCommand) Exists(ctx context.Context, key, field string) *redis.BoolCmd {
	return Self.client.HExists(ctx, key, field)
}

func (Self *HashCommand) Get(ctx context.Context, key, field string) *redis.StringCmd {
	return Self.client.HGet(ctx, key, field)
}

func (Self *HashCommand) IncrBy(ctx context.Context, key, field string, incr int64) *redis.IntCmd {
	return Self.client.HIncrBy(ctx, key, field, incr)
}

func (Self *HashCommand) IncrByFloat(ctx context.Context, key, field string, incr float64) *redis.FloatCmd {
	return Self.client.HIncrByFloat(ctx, key, field, incr)
}

func (Self *HashCommand) Keys(ctx context.Context, key string) *redis.StringSliceCmd {
	return Self.client.HKeys(ctx, key)
}

func (Self *HashCommand) Len(ctx context.Context, key string) *redis.IntCmd {
	return Self.client.HLen(ctx, key)
}

func (Self *HashCommand) MGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd {
	return Self.client.HMGet(ctx, key, fields...)
}

func (Self *HashCommand) Set(ctx context.Context, key string, values ...any) *redis.IntCmd {
	return Self.client.HSet(ctx, key, values...)
}

func (Self *HashCommand) MSet(ctx context.Context, key string, values ...any) *redis.BoolCmd {
	return Self.client.HMSet(ctx, key, values...)
}

func (Self *HashCommand) SetNX(ctx context.Context, key, field string, value any) *redis.BoolCmd {
	return Self.client.HSetNX(ctx, key, field, value)
}

func (Self *HashCommand) Vals(ctx context.Context, key string) *redis.StringSliceCmd {
	return Self.client.HVals(ctx, key)
}

func (Self *HashCommand) RandField(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	return Self.client.HRandField(ctx, key, count)
}

func (Self *HashCommand) RandFieldWithValues(ctx context.Context, key string, count int) *redis.KeyValueSliceCmd {
	return Self.client.HRandFieldWithValues(ctx, key, count)
}

func (Self *HashCommand) Scan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return Self.client.HScan(ctx, key, cursor, match, count)
}

func (Self *HashCommand) ScanNoValues(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return Self.client.HScanNoValues(ctx, key, cursor, match, count)
}

func (Self *HashCommand) Expire(ctx context.Context, key string, expiration time.Duration, fields ...string) *redis.IntSliceCmd {
	return Self.client.HExpire(ctx, key, expiration, fields...)
}

func (Self *HashCommand) ExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	return Self.client.HExpireWithArgs(ctx, key, expiration, expirationArgs, fields...)
}

func (Self *HashCommand) PExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *redis.IntSliceCmd {
	return Self.client.HPExpire(ctx, key, expiration, fields...)
}

func (Self *HashCommand) PExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	return Self.client.HPExpireWithArgs(ctx, key, expiration, expirationArgs, fields...)
}

func (Self *HashCommand) ExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *redis.IntSliceCmd {
	return Self.client.HExpireAt(ctx, key, tm, fields...)
}

func (Self *HashCommand) ExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	return Self.client.HExpireAtWithArgs(ctx, key, tm, expirationArgs, fields...)
}

func (Self *HashCommand) PExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *redis.IntSliceCmd {
	return Self.client.HPExpireAt(ctx, key, tm, fields...)
}

func (Self *HashCommand) PExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	return Self.client.HPExpireAtWithArgs(ctx, key, tm, expirationArgs, fields...)
}

func (Self *HashCommand) Persist(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	return Self.client.HPersist(ctx, key, fields...)
}

func (Self *HashCommand) ExpireTime(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	return Self.client.HExpireTime(ctx, key, fields...)
}

func (Self *HashCommand) PExpireTime(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	return Self.client.HPExpireTime(ctx, key, fields...)
}

func (Self *HashCommand) TTL(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	return Self.client.HTTL(ctx, key, fields...)
}

func (Self *HashCommand) PTTL(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	return Self.client.HPTTL(ctx, key, fields...)
}

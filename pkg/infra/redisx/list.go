package redisx

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var listCommand *ListCommand
var listCommandInit sync.Once

func List() *ListCommand {
	listCommandInit.Do(func() {
		listCommand = &ListCommand{client: getInstance()}
	})
	return listCommand
}

type ListCommand struct {
	client redis.Cmdable
}

func (Self *ListCommand) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	return Self.client.BLPop(ctx, timeout, keys...)
}

func (Self *ListCommand) BLMPop(ctx context.Context, timeout time.Duration, direction string, count int64, keys ...string) *redis.KeyValuesCmd {
	return Self.client.BLMPop(ctx, timeout, direction, count, keys...)
}

func (Self *ListCommand) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	return Self.client.BRPop(ctx, timeout, keys...)
}

func (Self *ListCommand) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *redis.StringCmd {
	return Self.client.BRPopLPush(ctx, source, destination, timeout)
}

func (Self *ListCommand) Index(ctx context.Context, key string, index int64) *redis.StringCmd {
	return Self.client.LIndex(ctx, key, index)
}

func (Self *ListCommand) MPop(ctx context.Context, direction string, count int64, keys ...string) *redis.KeyValuesCmd {
	return Self.client.LMPop(ctx, direction, count, keys...)
}

func (Self *ListCommand) Insert(ctx context.Context, key, op string, pivot, value interface{}) *redis.IntCmd {
	return Self.client.LInsert(ctx, key, op, pivot, value)
}

func (Self *ListCommand) InsertBefore(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	return Self.client.LInsertBefore(ctx, key, pivot, value)
}

func (Self *ListCommand) InsertAfter(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	return Self.client.LInsertAfter(ctx, key, pivot, value)
}

func (Self *ListCommand) Len(ctx context.Context, key string) *redis.IntCmd {
	return Self.client.LLen(ctx, key)
}

func (Self *ListCommand) Pop(ctx context.Context, key string) *redis.StringCmd {
	return Self.client.LPop(ctx, key)
}

func (Self *ListCommand) PopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	return Self.client.LPopCount(ctx, key, count)
}

func (Self *ListCommand) Pos(ctx context.Context, key string, value string, a redis.LPosArgs) *redis.IntCmd {
	return Self.client.LPos(ctx, key, value, a)
}

func (Self *ListCommand) PosCount(ctx context.Context, key string, value string, count int64, a redis.LPosArgs) *redis.IntSliceCmd {
	return Self.client.LPosCount(ctx, key, value, count, a)
}

func (Self *ListCommand) Push(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return Self.client.LPush(ctx, key, values...)
}

func (Self *ListCommand) PushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return Self.client.LPushX(ctx, key, values...)
}

func (Self *ListCommand) Range(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	return Self.client.LRange(ctx, key, start, stop)
}

func (Self *ListCommand) Rem(ctx context.Context, key string, count int64, value interface{}) *redis.IntCmd {
	return Self.client.LRem(ctx, key, count, value)
}

func (Self *ListCommand) Set(ctx context.Context, key string, index int64, value interface{}) *redis.StatusCmd {
	return Self.client.LSet(ctx, key, index, value)
}

func (Self *ListCommand) Trim(ctx context.Context, key string, start, stop int64) *redis.StatusCmd {
	return Self.client.LTrim(ctx, key, start, stop)
}

func (Self *ListCommand) RPop(ctx context.Context, key string) *redis.StringCmd {
	return Self.client.RPop(ctx, key)
}

func (Self *ListCommand) RPopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	return Self.client.RPopCount(ctx, key, count)
}

func (Self *ListCommand) RPopLPush(ctx context.Context, source, destination string) *redis.StringCmd {
	return Self.client.RPopLPush(ctx, source, destination)
}

func (Self *ListCommand) RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return Self.client.RPush(ctx, key, values...)
}

func (Self *ListCommand) RPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return Self.client.RPushX(ctx, key, values...)
}

func (Self *ListCommand) Move(ctx context.Context, source, destination, srcpos, destpos string) *redis.StringCmd {
	return Self.client.LMove(ctx, source, destination, srcpos, destpos)
}

func (Self *ListCommand) BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *redis.StringCmd {
	return Self.client.BLMove(ctx, source, destination, srcpos, destpos, timeout)
}

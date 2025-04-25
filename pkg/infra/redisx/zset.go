package redisx

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

// ZSetCommand 是 Redis ZSet 操作的封装结构体
type ZSetCommand struct {
	client redis.Cmdable
}

func (Self *ZSetCommand) BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	return Self.client.BZPopMax(ctx, timeout, keys...)
}

func (Self *ZSetCommand) BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	return Self.client.BZPopMin(ctx, timeout, keys...)
}

func (Self *ZSetCommand) BZMPop(ctx context.Context, timeout time.Duration, order string, count int64, keys ...string) *redis.ZSliceWithKeyCmd {
	return Self.client.BZMPop(ctx, timeout, order, count, keys...)
}

func (Self *ZSetCommand) AddArgs(ctx context.Context, key string, args redis.ZAddArgs) *redis.IntCmd {
	return Self.client.ZAddArgs(ctx, key, args)
}

func (Self *ZSetCommand) AddArgsIncr(ctx context.Context, key string, args redis.ZAddArgs) *redis.FloatCmd {
	return Self.client.ZAddArgsIncr(ctx, key, args)
}

func (Self *ZSetCommand) Add(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	return Self.client.ZAdd(ctx, key, members...)
}

func (Self *ZSetCommand) AddLT(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	return Self.client.ZAddLT(ctx, key, members...)
}

func (Self *ZSetCommand) AddGT(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	return Self.client.ZAddGT(ctx, key, members...)
}

func (Self *ZSetCommand) AddNX(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	return Self.client.ZAddNX(ctx, key, members...)
}

func (Self *ZSetCommand) AddXX(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	return Self.client.ZAddXX(ctx, key, members...)
}

func (Self *ZSetCommand) Card(ctx context.Context, key string) *redis.IntCmd {
	return Self.client.ZCard(ctx, key)
}

func (Self *ZSetCommand) Count(ctx context.Context, key, min, max string) *redis.IntCmd {
	return Self.client.ZCount(ctx, key, min, max)
}

func (Self *ZSetCommand) LexCount(ctx context.Context, key, min, max string) *redis.IntCmd {
	return Self.client.ZLexCount(ctx, key, min, max)
}

func (Self *ZSetCommand) IncrBy(ctx context.Context, key string, increment float64, member string) *redis.FloatCmd {
	return Self.client.ZIncrBy(ctx, key, increment, member)
}

func (Self *ZSetCommand) InterStore(ctx context.Context, destination string, store *redis.ZStore) *redis.IntCmd {
	return Self.client.ZInterStore(ctx, destination, store)
}

func (Self *ZSetCommand) Inter(ctx context.Context, store *redis.ZStore) *redis.StringSliceCmd {
	return Self.client.ZInter(ctx, store)
}

func (Self *ZSetCommand) InterWithScores(ctx context.Context, store *redis.ZStore) *redis.ZSliceCmd {
	return Self.client.ZInterWithScores(ctx, store)
}

func (Self *ZSetCommand) InterCard(ctx context.Context, limit int64, keys ...string) *redis.IntCmd {
	return Self.client.ZInterCard(ctx, limit, keys...)
}

func (Self *ZSetCommand) MPop(ctx context.Context, order string, count int64, keys ...string) *redis.ZSliceWithKeyCmd {
	return Self.client.ZMPop(ctx, order, count, keys...)
}

func (Self *ZSetCommand) MScore(ctx context.Context, key string, members ...string) *redis.FloatSliceCmd {
	return Self.client.ZMScore(ctx, key, members...)
}

func (Self *ZSetCommand) PopMax(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd {
	return Self.client.ZPopMax(ctx, key, count...)
}

func (Self *ZSetCommand) PopMin(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd {
	return Self.client.ZPopMin(ctx, key, count...)
}

func (Self *ZSetCommand) RangeArgs(ctx context.Context, z redis.ZRangeArgs) *redis.StringSliceCmd {
	return Self.client.ZRangeArgs(ctx, z)
}

func (Self *ZSetCommand) RangeArgsWithScores(ctx context.Context, z redis.ZRangeArgs) *redis.ZSliceCmd {
	return Self.client.ZRangeArgsWithScores(ctx, z)
}

func (Self *ZSetCommand) Range(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	return Self.client.ZRange(ctx, key, start, stop)
}

func (Self *ZSetCommand) RangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	return Self.client.ZRangeWithScores(ctx, key, start, stop)
}

func (Self *ZSetCommand) RangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	return Self.client.ZRangeByScore(ctx, key, opt)
}

func (Self *ZSetCommand) RangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	return Self.client.ZRangeByLex(ctx, key, opt)
}

func (Self *ZSetCommand) RangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	return Self.client.ZRangeByScoreWithScores(ctx, key, opt)
}

func (Self *ZSetCommand) RangeStore(ctx context.Context, dst string, z redis.ZRangeArgs) *redis.IntCmd {
	return Self.client.ZRangeStore(ctx, dst, z)
}

func (Self *ZSetCommand) Rank(ctx context.Context, key, member string) *redis.IntCmd {
	return Self.client.ZRank(ctx, key, member)
}

func (Self *ZSetCommand) RankWithScore(ctx context.Context, key, member string) *redis.RankWithScoreCmd {
	return Self.client.ZRankWithScore(ctx, key, member)
}

func (Self *ZSetCommand) Rem(ctx context.Context, key string, members ...any) *redis.IntCmd {
	return Self.client.ZRem(ctx, key, members...)
}

func (Self *ZSetCommand) RemRangeByRank(ctx context.Context, key string, start, stop int64) *redis.IntCmd {
	return Self.client.ZRemRangeByRank(ctx, key, start, stop)
}

func (Self *ZSetCommand) RemRangeByScore(ctx context.Context, key, min, max string) *redis.IntCmd {
	return Self.client.ZRemRangeByScore(ctx, key, min, max)
}

func (Self *ZSetCommand) RemRangeByLex(ctx context.Context, key, min, max string) *redis.IntCmd {
	return Self.client.ZRemRangeByLex(ctx, key, min, max)
}

func (Self *ZSetCommand) RevRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	return Self.client.ZRevRange(ctx, key, start, stop)
}

func (Self *ZSetCommand) RevRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	return Self.client.ZRevRangeWithScores(ctx, key, start, stop)
}

func (Self *ZSetCommand) RevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	return Self.client.ZRevRangeByScore(ctx, key, opt)
}

func (Self *ZSetCommand) RevRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	return Self.client.ZRevRangeByLex(ctx, key, opt)
}

func (Self *ZSetCommand) RevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	return Self.client.ZRevRangeByScoreWithScores(ctx, key, opt)
}

func (Self *ZSetCommand) RevRank(ctx context.Context, key, member string) *redis.IntCmd {
	return Self.client.ZRevRank(ctx, key, member)
}

func (Self *ZSetCommand) RevRankWithScore(ctx context.Context, key, member string) *redis.RankWithScoreCmd {
	return Self.client.ZRevRankWithScore(ctx, key, member)
}

func (Self *ZSetCommand) Score(ctx context.Context, key, member string) *redis.FloatCmd {
	return Self.client.ZScore(ctx, key, member)
}

func (Self *ZSetCommand) Union(ctx context.Context, store redis.ZStore) *redis.StringSliceCmd {
	return Self.client.ZUnion(ctx, store)
}

func (Self *ZSetCommand) UnionWithScores(ctx context.Context, store redis.ZStore) *redis.ZSliceCmd {
	return Self.client.ZUnionWithScores(ctx, store)
}

func (Self *ZSetCommand) UnionStore(ctx context.Context, dest string, store *redis.ZStore) *redis.IntCmd {
	return Self.client.ZUnionStore(ctx, dest, store)
}

// ZRandMember redis-server version >= 6.2.0.
func (Self *ZSetCommand) RandMember(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	return Self.client.ZRandMember(ctx, key, count)
}

// ZRandMemberWithScores redis-server version >= 6.2.0.
func (Self *ZSetCommand) RandMemberWithScores(ctx context.Context, key string, count int) *redis.ZSliceCmd {
	return Self.client.ZRandMemberWithScores(ctx, key, count)
}

// ZDiff redis-server version >= 6.2.0.
func (Self *ZSetCommand) Diff(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	return Self.client.ZDiff(ctx, keys...)
}

// ZDiffWithScores redis-server version >= 6.2.0.
func (Self *ZSetCommand) DiffWithScores(ctx context.Context, keys ...string) *redis.ZSliceCmd {
	return Self.client.ZDiffWithScores(ctx, keys...)
}

// ZDiffStore redis-server version >=6.2.0.
func (Self *ZSetCommand) DiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	return Self.client.ZDiffStore(ctx, destination, keys...)
}

func (Self *ZSetCommand) Scan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return Self.client.ZScan(ctx, key, cursor, match, count)
}

package redisx

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type SetCommand struct {
	client redis.Cmdable
}

//-------------------------
// Set API
//-------------------------

func (Self *SetCommand) Add(ctx context.Context, key string, members ...any) *redis.IntCmd {
	return Self.client.SAdd(ctx, key, members...)
}

func (Self *SetCommand) Card(ctx context.Context, key string) *redis.IntCmd {
	return Self.client.SCard(ctx, key)
}

func (Self *SetCommand) Diff(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	return Self.client.SDiff(ctx, keys...)
}

func (Self *SetCommand) DiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	return Self.client.SDiffStore(ctx, destination, keys...)
}

func (Self *SetCommand) Inter(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	return Self.client.SInter(ctx, keys...)
}

func (Self *SetCommand) InterCard(ctx context.Context, limit int64, keys ...string) *redis.IntCmd {
	return Self.client.SInterCard(ctx, limit, keys...)
}

func (Self *SetCommand) InterStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	return Self.client.SInterStore(ctx, destination, keys...)
}

func (Self *SetCommand) IsMember(ctx context.Context, key string, member any) *redis.BoolCmd {
	return Self.client.SIsMember(ctx, key, member)
}

func (Self *SetCommand) MIsMember(ctx context.Context, key string, members ...any) *redis.BoolSliceCmd {
	return Self.client.SMIsMember(ctx, key, members...)
}

func (Self *SetCommand) Members(ctx context.Context, key string) *redis.StringSliceCmd {
	return Self.client.SMembers(ctx, key)
}

func (Self *SetCommand) MembersMap(ctx context.Context, key string) *redis.StringStructMapCmd {
	return Self.client.SMembersMap(ctx, key)
}

func (Self *SetCommand) Move(ctx context.Context, source, destination string, member any) *redis.BoolCmd {
	return Self.client.SMove(ctx, source, destination, member)
}

func (Self *SetCommand) Pop(ctx context.Context, key string) *redis.StringCmd {
	return Self.client.SPop(ctx, key)
}

func (Self *SetCommand) PopN(ctx context.Context, key string, count int64) *redis.StringSliceCmd {
	return Self.client.SPopN(ctx, key, count)
}

func (Self *SetCommand) RandMember(ctx context.Context, key string) *redis.StringCmd {
	return Self.client.SRandMember(ctx, key)
}

func (Self *SetCommand) RandMemberN(ctx context.Context, key string, count int64) *redis.StringSliceCmd {
	return Self.client.SRandMemberN(ctx, key, count)
}

func (Self *SetCommand) Rem(ctx context.Context, key string, members ...any) *redis.IntCmd {
	return Self.client.SRem(ctx, key, members...)
}

func (Self *SetCommand) Scan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return Self.client.SScan(ctx, key, cursor, match, count)
}

func (Self *SetCommand) Union(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	return Self.client.SUnion(ctx, keys...)
}

func (Self *SetCommand) UnionStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	return Self.client.SUnionStore(ctx, destination, keys...)
}

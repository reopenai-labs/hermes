package redisx

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
)

var pubSubCommand *PubSubCommand
var pubSubCommandInit sync.Once

func PubSub() *PubSubCommand {
	pubSubCommandInit.Do(func() {
		pubSubCommand = &PubSubCommand{client: getInstance()}
	})
	return pubSubCommand
}

//-------------------------
// pub-sub API
//-------------------------

type PubSubCommand struct {
	client redis.Cmdable
}

func (Self *PubSubCommand) Publish(ctx context.Context, channel string, message any) *redis.IntCmd {
	return Self.client.Publish(ctx, channel, message)
}

func (Self *PubSubCommand) SPublish(ctx context.Context, channel string, message any) *redis.IntCmd {
	return Self.client.SPublish(ctx, channel, message)
}

func (Self *PubSubCommand) PubSubChannels(ctx context.Context, pattern string) *redis.StringSliceCmd {
	return Self.client.PubSubChannels(ctx, pattern)
}

func (Self *PubSubCommand) PubSubNumSub(ctx context.Context, channels ...string) *redis.MapStringIntCmd {
	return Self.client.PubSubNumSub(ctx, channels...)
}

func (Self *PubSubCommand) PubSubShardChannels(ctx context.Context, pattern string) *redis.StringSliceCmd {
	return Self.client.PubSubShardChannels(ctx, pattern)
}

func (Self *PubSubCommand) PubSubShardNumSub(ctx context.Context, channels ...string) *redis.MapStringIntCmd {
	return Self.client.PubSubShardNumSub(ctx, channels...)
}

func (Self *PubSubCommand) PubSubNumPat(ctx context.Context) *redis.IntCmd {
	return Self.client.PubSubNumPat(ctx)
}

package redisx

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type PipelineCommand struct {
	client redis.Pipeliner
}

func (Self *PipelineCommand) Do(ctx context.Context, args ...interface{}) *redis.Cmd {
	return Self.client.Do(ctx, args...)
}

func (Self *PipelineCommand) Process(ctx context.Context, cmd redis.Cmder) error {
	return Self.client.Process(ctx, cmd)
}

func (Self *PipelineCommand) Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return Self.client.Pipelined(ctx, fn)
}

func (Self *PipelineCommand) Pipeline() redis.Pipeliner {
	return Self.client.Pipeline()
}

func (Self *PipelineCommand) TxPipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return Self.client.TxPipelined(ctx, fn)
}

func (Self *PipelineCommand) TxPipeline() redis.Pipeliner {
	return Self.client.TxPipeline()
}

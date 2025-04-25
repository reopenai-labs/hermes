package logger

import (
	"context"
	"github.com/rs/zerolog/log"
)

func Debug(msg string) {
	log.Debug().Msg(msg)
}

func DebugWithCtx(ctx context.Context, msg string) {
	log.Debug().Ctx(ctx).Msg(msg)
}

func Debugf(format string, args ...any) {
	log.Debug().Msgf(format, args...)
}

func DebugfWithCtx(ctx context.Context, format string, args ...any) {
	log.Debug().Ctx(ctx).Msgf(format, args...)
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func InfoWithCtx(ctx context.Context, msg string) {
	log.Info().Ctx(ctx).Msg(msg)
}

func Infof(format string, args ...any) {
	log.Info().Msgf(format, args...)
}

func InfofWithCtx(ctx context.Context, format string, args ...any) {
	log.Info().Ctx(ctx).Msgf(format, args...)
}

func Error(msg string) {
	log.Error().Msg(msg)
}

func ErrorWithCtx(ctx context.Context, msg string) {
	log.Error().Ctx(ctx).Msg(msg)
}

func Errorf(format string, args ...any) {
	log.Error().Msgf(format, args...)
}

func ErrorfWithCtx(ctx context.Context, format string, args ...any) {
	log.Error().Ctx(ctx).Msgf(format, args...)
}

func Fatal(msg string) {
	log.Fatal().Msg(msg)
}

func FatalWithCtx(ctx context.Context, msg string) {
	log.Fatal().Ctx(ctx).Msg(msg)
}

func Fatalf(format string, args ...any) {
	log.Fatal().Msgf(format, args...)
}

func FatalfWithCtx(ctx context.Context, format string, args ...any) {
	log.Fatal().Ctx(ctx).Msgf(format, args...)
}

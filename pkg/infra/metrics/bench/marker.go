package bench

import (
	"context"
	"fmt"
	"qiyibyte.com/hermes/pkg/builtin/constx"
	"qiyibyte.com/hermes/pkg/utils/ctxutil"
	"qiyibyte.com/hermes/pkg/utils/timeutil"
	"strings"
)

type Marker struct {
	builder   strings.Builder
	StartTime int64
	EndTime   int64
	result    string
	total     int64
}

func (Self *Marker) Mark(flag string) {
	now := timeutil.CachedNow().UnixMilli()
	Self.total = now - Self.StartTime
	Self.builder.WriteRune(' ')
	Self.builder.WriteString(flag)
	Self.builder.WriteRune(':')
	Self.builder.WriteString(fmt.Sprintf("%d", now-Self.EndTime))
	Self.builder.WriteRune(':')
	Self.builder.WriteString(fmt.Sprintf("%d", Self.total))
	Self.EndTime = now
}

func (Self *Marker) String() string {
	if Self.result == "" {
		Self.builder.WriteString(" Total=")
		Self.builder.WriteString(fmt.Sprintf("%d", Self.total))
		Self.result = Self.builder.String()
	}
	return Self.result
}

func (Self *Marker) Total() int64 {
	return Self.total
}

func Mark(ctx context.Context, flag string) {
	marker := ctxutil.GetPtrValue[Marker](ctx, constx.ContextKeyBenchMarker)
	if marker != nil {
		marker.Mark(flag)
	}
}

func NewMarker() *Marker {
	now := timeutil.CachedNow().UnixMilli()
	return &Marker{
		StartTime: now,
		EndTime:   now,
		builder:   strings.Builder{},
	}
}

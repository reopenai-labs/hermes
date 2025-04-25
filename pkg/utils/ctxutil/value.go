package ctxutil

import "context"

func GetPtrValue[T any](ctx context.Context, key any) *T {
	value := ctx.Value(key)
	if value == nil {
		return nil
	}
	if data, ok := value.(*T); ok {
		return data
	}
	return nil
}

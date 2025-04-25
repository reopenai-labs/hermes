package timeutil

import (
	"qiyibyte.com/hermes/pkg/builtin/sys"
	"time"
)

var now = time.Now()

func init() {
	go refreshNow()
}

func CachedNow() time.Time {
	return now
}

func refreshNow() {
	ticker := time.NewTicker(time.Millisecond)
	for {
		select {
		case <-ticker.C:
			now = time.Now()
		case <-sys.Done():
			break
		}
	}
}

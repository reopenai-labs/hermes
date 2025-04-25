package sys

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type shutdownHookFunc func()

var instance context.Context
var cancelFunc context.CancelFunc
var shutdownHooks []shutdownHookFunc

func init() {
	instance, cancelFunc = context.WithCancel(context.Background())
	go signalWatcher()
}

func Done() <-chan struct{} {
	return instance.Done()
}

func Err() error {
	return instance.Err()
}

func Value(key any) any {
	return instance.Value(key)
}

func Deadline() (deadline time.Time, ok bool) {
	return instance.Deadline()
}

func ShutdownHook(hook func()) {
	shutdownHooks = append(shutdownHooks, hook)
}

func signalWatcher() {
	signalToNotify := []os.Signal{syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM}
	if signal.Ignored(syscall.SIGHUP) {
		signalToNotify = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, signalToNotify...)
	for sig := range signals {
		switch sig {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
			cancelFunc()
			for _, hook := range shutdownHooks {
				hook()
			}
		}
	}
}

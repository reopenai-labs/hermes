package eventbus

import (
	"qiyibyte.com/hermes/pkg/builtin/sys"
	"sync"
)

var bus = NewEventBus()

func init() {
	sys.ShutdownHook(func() {
		Publish(TopicShutdown)
	})
}

func Publish(topic string, args ...any) {
	bus.Publish(topic, args...)
}

// Subscribe to events
func Subscribe(topic string, fn func(args ...any)) {
	bus.Subscribe(topic, fn)
}

func NewEventBus() *EventBus {
	return &EventBus{
		channels: make(map[string]*eventChannel),
		locker:   sync.Mutex{},
	}
}

type eventHandler func(args ...any)

type EventBus struct {
	channels map[string]*eventChannel
	locker   sync.Mutex
}

func (Self *EventBus) Publish(topic string, args ...any) {
	Self.getChannel(topic).publish(args...)
}

func (Self *EventBus) Subscribe(topic string, fn func(args ...any)) {
	Self.getChannel(topic).subscribe(fn)
}

func (Self *EventBus) getChannel(topic string) *eventChannel {
	channel, ok := Self.channels[topic]
	if !ok {
		Self.locker.Lock()
		defer Self.locker.Unlock()
		channel, ok = Self.channels[topic]
		if !ok {
			channel = &eventChannel{handlers: make([]eventHandler, 0), topic: topic}
			Self.channels[topic] = channel
		}
	}
	return channel
}

type eventChannel struct {
	topic    string
	lock     sync.RWMutex
	handlers []eventHandler
}

func (Self *eventChannel) publish(args ...any) {
	Self.lock.RLock()
	defer Self.lock.RUnlock()
	for _, handler := range Self.handlers {
		handler(args...)
	}
}

func (Self *eventChannel) subscribe(handler eventHandler) {
	Self.lock.Lock()
	defer Self.lock.Unlock()
	Self.handlers = append(Self.handlers, handler)
}

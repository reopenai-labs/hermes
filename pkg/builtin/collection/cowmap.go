package collection

import (
	"sync"
	"sync/atomic"
)

type CowMap[K comparable, V any] struct {
	entities map[K]V
	snapshot atomic.Pointer[map[K]V]
	locker   sync.Mutex
}

func NewCowMap[K comparable, V any]() *CowMap[K, V] {
	m := &CowMap[K, V]{
		entities: make(map[K]V),
	}
	initial := make(map[K]V)
	m.snapshot.Store(&initial)
	return m
}

func (Self *CowMap[K, V]) Put(key K, value V) {
	Self.locker.Lock()
	defer Self.locker.Unlock()
	Self.entities[key] = value
	Self.createSnapshot()
}

func (Self *CowMap[K, V]) PutAll(entities map[K]V) {
	Self.locker.Lock()
	defer Self.locker.Unlock()
	for k, v := range entities {
		Self.entities[k] = v
	}
	Self.createSnapshot()
}

func (Self *CowMap[K, V]) Remove(key K) {
	Self.locker.Lock()
	defer Self.locker.Unlock()
	delete(Self.entities, key)
	Self.createSnapshot()
}

func (Self *CowMap[K, V]) Get(key K) (V, bool) {
	snap := Self.snapshot.Load()
	value, ok := (*snap)[key]
	return value, ok
}

func (Self *CowMap[K, V]) GetOrCreate(key K, fn func() V) V {
	value, ok := Self.Get(key)
	if !ok {
		Self.locker.Lock()
		defer Self.locker.Unlock()
		value, ok = Self.Get(key)
		if !ok {
			value = fn()
			Self.entities[key] = value
			Self.createSnapshot()
		}
	}
	return value
}

func (Self *CowMap[K, V]) Range(fn func(key K, value V) bool) {
	snap := Self.snapshot.Load()
	for k, v := range *snap {
		if !fn(k, v) {
			break
		}
	}
}

func (Self *CowMap[K, V]) createSnapshot() {
	snapshot := make(map[K]V, len(Self.entities))
	for k, v := range Self.entities {
		snapshot[k] = v
	}
	Self.snapshot.Store(&snapshot)
}

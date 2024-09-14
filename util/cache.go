package util

import (
	"time"
)

type CacheItem[T any] struct {
	expiresAt time.Time
	data      T
}

type Cache[T any] map[string]CacheItem[T]

func NewCache[T any]() Cache[T] {
	return Cache[T]{}
}

func (m *Cache[T]) Get(k string) T {
	if c, ok := (*m)[k]; ok && c.expiresAt.After(time.Now()) {
		return c.data
	} else {
		if ok { // exists
			m.Delete(k) // clear
		}
		return ZeroValue[T]()
	}
}

func (m *Cache[T]) Set(k string, v T, ttl time.Duration) {
	(*m)[k] = CacheItem[T]{time.Now().Add(ttl), v}
}

func (m *Cache[T]) Delete(k string) {
	delete(*m, k)
}

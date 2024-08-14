package ratelimit

import (
	"context"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type RateLimit[T any] struct {
	duration    time.Duration
	burst       int
	mapKey      func(r T) string // returning "" represents to be invalid
	ctx         context.Context
	cancel      context.CancelFunc
	mu          *sync.Mutex
	limiterMap  map[string]*rate.Limiter
	lastTimeMap map[string]time.Time
}

func New[T any](duration time.Duration, burst int, mapKey func(r T) string) *RateLimit[T] {
	ctx, cancel := context.WithCancel(context.Background())
	mu := new(sync.Mutex)
	limiterMap := make(map[string]*rate.Limiter)
	lastTimeMap := make(map[string]time.Time)
	l := &RateLimit[T]{duration, burst, mapKey, ctx, cancel, mu, limiterMap, lastTimeMap}

	// 自动启动IP清理的goroutine
	go l.keepCleaning()
	return l
}

func (l *RateLimit[T]) Destroy() {
	l.cancel()
}

func (l *RateLimit[T]) Allow(r T) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	key := l.mapKey(r)
	if key == "" { // invalid key
		return false // always reject
	}
	l.lastTimeMap[key] = time.Now()

	if _, exists := l.limiterMap[key]; !exists {
		l.limiterMap[key] = rate.NewLimiter(rate.Every(l.duration), l.burst)
	}
	return l.limiterMap[key].Allow()
}

// 定期清理不再活跃的IP地址
func (l *RateLimit[T]) keepCleaning() {
	// interval := l.duration * 5
	interval := l.duration + 2*time.Minute
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// for range ticker.C {
	for {
		select {
		case <-l.ctx.Done():
			return // exit and stop
		case <-ticker.C:
			func() {
				l.mu.Lock()
				defer l.mu.Unlock()

				currentTime := time.Now()
				for key, lastTime := range l.lastTimeMap {
					if currentTime.Sub(lastTime) > interval { // 如果超过指定间隔无请求
						delete(l.limiterMap, key)
						delete(l.lastTimeMap, key)
					}
				}
			}()
		}
	}
}

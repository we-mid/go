package ratelimit

import (
	"net/http"
	"sync"
	"time"

	"gitee.com/we-mid/go/bec_http"
	"golang.org/x/time/rate"
)

type IpRateLimiter struct {
	duration    time.Duration
	burst       int
	mu          *sync.Mutex
	limiterMap  map[string]*rate.Limiter
	lastTimeMap map[string]time.Time
}

func NewIpRateLimiter(duration time.Duration, burst int) *IpRateLimiter {
	mu := new(sync.Mutex)
	limiterMap := make(map[string]*rate.Limiter)
	lastTimeMap := make(map[string]time.Time)
	l := &IpRateLimiter{duration, burst, mu, limiterMap, lastTimeMap}

	// 启动IP清理的goroutine
	go l.keepCleaning()
	return l
}

func (l *IpRateLimiter) Allow(r *http.Request) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	ip := bec_http.GetClientIP(r)
	l.lastTimeMap[ip] = time.Now()

	if _, exists := l.limiterMap[ip]; !exists {
		l.limiterMap[ip] = rate.NewLimiter(rate.Every(l.duration), l.burst)
	}
	return l.limiterMap[ip].Allow()
}

// cleanupOldIPs 定期清理不再活跃的IP地址
func (l *IpRateLimiter) keepCleaning() {
	interval := l.duration * 5
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		func() {
			l.mu.Lock()
			defer l.mu.Unlock()

			currentTime := time.Now()
			for ip, lastReqTime := range l.lastTimeMap {
				if currentTime.Sub(lastReqTime) > interval { // 如果超过指定间隔无请求
					delete(l.limiterMap, ip)
					delete(l.lastTimeMap, ip)
				}
			}
		}()
	}
}

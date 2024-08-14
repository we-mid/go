package bec_http

import (
	"net/http"
	"time"

	rl "gitee.com/we-mid/go/ratelimit"
)

func NewIPRateLimit(duration time.Duration, burst int) *rl.RateLimit[*http.Request] {
	// 注意 此处RateLimiter会启动一个goroutine进行keepCleaning
	return rl.New(duration, burst, func(r *http.Request) string {
		return GetClientIP(r)
	})
}

func RateLimitWrap(l *rl.RateLimit[*http.Request], logic Logic) Logic {
	return func(w http.ResponseWriter, r *http.Request) (any, error) {
		if !l.Allow(r) {
			return nil, Err429
		}
		return logic(w, r)
	}
}

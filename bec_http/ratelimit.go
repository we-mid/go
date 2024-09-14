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

func RateLimitWrap[T any](l *rl.RateLimit[*http.Request], logic Logic[T]) Logic[T] {
	return func(w http.ResponseWriter, r *http.Request) (T, error) {
		if !l.Allow(r) {
			// fix: compiler: cannot use nil as T value in return statement
			// return nil, Err429
			// fix: compiler: invalid composite literal type T
			// return T{}, Err429
			var zero T
			return zero, Err429
		}
		return logic(w, r)
	}
}

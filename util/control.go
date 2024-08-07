package util

import (
	"sync"
	"time"
)

// throttledFunction wraps an original function with throttle logic.
type throttledFunction func()

// Throttle creates and returns a throttled version of the given function that will only be executed once every `interval`.
func Throttle(interval time.Duration, originalFunc func()) throttledFunction {
	var mutex sync.Mutex
	lastExecutionTime := time.Time{}

	return func() {
		mutex.Lock()
		defer mutex.Unlock()

		now := time.Now()
		timeSinceLastExecution := now.Sub(lastExecutionTime)

		if timeSinceLastExecution >= interval {
			originalFunc()
			lastExecutionTime = now
		}
	}
}

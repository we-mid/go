package util

import (
	"sync"
	"time"
)

// Example usage:
// throttledFunc := Throttle(time.Second*3, func() { fmt.Println("Executed") })
// throttledFunc() // Will execute immediately.
// throttledFunc() // Will not execute until 3 seconds have passed from the last execution.

func Throttle(interval time.Duration, originalFunc func()) func() {
	var timer *time.Timer
	var mutex sync.Mutex
	lastExecutionTime := time.Time{}

	return func() {
		mutex.Lock()
		var onceUnlocked bool
		defer func() {
			if !onceUnlocked {
				mutex.Unlock()
			}
		}()

		if timer != nil {
			return // skip
		}
		// elapsed := time.Since(lastExecutionTime)
		elapsed := time.Now().Sub(lastExecutionTime)
		wait := interval - elapsed

		if wait <= 0 {
			// 如果已经超过了间隔时间，直接执行函数并更新时间
			lastExecutionTime = time.Now()
			// go originalFunc() // defer
			onceUnlocked = true
			mutex.Unlock() // 提前解锁
			originalFunc()
			return
		}
		// 如果还没有超过间隔时间，设置定时器
		timer = time.AfterFunc(wait, func() {
			timer = nil
			// 注意：这里不需要再加锁，因为我们是在定时器回调中直接操作
			lastExecutionTime = time.Now()
			originalFunc()
			// 定时器会自动被停止，并且timer变量将不再被引用（成为垃圾回收的目标）
		})
	}
}

package util

import (
	"sync"
	"sync/atomic"
	"syscall"
)

type ReentrantLock struct {
	mutex  sync.Mutex
	holder int64 // 当前线程/协程ID，用于记录锁的持有者
	count  int32 // 重入计数
}

func (r *ReentrantLock) Lock() {
	currentGoroutine := getThreadId()
	if atomic.LoadInt64(&r.holder) == int64(currentGoroutine) {
		r.count++
		return
	}

	r.mutex.Lock()
	atomic.StoreInt64(&r.holder, int64(currentGoroutine))
	r.count = 1
}

func (r *ReentrantLock) Unlock() {
	currentGoroutine := getThreadId()
	if atomic.LoadInt64(&r.holder) != int64(currentGoroutine) {
		panic("Unlock called by a non-lock holder")
	}
	if r.count > 1 {
		r.count--
		return
	}

	atomic.StoreInt64(&r.holder, 0)
	r.mutex.Unlock()
}

func getThreadId() int64 {
	// if runtime.GOOS == "windows" {
	// 	return syscall.GetCurrentThreadId()
	// }
	// if runtime.GOOS == "linux" {
	// 	return syscall.Gettid()
	// }
	// return 0 // 简化示例，实际情况可能需要更复杂的协程ID获取方式

	// ???
	return syscall.SYS_THREAD_SELFID
}

// func main() {
// 	lock := &ReentrantLock{}

// 	go func() {
// 		lock.Lock()
// 		fmt.Println("First lock acquired")
// 		lock.Lock() // 重入
// 		fmt.Println("Second lock acquired")
// 		lock.Unlock()
// 		fmt.Println("First unlock")
// 		lock.Unlock()
// 		fmt.Println("Second unlock")
// 	}()

// 	go func() {
// 		runtime.Gosched() // 确保上面的goroutine先执行Lock
// 		lock.Lock()
// 		fmt.Println("Another goroutine acquired the lock")
// 		lock.Unlock()
// 		fmt.Println("Another goroutine released the lock")
// 	}()

// 	select {}
// }

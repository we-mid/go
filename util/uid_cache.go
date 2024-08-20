package util

import (
	"sync"
	"time"
)

// UIDCache 结构体用于存储UID及其接收时间
type UIDCache struct {
	// mu sync.RWMutex
	mu         sync.Mutex
	cancel     chan struct{}
	store      map[string]time.Time // 存储UID及其接收时间
	timeWindow time.Duration        // 时间窗口，例如2分钟
}

// NewUIDCache 创建一个新的UIDCache实例
func NewUIDCache(timeWindow time.Duration) *UIDCache {
	c := &UIDCache{
		cancel:     make(chan struct{}),
		store:      make(map[string]time.Time),
		timeWindow: timeWindow,
	}
	go c.keepGC()
	return c
}

func (c *UIDCache) Destroy() {
	c.cancel <- struct{}{}
}

// Add 尝试添加一个UID到存储中，如果UID在过去的时间窗口内已存在，则返回false表示幂等
func (c *UIDCache) Add(UID string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 检查UID是否已经在时间窗口内存在
	if t, exists := c.store[UID]; exists {
		// 如果存在，且未过期，则不处理，返回false
		if time.Since(t) <= c.timeWindow {
			return false
			// } else {
			// 	// 如果已过期，则删除旧记录，允许添加新记录
			// 	delete(c.store, UID)
		}
	}
	// 添加新UID及其接收时间
	c.store[UID] = time.Now()
	return true
}

// StartGC 启动垃圾回收，定期清理过期的UID
func (c *UIDCache) keepGC() {
	for {
		select {
		case <-c.cancel:
			return // exit loop
		case <-time.After(2 * c.timeWindow):
			c.cleanExpired()
		}
	}
}

// cleanExpired 清理过期的UID
func (c *UIDCache) cleanExpired() {
	c.mu.Lock()
	defer c.mu.Unlock()
	var n, total int
	currentTime := time.Now()
	for UID, t := range c.store {
		if currentTime.Sub(t) > c.timeWindow {
			delete(c.store, UID)
			n++
		}
		total++
	}
	// if total > 0 {
	// 	log.Printf("[UIDCache] 已清理 %d / %d 个key\n", n, total)
	// }
}

// func main() {
// 	// 创建一个2分钟时间窗口的UID存储实例
// 	c := NewUIDCache(2 * time.Minute)

// 	// 示例：尝试添加相同的UID
// 	UID := "some-unique-id"
// 	fmt.Println("Adding UID for the first time:", c.Add(UID)) // 应该输出true
// 	time.Sleep(1 * time.Minute)                                   // 等待一段时间
// 	fmt.Println("Adding the same UID again:", c.Add(UID))    // 在时间窗口内，应输出false，表示幂等
// }

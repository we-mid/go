package util

import (
	"testing"
	"time"

	"sync/atomic"
)

func TestThrottle(t *testing.T) {
	var calls int32
	throttle := Throttle(100*time.Millisecond, func() { atomic.AddInt32(&calls, 1) })

	// 快速连续调用多次，期望在第一个100ms内只被计数一次
	for i := 0; i < 10; i++ {
		throttle()
	}
	time.Sleep(200 * time.Millisecond) // 等待超过节流间隔
	if atomic.LoadInt32(&calls) != 2 {
		t.Errorf("Expected 2 call, got %d", atomic.LoadInt32(&calls))
	}

	// 再次快速连续调用多次，这次应该在下一个100ms周期开始时再次被计数
	atomic.StoreInt32(&calls, 0) // reset
	for i := 0; i < 10; i++ {
		throttle()
	}
	time.Sleep(200 * time.Millisecond) // 等待超过节流间隔
	if atomic.LoadInt32(&calls) != 2 {
		t.Errorf("Expected 2 calls, got %d", atomic.LoadInt32(&calls))
	}

	// 检查在节流间隔内调用是否会增加计数
	atomic.StoreInt32(&calls, 0) // reset
	throttle()
	time.Sleep(50 * time.Millisecond)
	throttle()
	if atomic.LoadInt32(&calls) != 1 {
		t.Errorf("Expected no additional calls within throttle interval")
	}

	// 检查在节流间隔之后调用是否会增加计数
	time.Sleep(60 * time.Millisecond) // 确保超过节流间隔
	throttle()
	if atomic.LoadInt32(&calls) != 2 {
		t.Errorf("Expected an additional call after throttle interval")
	}
}

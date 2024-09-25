package util

import "sync"

type GoroutineWaitGroup struct {
	sync.WaitGroup
}

func (c *GoroutineWaitGroup) Go(fn func()) {
	c.Add(1) // 注意 wg.Add 必须放在 go func 之外
	go func() {
		defer c.Done()
		fn()
	}()
}

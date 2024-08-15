package util

import (
	"sync"
	"testing"
)

type worker struct {
	id   int
	done bool
}

func (w *worker) Close() {
	w.done = true
}

func TestPool(t *testing.T) {
	// 突出潜在的并发问题
	// 如 fatal error: concurrent map iteration and map write`
	// n := 1000
	n := 100000
	// 1. 串行
	for i := 0; i < n; i++ {
		runOnce(t)
	}
	// 2. 并行
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			runOnce(t)
			wg.Done()
		}()
	}
	wg.Wait()
}

func runOnce(t *testing.T) {
	// new
	pool := NewPool[*worker](3, func() (*worker, error) {
		return &worker{}, nil
	}, func(w *worker) error {
		w.Close()
		return nil
	})
	n, c := len(pool.refs), len(pool.ch)
	if n != 0 {
		t.Fatalf("n: want 0, got %d", n)
	}
	if c != 0 {
		t.Fatalf("c: want 0, got %d", c)
	}

	// get
	c1, err := pool.Get()
	if err != nil {
		// (*testing.common).Fatalf does not support error-wrapping directive %w
		// t.Fatalf("pool.Get: %w", err)
		t.Fatalf("pool.Get: %s", err)
	}
	n, c = len(pool.refs), len(pool.ch)
	if n != 1 {
		t.Fatalf("n: want 1, got %d", n)
	}
	if c != 0 {
		t.Fatalf("c: want 0, got %d", c)
	}
	c2, _ := pool.Get()
	n, c = len(pool.refs), len(pool.ch)
	if n != 2 {
		t.Fatalf("n: want 2, got %d", n)
	}
	if c != 0 {
		t.Fatalf("c: want 0, got %d", c)
	}
	c3, _ := pool.Get()
	n, c = len(pool.refs), len(pool.ch)
	if n != 3 {
		t.Fatalf("n: want 3, got %d", n)
	}
	if c != 0 {
		t.Fatalf("c: want 0, got %d", c)
	}

	// block
	var c4 *worker
	select {
	// case c4, _ = pool.Get():
	case c4 = <-pool.ch:
		t.Fatalf("pool.ch: did not block")
	default:
		// expect
	}
	n, c = len(pool.refs), len(pool.ch)
	if n != 3 {
		t.Fatalf("n: want 3, got %d", n)
	}
	if c != 0 {
		t.Fatalf("c: want 0, got %d", c)
	}

	// put
	err = pool.Put(c1)
	if err != nil {
		t.Fatalf("pool.Put: %s", err)
	}
	n, c = len(pool.refs), len(pool.ch)
	if n != 3 {
		t.Fatalf("n: want 3, got %d", n)
	}
	if c != 1 {
		t.Fatalf("c: want 1, got %d", c)
	}
	pool.Put(c2)
	n, c = len(pool.refs), len(pool.ch)
	if n != 3 {
		t.Fatalf("n: want 3, got %d", n)
	}
	if c != 2 {
		t.Fatalf("c: want 2, got %d", c)
	}
	pool.Put(&worker{})
	n, c = len(pool.refs), len(pool.ch)
	if n != 4 {
		t.Fatalf("n: want 4, got %d", n)
	}
	if c != 3 {
		t.Fatalf("c: want 3, got %d", c)
	}
	pool.Put(&worker{})
	n, c = len(pool.refs), len(pool.ch)
	if n != 4 {
		t.Fatalf("n: want 4, got %d", n)
	}
	if c != 3 {
		t.Fatalf("c: want 3, got %d", c)
	}
	pool.Put(c3)
	n, c = len(pool.refs), len(pool.ch)
	if n != 3 {
		t.Fatalf("n: want 3, got %d", n)
	}
	if c != 3 {
		t.Fatalf("c: want 3, got %d", c)
	}
	_ = c4 // ignore

	// destroy
	if err = pool.Destroy(); err != nil {
		t.Fatalf("pool.Destroy: %s", err)
	}
	n, _ = len(pool.refs), len(pool.ch)
	if n != 0 {
		t.Fatalf("n: want 0, got %d", n)
	}
	if _, ok := <-pool.ch; ok {
		t.Fatal("pool.ch: should be closed")
	}
	if !pool.destroyed {
		t.Fatalf("pool.destroyed: want true, got %v", pool.destroyed)
	}
	if _, err := pool.Get(); err != errDestroyed {
		t.Fatalf("pool.Get: want errDestroyed, got %v", err)
	}
}

package util

import (
	"errors"
	"sync"
)

var errDestroyed = errors.New("pool has been destroyed")

// 连接池结构体
type Pool[T comparable] struct {
	maxConns  int            // 最大连接数量
	refs      map[T]struct{} // 字典用于收集连接实例引用
	ch        chan T         // 通道用于排队获取连接实例
	mu        sync.Mutex     // 互斥锁用于同步连接的创建
	create    func() (T, error)
	free      func(T) error
	destroyed bool
}

// 创建一个新的连接池
func NewPool[T comparable](maxConns int, create func() (T, error), free func(T) error) *Pool[T] {
	return &Pool[T]{
		maxConns: maxConns,
		refs:     make(map[T]struct{}),
		ch:       make(chan T, maxConns),
		create:   create,
		free:     free,
	}
}

// 释放/存放一个连接实例到连接池
func (p *Pool[T]) Put(conn T) error {
	if p.destroyed {
		return errDestroyed
	}
	p.mu.Lock()
	defer p.mu.Unlock()

	select {
	case p.ch <- conn:
		p.refs[conn] = struct{}{}
	default: // 通道已满
		if err := p.free(conn); err != nil { // abandan
			return err // skip, keep in set to expose
		}
		delete(p.refs, conn)
	}
	return nil
}

// 获取一个连接实例
func (p *Pool[T]) Get() (T, error) {
	if p.destroyed {
		return *new(T), errDestroyed
	}
	p.mu.Lock()
	defer p.mu.Unlock()

	select {
	case conn := <-p.ch:
		return conn, nil
	default: // 通道为空
		// 如果达到最大连接数量，阻塞直到有空闲的连接
		// if len(p.ch) >= p.maxConns { // buggy
		if len(p.refs) >= p.maxConns {
			return <-p.ch, nil
		}
		conn, err := p.create()
		if err != nil {
			// return nil, err
			// return T{}, err
			// return T(nil), err
			return *new(T), err
		}
		p.refs[conn] = struct{}{}
		return conn, nil
	}
}

func (p *Pool[T]) Destroy() error {
	p.destroyed = true
	close(p.ch)
	for range p.ch {
		// remove all
	}
	var err error
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(p.refs))

	// fix: fatal error: concurrent map iteration and map write (2)
	var conns []T
	for conn := range p.refs {
		conns = append(conns, conn)
	}
	// for conn := range p.refs {
	for _, conn := range conns {
		// 循环闭包问题：在 Destroy 方法中，你提到注释中有一个循环闭包问题。
		// 这是因为 Go 的循环闭包捕获的是循环变量的引用，而不是值。
		// 这意味着所有 goroutines 都会引用同一个 conn 变量，导致错误的行为。
		// 解决这个问题的方法是在闭包中传递每个 conn 的副本。
		c := conn // 捕获当前迭代的值
		// note: loopclosure: loop variable conn captured by func literal
		go func(c T) {
			defer wg.Done()
			if e := p.free(c); e != nil {
				if err == nil {
					err = e // return the first error
				}
				return
			}
			// fix: fatal error: concurrent map writes (1)
			mu.Lock()
			defer mu.Unlock()
			delete(p.refs, c)
		}(c)
	}
	wg.Wait()
	return err
}

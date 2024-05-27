package util

import (
	"errors"
	"os"
	"syscall"
)

// Go语言巧用文件锁避免多个进程同时存在的问题
// https://linkscue.com/posts/2018-09-07-golang-flock-example/
// File Locking In Go Part
// https://magdy.dev/blog/2021/02/07/file-locking-in-go-part-i/
// https://magdy.dev/blog/2021/02/14/file-locking-in-go-part-ii/

type Flock struct {
	Filename string
	file     *os.File
}

// 创建文件锁，配合 defer f.Release() 来使用
func FlockCreate(filename string) (f *Flock, e error) {
	if filename == "" {
		e = errors.New("cannot create flock on empty path")
		return
	}
	file, e := os.Create(filename)
	if e != nil {
		return
	}
	return &Flock{
		Filename: filename,
		file:     file,
	}, nil
}

// 释放文件锁
func (f *Flock) Release() {
	if f != nil && f.file != nil {
		f.file.Close()
		os.Remove(f.Filename)
	}
}

// 上锁，配合 defer f.Unlock() 来使用
func (f *Flock) Lock() (e error) {
	if f == nil {
		e = errors.New("cannot use lock on a nil flock")
		return
	}
	return syscall.Flock(int(f.file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
}

// 解锁
func (f *Flock) Unlock() {
	if f != nil {
		syscall.Flock(int(f.file.Fd()), syscall.LOCK_UN)
	}
}

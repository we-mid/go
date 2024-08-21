package db

import (
	"fmt"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

// https://github.com/lionsoul2014/ip2region/tree/master/binding/golan
// var dbPath = "ip2region.xdb"

// 方式一、完全基于文件的查询:
// 备注：并发使用，每个 goroutine 需要创建一个独立的 searcher 对象。
func Load01FileOnly(dbPath string) (*xdb.Searcher, error) {
	searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create searcher: %w\n", err)
	}
	return searcher, nil
}

// 方式二、缓存 VectorIndex 索引:
// 备注：并发使用，全部 goroutine 共享全局的只读 vIndex 缓存，每个 goroutine 创建一个独立的 searcher 对象
func Load02IndexCache(dbPath string) (*xdb.Searcher, error) {
	// // 1、从 dbPath 加载 VectorIndex 缓存，把下述 vIndex 变量全局到内存里面。
	vIndex, err := xdb.LoadVectorIndexFromFile(dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load vector index from `%s`: %w", dbPath, err)
	}
	// // 2、用全局的 vIndex 创建带 VectorIndex 缓存的查询对象。
	searcher, err := xdb.NewWithVectorIndex(dbPath, vIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to create searcher with vector index: %w", err)
	}
	return searcher, nil
}

// 方式三、缓存整个 xdb 数据：
// 备注：并发使用，用整个 xdb 缓存创建的 searcher 对象可以安全用于并发。
func Load03FullCache(dbPath string) (*xdb.Searcher, error) {
	// 1、从 dbPath 加载整个 xdb 到内存
	cBuff, err := xdb.LoadContentFromFile(dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load content from `%s`: %w", dbPath, err)
	}
	// 2、用全局的 cBuff 创建完全基于内存的查询对象。
	searcher, err := xdb.NewWithBuffer(cBuff)
	if err != nil {
		return nil, fmt.Errorf("failed to create searcher with content: %w", err)
	}
	return searcher, nil
}

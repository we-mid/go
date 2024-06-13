package util

import "sort"

// LessFunc 定义一个比较函数类型
type LessFunc[T comparable] func(T, T) bool

// SortedKeys 获取并返回一个已排序的map键切片
func SortedKeys[M ~map[K]V, K comparable, V any](m M, less LessFunc[K]) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return less(keys[i], keys[j])
	})
	return keys
}

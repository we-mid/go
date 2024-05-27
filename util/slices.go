package util

// RemoveElements 从 slice 中移除 elements 中的所有元素
func RemoveElements[T comparable](slice, elements []T) []T {
	seen := make(map[T]struct{}, len(elements))
	for _, v := range elements {
		seen[v] = struct{}{}
	}
	// 创建一个新的切片来存储结果
	result := make([]T, 0, len(slice)-len(elements))
	for _, v := range slice {
		if _, ok := seen[v]; !ok {
			// 如果 v 不在 elements 中，则将其添加到结果切片中
			result = append(result, v)
		}
	}
	return result
}

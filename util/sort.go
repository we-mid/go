package util

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// var reDigitSuffix = regexp.MustCompile(`\.log(\.\d+)?$`) // 只支持整数
var reDigitSuffix = regexp.MustCompile(`\.log(\.\d+(\.\d*)?)?$`) // 支持小数 灵活操控

// 数字大的表示越早的 排在前面
func SortLogFilesDigitSuffix(filenames []string) {
	sort.Slice(filenames, func(i, j int) bool {
		// 提取文件名中的数字部分
		matchesI := reDigitSuffix.FindStringSubmatch(filenames[i])
		matchesJ := reDigitSuffix.FindStringSubmatch(filenames[j])
		// 比较数字部分，没有数字的视为0
		// var indexI, indexJ int
		var indexI, indexJ float64
		if matchesI[1] != "" {
			indexStrI := strings.TrimPrefix(matchesI[1], ".")
			// indexI, _ = strconv.Atoi(indexStrI) // 忽略错误，因为我们已经用正则匹配了
			indexI, _ = strconv.ParseFloat(indexStrI, 64)
		}
		if matchesJ[1] != "" {
			indexStrJ := strings.TrimPrefix(matchesJ[1], ".")
			// indexJ, _ = strconv.Atoi(indexStrJ) // 忽略错误，因为我们已经用正则匹配了
			indexJ, _ = strconv.ParseFloat(indexStrJ, 64)
		}
		// 数字大的排在前面
		return indexJ < indexI
	})
}

func RangeMapSorted[M ~map[K]V, K comparable, V any](m M, c LessFunc[K], f func(k K, v V) bool) {
	keys := SortedKeys(m, c)
	for _, k := range keys {
		if !f(k, m[k]) {
			break
		}
	}
}

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

// 用法 util.RangeMapSorted(myMap, util.StrLess, func (...) bool {...})
func StrLess(a, b string) bool {
	return a < b
}

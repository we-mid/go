package util

import (
	"math"
	"strings"
)

// CountOccurrences counts the number of occurrences of subStr in s.
func CountOccurrences(s, subStr string) int {
	return CountOccurrencesN(s, subStr, math.MaxInt)
}
func CountOccurrencesN(s, subStr string, n int) int {
	if s == "" || subStr == "" {
		return 0
	}
	count := 0
	startIndex := 0
	for count < n && startIndex < len(s) {
		// Find the next occurrence of subStr starting from startIndex
		index := strings.Index(s[startIndex:], subStr)
		if index == -1 {
			break // No more occurrences found
		}
		count++
		// Move startIndex forward to continue search after the current match
		startIndex += index + len(subStr)
	}
	return count
}

// func CountOccurrencesBuggy(s, subStr string) int {
// 	count := 0
// 	index := 0
// 	for index < len(s) {
// 		// Find the next occurrence of subStr starting from index
// 		index = strings.Index(s[index:], subStr)
// 		if index == -1 {
// 			break // No more occurrences found
// 		}
// 		count++
// 		// Move index forward to continue search after the current match
// 		index += len(subStr)
// 	}
// 	return count
// }

package util

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	s        string
	subStr   string
	expected int
}{
	{"", "", 0},
	{"", "aaa", 0},
	{"aaa", "", 0},
	{"hello world", "world", 1},
	{"hello world world", "world", 2},
	{"hellohellohello", "hello", 3},
	{"aaaabaaaacaaaad", "aaaa", 3},
	{"abcdefg", "h", 0},
	{"ababababab", "ab", 5},
}

// TestCountOccurrences tests the CountOccurrences function with various inputs.
func TestCountOccurrences(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s_in_%s", tc.subStr, tc.s), func(t *testing.T) {
			result := CountOccurrences(tc.s, tc.subStr)
			if result != tc.expected {
				t.Errorf("CountOccurrences(%q, %q) = %d; want %d", tc.s, tc.subStr, result, tc.expected)
			}
		})
	}
}

// func TestCountOccurrencesBuggy(t *testing.T) {
// 	for _, tc := range testCases {
// 		t.Run(fmt.Sprintf("%s_in_%s", tc.subStr, tc.s), func(t *testing.T) {
// 			result := CountOccurrencesBuggy(tc.s, tc.subStr)
// 			if result != tc.expected {
// 				t.Errorf("CountOccurrences(%q, %q) = %d; want %d", tc.s, tc.subStr, result, tc.expected)
// 			}
// 		})
// 	}
// }

// BenchmarkCountOccurrences benchmarks the CountOccurrences function with different data sizes.
func BenchmarkCountOccurrences(b *testing.B) {
	for _, tc := range testCases {
		b.Run(fmt.Sprintf("%s_in_%s", tc.subStr, tc.s), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				CountOccurrences(tc.s, tc.subStr)
			}
		})
	}
}

// func BenchmarkCountOccurrencesBuggy(b *testing.B) {
// 	for _, tc := range testCases {
// 		b.Run(fmt.Sprintf("%s_in_%s", tc.subStr, tc.s), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				CountOccurrencesBuggy(tc.s, tc.subStr)
// 			}
// 		})
// 	}
// }

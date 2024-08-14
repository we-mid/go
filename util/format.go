package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	datePrintPattern = "01/02 15:04"
)

func RuneLength(s string) int {
	return len([]rune(s))
}
func TruncateString(s string, length int) string {
	runes := []rune(s)
	if len(runes) <= length {
		return s
	}
	return string(runes[:length])
}

// 四舍五入一个float64到指定位数，并返回格式化后的字符串
func FormatRound(num float64, precision int) string {
	// 计算小数点移动的倍数
	shift := math.Pow(10, float64(precision))
	// 四舍五入
	roundedNum := math.Round(num*shift) / shift
	// 使用Sprintf格式化，precision决定了小数点后保留的位数
	return fmt.Sprintf("%.*f", precision, roundedNum)
}

func PadWithSpaces(num int, width int) string {
	str := strconv.Itoa(num)
	padding := strings.Repeat(" ", max(0, width-len(str)))
	return padding + str
}

func PadWithSpacesStr(str string, width int) string {
	padding := strings.Repeat(" ", max(0, width-len(str)))
	return padding + str
}

func DatePrintf(layout string, args ...any) (int, error) {
	layout = "%s " + layout
	xargs := []any{time.Now().Format(datePrintPattern)}
	xargs = append(xargs, args...)
	return fmt.Printf(layout, xargs...)
}
func DatePrintln(args ...any) (int, error) {
	xargs := []any{time.Now().Format(datePrintPattern)}
	xargs = append(xargs, args...)
	return fmt.Println(xargs...)
}

func NormalizeTimestamp(t int64) int64 {
	if IsLikeJsTimestamp(t) {
		return t / 1000
	}
	return t
}
func IsLikeJsTimestamp(t int64) bool {
	return t >= 1000000000000
}

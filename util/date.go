package util

import (
	"fmt"
	"time"
)

// DaysBetween 计算两个日期之间相差的天数
// 假设输入的日期字符串都遵循 "YYYY-MM-DD" 格式
func DaysBetweenDates(dateStr, prevDateStr string) (int, error) {
	// 解析日期字符串为 time.Time 类型
	layout := "2006-01-02" // Go 的日期布局
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return 0, fmt.Errorf("无法解析日期 %s: %w", dateStr, err)
	}
	prev, err := time.Parse(layout, prevDateStr)
	if err != nil {
		return 0, fmt.Errorf("无法解析日期 %s: %w", prevDateStr, err)
	}
	return DaysBetween(t, prev)
}
func DaysBetween(t, prev time.Time) (int, error) {
	return int(t.Sub(prev).Hours() / 24), nil
}

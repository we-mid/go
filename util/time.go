package util

import (
	"fmt"
	"math"
	"time"
)

// format of extraTZ: `+0800`, `-0700`
func TimeParseWithExtraTZ(layout, timeStrWithoutTZ, extraTZ string) (time.Time, error) {
	l := fmt.Sprintf("%s %s", layout, "-0700")
	s := fmt.Sprintf("%s %s", timeStrWithoutTZ, extraTZ)
	return time.Parse(l, s)
}

// DurationScale takes a time.Duration and a float64 multiplier,
// returns the duration multiplied by the given multiplier.
func DurationScale(d time.Duration, mult float64) time.Duration {
	// 将小数转换为分数形式的整数比例
	numerator := int64(math.Round(mult * 1e9)) // 1e9代表十亿，即纳秒的数量
	denominator := int64(1e9)

	// 将时间间隔乘以分子，再除以分母
	return d * time.Duration(numerator) / time.Duration(denominator)
}

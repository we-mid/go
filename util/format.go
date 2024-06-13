package util

import (
	"strconv"
	"strings"
)

func PadWithSpaces(num int, width int) string {
	str := strconv.Itoa(num)
	padding := strings.Repeat(" ", width-len(str))
	return padding + str
}

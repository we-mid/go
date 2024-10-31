package util

import (
	"regexp"
	"unicode"
)

func ContainsChinese(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

// 正则表达式匹配标准的电子邮件地址
var reEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// isValidEmail 使用正则表达式检查给定的字符串是否符合电子邮件地址的标准格式
func IsValidEmail(email string) bool {
	return reEmail.MatchString(email)
}

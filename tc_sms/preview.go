package tc_sms

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

// 注意：关于单条短信计费fee -- 1. 汉字、字母、数字、标点符号（不区分全角/半角）以及空格等，都按1个字计算
// 2. 国内短信短信长度（签名+正文）不超过70字时，按照1条短信计费；超过70字即为长短信时，按67字/条分隔成多条计费，但会在1条短信内呈现。例如，短信长度为150字，则按照67字/67字/16字分隔成3条计费
// https://console.cloud.tencent.com/smsv2/csms-template/create
const nPerFee = 67

// 发送前预览完整文本，自动计算fee
func SmsPreview(template string, params []string) (string, int, int) {
	out := replacePlaceholders(template, params)
	n := len([]rune(out)) // 注意是 字符数
	return out, n, int(math.Ceil(float64(n) / float64(nPerFee)))
}

func replacePlaceholders(s string, params []string) string {
	// 正则表达式匹配形如 {数字} 的模式
	reTmpl := regexp.MustCompile(`\{(\d+)\}`)
	return reTmpl.ReplaceAllStringFunc(s, func(match string) string {
		// 提取数字部分并转换为整型
		indexStr := strings.Trim(match, "{}")
		index, err := strconv.Atoi(indexStr)
		index -= 1
		if err != nil {
			return match // 如果转换失败，返回原样
		}
		// 确保索引在切片范围内
		if index < 0 || index >= len(params) {
			return match // 如果索引越界，返回原样
		}
		return params[index]
	})
}

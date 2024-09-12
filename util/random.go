package util

import (
	crand "crypto/rand"
	"encoding/base64"
	"math"
	"math/rand"
)

func RandomBase64(lenStr int) (string, error) {
	// 计算至少需要的原始字节数量，向上取整
	minBytesNeeded := int(math.Ceil(float64(lenStr*3) / 4))
	// 生成原始字节
	b := make([]byte, minBytesNeeded)
	if _, err := crand.Read(b); err != nil {
		return "", err
	}
	// 进行Base64编码
	base64Str := base64.URLEncoding.EncodeToString(b)
	// 注意这可能会有'='填充字符 截断处理
	return base64Str[:lenStr], nil
}

func RandomID(lenStr int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits of letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	// 计算至少需要的原始字节数量，向上取整
	b := make([]byte, lenStr)
	for i, cache, remain := lenStr-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)[:lenStr]
}

func RandomCode(length int) string {
	// 创建一个长度为length的切片，用于存放生成的随机数
	randomNumbers := make([]byte, length)
	// 填充切片，生成指定长度的随机数
	// rand.Intn(base) 返回一个[0, base)之间的伪随机整数
	// 这里base是10，因为我们只需要0-9的数字
	for i := range randomNumbers {
		randomNumbers[i] = byte(rand.Intn(10)) + '0' // 将整数转换为对应的字符'0'-'9'
	}
	// 将切片转换为字符串
	return string(randomNumbers)
}

package node

import (
	crand "crypto/rand"
	"encoding/base64"
	"math"
)

func randomBase64ID(lenStr int) (string, error) {
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

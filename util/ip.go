package util

import (
	"net"
	"strings"
)

func ExtractIP(addr string) string {
	// 尝试将remoteAddr解析为TCPAddr，如果成功则直接获取IP
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err == nil {
		return tcpAddr.IP.String()
	}

	// 如果解析失败（可能是因为包含了非IP部分，如IPv6的"[::1]:80"），则使用字符串操作提取IP
	parts := strings.Split(addr, ":")
	if len(parts) > 0 {
		// 检查是否是IPv6地址（用方括号括起来）
		if strings.HasPrefix(parts[0], "[") && strings.HasSuffix(parts[0], "]") {
			return parts[0][1 : len(parts[0])-1] // 去掉方括号
		}
		return parts[0] // 假设是IPv4地址或其他没有端口的IP地址
	}

	// 如果以上都失败，返回空字符串或错误处理
	return ""
}

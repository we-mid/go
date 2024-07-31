package util

import (
	"fmt"
	"net"
	"strings"
)

func IsIPInCIDR(ip, cidr string) (bool, error) {
	// 解析CIDR
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false, fmt.Errorf("net.ParseCIDR: %w", err)
	}
	// 将IP地址字符串转换为net.IP类型
	ipObj := net.ParseIP(ip)
	// 检查IP是否在CIDR区间内
	return ipNet.Contains(ipObj), nil
}

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

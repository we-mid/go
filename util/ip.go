package util

import (
	"net"
	"strings"
)

// isIPInRange 检查给定的 IP 地址是否在指定的起始 IP 和结束 IP 之间
func IsIPInRange(ipStr, startStr, endStr string) bool {
	ip := net.ParseIP(ipStr)
	start := net.ParseIP(startStr)
	end := net.ParseIP(endStr)
	return isNetIPInRange(ip, start, end)
}
func isNetIPInRange(ip, start, end net.IP) bool {
	ipInt := ipToInt(ip)
	startInt := ipToInt(start)
	endInt := ipToInt(end)
	return startInt <= ipInt && ipInt <= endInt
}
func ipToInt(ip net.IP) uint32 {
	return uint32(ip[12])<<24 | uint32(ip[13])<<16 | uint32(ip[14])<<8 | uint32(ip[15])
}

func IsIPv4(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		// 无法解析为IP地址
		return false
	}
	return ip.To4() != nil
}

// func IsIPInCIDR(ip, cidr string) (bool, error) {
// 	// 解析CIDR
// 	_, ipNet, err := net.ParseCIDR(cidr)
// 	if err != nil {
// 		return false, fmt.Errorf("net.ParseCIDR: %w", err)
// 	}
// 	// 将IP地址字符串转换为net.IP类型
// 	ipObj := net.ParseIP(ip)
// 	// 检查IP是否在CIDR区间内
// 	return ipNet.Contains(ipObj), nil
// }

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

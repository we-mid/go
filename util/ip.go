package util

import (
	"net"
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

func ExtractAddrIP(addr string) string {
	// 尝试将remoteAddr解析为TCPAddr，如果成功则直接获取IP
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err == nil {
		return tcpAddr.IP.String()
	}
	return addr // is IP
}

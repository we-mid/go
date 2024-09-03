package ip2r

import (
	"log"
	"net"
	"strings"
)

// isIPInRange 检查给定的 IP 地址是否在指定的起始 IP 和结束 IP 之间
func isIPInRange(ipStr, startStr, endStr string) bool {
	ipStr = strings.Trim(ipStr, " \t\n\r")
	startStr = strings.Trim(startStr, " \t\n\r")
	endStr = strings.Trim(endStr, " \t\n\r")
	ip := net.ParseIP(ipStr)
	start := net.ParseIP(startStr)
	end := net.ParseIP(endStr)
	if ip == nil || start == nil || end == nil {
		log.Printf("[ip2r] isIPInRange: invalid ip=%q, start=%q, end=%q\n", ipStr, startStr, endStr)
		return false
	}
	return isNetIPInRange(ip, start, end)
}
func isNetIPInRange(ip, start, end net.IP) bool {
	ipInt := ipToInt(ip)
	startInt := ipToInt(start)
	endInt := ipToInt(end)
	return startInt <= ipInt && ipInt <= endInt
}
func ipToInt(ip net.IP) uint32 {
	if ip == nil || len(ip) < 16 { // avoid out of range
		log.Printf("[ip2r] ipToInt: invalid net.IP=%q\n", ip)
		return 0
	}
	return uint32(ip[12])<<24 | uint32(ip[13])<<16 | uint32(ip[14])<<8 | uint32(ip[15])
}

func isIPv4(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		// 无法解析为IP地址
		return false
	}
	return ip.To4() != nil
}

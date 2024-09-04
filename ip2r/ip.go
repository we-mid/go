package ip2r

import (
	"log"
	"net"
	"strings"
)

func ipToInt(ipStr string) uint32 {
	ipStr = strings.Trim(ipStr, " \t\n\r")
	ip := net.ParseIP(ipStr)
	if ip == nil {
		log.Printf("[ip2r] ipToInt: invalid ipStr=%q\n", ipStr)
	}
	return netIPToInt(ip)
}
func netIPToInt(ip net.IP) uint32 {
	if ip == nil || len(ip) < 16 { // avoid out of range
		log.Printf("[ip2r] netIPToInt: invalid net.IP=%q\n", ip)
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

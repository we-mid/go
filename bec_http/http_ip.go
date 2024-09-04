package bec_http

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"gitee.com/we-mid/go/util"
)

var (
	isGatewayTrusted = os.Getenv("GW_TRUSTED") != "0"
)

func FormatIPList(r *http.Request) string {
	ip, ips := r.RemoteAddr, GetHeaderIpsUnsafe(r)
	if ip == ips || ips == "" {
		return ip
	}
	if isGatewayTrusted {
		return ips // Will hide the ip of gateway itself
	}
	// Braces in `(ip,ip)-ip` represents that they might be untrusted and unreliable.
	return fmt.Sprintf("(%s)-%s", ips, ip)
}

// Correct way of getting Client's IP Addresses from http.Request
// https://stackoverflow.com/questions/27234861/correct-way-of-getting-clients-ip-addresses-from-http-request
func GetClientAddr(r *http.Request) string {
	if isGatewayTrusted {
		ipsStr := GetHeaderIpsUnsafe(r)
		ips := strings.Split(ipsStr, ",")
		if len(ips) > 0 {
			// 网关追加X-Forwarded-For请求头的方式是向右追加，并且最左边是最原始客户端的IP地址。
			// seg := strings.TrimSpace(ips[0])
			seg := strings.TrimSpace(ips[len(ips)-1]) // 注意这里取的是最接近直接请求端而非最接近真实用户
			if seg != "" {
				return seg
			}
		}
	}
	return r.RemoteAddr
}
func GetClientIP(r *http.Request) string {
	return util.ExtractIP(GetClientAddr(r))
}

func GetHeaderIpsUnsafe(r *http.Request) string {
	ips := r.Header.Get("X-Forwarded-For")
	if ips == "" {
		ips = r.Header.Get("X-Real-Ip")
	}
	return ips
}

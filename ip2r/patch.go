package ip2r

// todo fixme
func patch(ip string) string {
	if isIPInRange(ip, "106.55.88.0", "106.55.255.255") {
		return "中国广东省广州市 腾讯云 数据中心" // fix
	} else if isIPInRange(ip, "42.192.0.0", "42.192.255.255") || // improve
		isIPInRange(ip, "43.142.0.0", "43.143.127.255") ||
		isIPInRange(ip, "49.234.0.0", "49.235.255.255") ||
		isIPInRange(ip, "101.34.0.0", "101.35.255.255") ||
		isIPInRange(ip, "101.43.0.0", "101.43.127.255") ||
		isIPInRange(ip, "110.42.128.0", "110.42.255.255") ||
		isIPInRange(ip, "118.89.66.0", "118.89.209.255") ||
		isIPInRange(ip, "122.51.0.0", "122.51.255.255") ||
		isIPInRange(ip, "123.206.97.0", "123.206.233.255") ||
		isIPInRange(ip, "124.220.0.0", "124.223.255.255") {
		return "中国上海市 腾讯云 数据中心" // fix
	} else if isIPInRange(ip, "111.55.13.0", "111.55.15.255") {
		return "中国广西壮族自治区 移动 移动网络" // fix
	} else if isIPInRange(ip, "27.38.251.0", "27.38.251.255") ||
		isIPInRange(ip, "163.125.210.0", "163.125.210.255") {
		return "中国广东省深圳市光明区 联通 城域网" // improve
	} else if isIPInRange(ip, "123.139.60.0", "123.139.61.255") {
		return "中国陕西省西安市 联通 移动网络" // improve
	} else if isIPInRange(ip, "39.144.69.0", "39.144.69.255") {
		return "中国海南省 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.87.64", "39.144.87.127") {
		return "中国河北省张家口市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.92.0", "39.144.92.255") {
		return "中国内蒙古自治区 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.124.0", "39.144.124.255") {
		return "中国浙江省杭州市 移动 城域网" // fix
	} else if isIPInRange(ip, "39.144.142.0", "39.144.142.127") {
		return "中国四川省宜宾市 移动 城域网" // fix
	} else if isIPInRange(ip, "39.146.224.0", "39.146.224.255") {
		return "中国安徽省合肥市 移动 城域网" // fix
	} else if isIPInRange(ip, "124.62.32.0", "124.62.63.255") {
		return "韩国首尔特别市 Uplus" // improve
	}
	return "" // no patch
}

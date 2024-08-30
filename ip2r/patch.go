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
		isIPInRange(ip, "110.40.128.0", "110.40.255.255") ||
		isIPInRange(ip, "110.42.128.0", "110.42.255.255") ||
		isIPInRange(ip, "111.229.0.0", "111.229.255.255") ||
		isIPInRange(ip, "111.231.0.0", "111.231.187.255") ||
		isIPInRange(ip, "115.159.0.0", "115.159.227.255") ||
		isIPInRange(ip, "118.25.0.0", "118.25.199.255") ||
		isIPInRange(ip, "118.89.66.0", "118.89.209.255") ||
		isIPInRange(ip, "122.51.0.0", "122.51.255.255") ||
		isIPInRange(ip, "123.206.97.0", "123.206.233.255") ||
		isIPInRange(ip, "123.207.176.0", "123.207.207.255") ||
		isIPInRange(ip, "124.220.0.0", "124.223.255.255") ||
		isIPInRange(ip, "150.158.0.0", "150.158.255.255") ||
		isIPInRange(ip, "182.254.128.0", "182.254.159.255") {
		return "中国上海市 腾讯云 数据中心" // fix
	} else if isIPInRange(ip, "27.38.251.0", "27.38.251.255") ||
		isIPInRange(ip, "163.125.210.0", "163.125.210.255") {
		return "中国广东省深圳市光明区 联通 城域网" // improve
	} else if isIPInRange(ip, "111.55.10.0", "111.55.10.255") {
		return "中国广西壮族自治区柳州市 移动 移动网络" // fix
	} else if isIPInRange(ip, "111.55.13.0", "111.55.15.255") {
		return "中国广西壮族自治区 移动 移动网络" // fix
	} else if isIPInRange(ip, "111.55.146.0", "111.55.146.255") {
		return "中国四川省 移动 城域网" // fix
	} else if isIPInRange(ip, "123.139.60.0", "123.139.61.255") {
		return "中国陕西省西安市 联通 移动网络" // improve
	} else if isIPInRange(ip, "39.144.29.0", "39.144.29.255") {
		return "中国河南省郑州市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.69.0", "39.144.69.255") {
		return "中国海南省 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.87.64", "39.144.87.127") {
		return "中国河北省张家口市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.92.0", "39.144.92.255") {
		return "中国内蒙古自治区 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.95.0", "39.144.95.255") {
		return "中国山西省太原市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.124.0", "39.144.124.255") {
		return "中国浙江省杭州市 移动 城域网" // fix
	} else if isIPInRange(ip, "39.144.133.0", "39.144.133.255") {
		return "中国广西壮族自治区 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.142.0", "39.144.142.127") {
		return "中国四川省宜宾市 移动 城域网" // fix
	} else if isIPInRange(ip, "39.144.146.0", "39.144.146.63") {
		return "中国云南省曲靖市 移动 城域网" // fix
	} else if isIPInRange(ip, "39.144.156.128", "39.144.156.255") {
		return "中国江苏省苏州市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.180.64", "39.144.180.127") {
		return "中国河南省许昌市 移动 移动网络" // improve
	} else if isIPInRange(ip, "39.144.187.0", "39.144.187.127") {
		return "中国河南省洛阳市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.251.0", "39.144.251.255") {
		return "中国福建省厦门市 移动 城域网" // fix
	} else if isIPInRange(ip, "39.146.81.0", "39.146.85.255") {
		return "中国安徽省滁州市定远县 移动" // fix
	} else if isIPInRange(ip, "39.146.224.0", "39.146.224.255") {
		return "中国安徽省合肥市 移动 城域网" // fix
	} else if isIPInRange(ip, "120.202.239.0", "120.202.239.255") {
		return "中国湖北省武汉市江汉区 移动" // fix
	} else if isIPInRange(ip, "112.224.164.0", "112.224.168.255") {
		return "中国山东省青岛市 联通 移动网络" // fix
	} else if isIPInRange(ip, "117.155.88.0", "117.155.89.255") {
		return "中国湖北省孝感市汉川市 移动 城域网" // improve
	} else if isIPInRange(ip, "124.62.32.0", "124.62.63.255") {
		return "韩国首尔特别市 Uplus" // improve
	} else if isIPInRange(ip, "120.225.74.0", "120.225.74.255") {
		return "中国山东省泰安市泰山区 移动" // fix
	} else if isIPInRange(ip, "120.244.44.0", "120.244.47.255") {
		return "中国北京市顺义区 移动" // fix
	} else if isIPInRange(ip, "202.85.208.0", "202.85.208.255") {
		return "中国北京市海淀区 电信&联通 数据中心" // improve
	} else if isIPInRange(ip, "220.197.4.0", "220.197.5.255") ||
		isIPInRange(ip, "220.197.13.0", "220.197.31.255") {
		return "中国贵州省贵阳市 联通 移动网络" // improve
	} else if isIPInRange(ip, "220.197.234.0", "220.197.236.255") {
		return "中国云南省昆明市 联通 移动网络"
	} else if isIPInRange(ip, "223.104.197.0", "223.104.197.255") {
		return "中国山西省 移动 移动网络" // improve
	} else if isIPInRange(ip, "223.104.199.0", "223.104.199.255") {
		return "中国山西省太原市 移动"
	} else if isIPInRange(ip, "223.104.204.0", "223.104.204.255") {
		return "中国陕西省西安市 移动 移动网络"
	} else if isIPInRange(ip, "223.104.205.0", "223.104.205.63") {
		return "中国陕西省渭南市 移动 移动网络"
		// todo ?
		// } else if isIPInRange(ip, "223.104.205.64", "223.104.205.127") {
		// 	return "中国陕西省咸阳市 移动 移动网络"
		// } else if isIPInRange(ip, "223.104.205.128", "223.104.205.191") {
		// 	return "中国陕西省榆林市 移动 移动网络"
	} else if isIPInRange(ip, "223.104.205.192", "223.104.205.255") {
		return "中国陕西省延安市 移动 移动网络"
	} else if isIPInRange(ip, "223.160.225.0", "223.160.231.255") {
		return "中国广东省 广电 移动网络" // fix
	}
	return "" // no patch
}

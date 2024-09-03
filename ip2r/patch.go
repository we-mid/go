package ip2r

// todo fixme
func patch(ip string) string {
	if isIPInRange(ip, "106.55.88.0", "106.55.255.255") {
		return "中国广东省广州市 腾讯云 数据中心" // fix
	} else if isIPInRange(ip, "42.192.0.0", "42.192.255.255") || // improve
		isIPInRange(ip, "43.142.0.0", "43.143.127.255") ||
		isIPInRange(ip, "49.234.0.0", "49.235.255.255") ||
		isIPInRange(ip, "81.68.0.0", "81.69.255.255") ||
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
		isIPInRange(ip, "123.206.96.0", "123.206.233.255") ||
		isIPInRange(ip, "123.207.176.0", "123.207.207.255") ||
		isIPInRange(ip, "124.220.0.0", "124.223.255.255") ||
		isIPInRange(ip, "150.158.0.0", "150.158.255.255") ||
		isIPInRange(ip, "175.24.0.0", "175.24.255.255") ||
		isIPInRange(ip, "182.254.128.0", "182.254.159.255") {
		return "中国上海市 腾讯云 数据中心" // fix
	} else if isIPInRange(ip, "27.38.208.0", "27.38.208.255") ||
		isIPInRange(ip, "27.38.251.0", "27.38.251.255") ||
		isIPInRange(ip, "163.125.210.0", "163.125.210.255") {
		return "中国广东省深圳市光明区 联通 城域网" // improve
	} else if isIPInRange(ip, "111.55.10.0", "111.55.10.255") {
		return "中国广西壮族自治区柳州市 移动 移动网络" // fix
	} else if isIPInRange(ip, "111.55.13.0", "111.55.15.255") {
		return "中国广西壮族自治区 移动 移动网络" // fix
	} else if isIPInRange(ip, "111.55.146.0", "111.55.146.255") {
		return "中国四川省眉山市 移动 城域网" // fix
	} else if isIPInRange(ip, "123.139.60.0", "123.139.61.255") {
		return "中国陕西省西安市 联通 移动网络" // improve
	} else if isIPInRange(ip, "35.80.0.0", "35.111.255.255") {
		return "美国俄勒冈州波特兰 亚马逊云 数据中心" // improve
	} else if isIPInRange(ip, "36.159.223.0", "36.159.223.255") {
		return "中国安徽省芜湖市 移动 城域网" // improve
	} else if isIPInRange(ip, "39.144.26.0", "39.144.26.127") ||
		isIPInRange(ip, "39.144.29.0", "39.144.29.255") {
		return "中国河南省郑州市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.26.128", "39.144.26.255") {
		return "中国河南省商丘市 移动 移动网络" // improve
	} else if isIPInRange(ip, "39.144.61.0", "39.144.66.255") {
		return "中国广西壮族自治区 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.69.0", "39.144.69.255") {
		return "中国海南省海口市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.74.0", "39.144.74.255") {
		return "中国宁夏回族自治区银川市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.82.0", "39.144.82.255") {
		return "中国天津市 移动" // fix
	} else if isIPInRange(ip, "39.144.86.0", "39.144.86.255") {
		return "中国河北省沧州市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.87.64", "39.144.87.127") {
		return "中国河北省张家口市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.87.128", "39.144.87.191") {
		return "中国河北省承德市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.92.0", "39.144.92.255") {
		return "中国内蒙古自治区呼和浩特市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.95.0", "39.144.95.255") {
		return "中国山西省太原市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.123.0", "39.144.123.255") {
		return "中国浙江省宁波市 移动 城域网" // fix
	} else if isIPInRange(ip, "39.144.124.0", "39.144.124.255") ||
		isIPInRange(ip, "39.144.129.0", "39.144.130.255") {
		return "中国浙江省杭州市 移动 城域网" // fix
	} else if isIPInRange(ip, "39.144.133.0", "39.144.133.255") {
		// return "中国广西壮族自治区 移动 移动网络" // fix
		return "中国广西省南宁市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.140.0", "39.144.140.255") {
		// return "中国四川省广元市 移动 城域网" // fix
		return "中国四川省阿坝藏族羌族自治州 移动 城域网" // fix
	} else if isIPInRange(ip, "39.144.142.0", "39.144.142.127") {
		return "中国四川省宜宾市 移动 城域网" // fix
	} else if isIPInRange(ip, "39.144.142.0", "39.144.142.127") {
		return "中国四川省宜宾市 移动 城域网" // fix
	} else if isIPInRange(ip, "39.144.143.0", "39.144.143.255") {
		return "中国四川省绵阳市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.146.64", "39.144.146.127") {
		return "中国云南省红河哈尼族彝族自治州" // fix
	} else if isIPInRange(ip, "39.144.147.96", "39.144.147.127") {
		return "中国云南省西双版纳傣族自治州 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.156.128", "39.144.156.255") {
		return "中国江苏省苏州市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.180.64", "39.144.180.127") {
		return "中国河南省许昌市 移动 移动网络" // improve
	} else if isIPInRange(ip, "39.144.187.0", "39.144.187.127") {
		return "中国河南省洛阳市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.210.0", "39.144.210.255") {
		return "中国甘肃省兰州市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.249.0", "39.144.249.127") {
		return "中国福建省漳州市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.249.128", "39.144.249.255") {
		return "中国福建省龙岩市 移动 移动网络" // fix
	} else if isIPInRange(ip, "39.144.251.0", "39.144.251.255") {
		return "中国福建省厦门市 移动 城域网" // fix
	} else if isIPInRange(ip, "39.146.81.0", "39.146.85.255") {
		return "中国安徽省滁州市定远县 移动" // fix
	} else if isIPInRange(ip, "39.146.224.0", "39.146.224.255") {
		return "中国安徽省合肥市 移动 城域网" // fix
	} else if isIPInRange(ip, "61.186.16.0", "61.186.31.255") {
		return "中国海南省海口市秀英区 电信 城域网" // improve
	} else if isIPInRange(ip, "72.143.216.0", "72.143.239.255") {
		return "加拿大不列颠哥伦比亚温哥华 Rogers" // fix
	} else if isIPInRange(ip, "101.226.98.128", "101.227.139.255") {
		return "中国上海市 电信 数据中心"
	} else if isIPInRange(ip, "106.101.96.0", "106.101.191.255") ||
		isIPInRange(ip, "124.62.32.0", "124.62.63.255") {
		return "韩国首尔特别市 Uplus" // improve
	} else if isIPInRange(ip, "111.29.200.0", "111.29.207.255") {
		return "中国海南省定安县 移动 城域网" // fix
	} else if isIPInRange(ip, "111.55.25.0", "111.55.27.255") {
		return "中国河北省保定市 移动" // fix
	} else if isIPInRange(ip, "112.66.32.0", "112.66.47.255") {
		return "中国海南省海口市美兰区 电信 城域网" // improve
	} else if isIPInRange(ip, "112.224.164.0", "112.224.168.255") {
		return "中国山东省青岛市 联通 移动网络" // fix
	} else if isIPInRange(ip, "113.118.88.0", "113.118.91.255") ||
		isIPInRange(ip, "119.123.44.0", "119.123.51.255") {
		return "中国广东省深圳市宝安区 电信 城域网" // improve
	} else if isIPInRange(ip, "114.92.220.0", "114.92.220.255") {
		return "中国上海市浦东新区 电信" // improve
	} else if isIPInRange(ip, "116.229.180.0", "116.229.183.255") {
		return "中国上海市嘉定区 电信 城域网" // improve
	} else if isIPInRange(ip, "117.136.120.0", "117.136.120.255") {
		return "中国上海市浦东新区 移动 城域网" // improve
	} else if isIPInRange(ip, "117.155.88.0", "117.155.89.255") {
		return "中国湖北省孝感市汉川市 移动 城域网" // improve
	} else if isIPInRange(ip, "120.202.239.0", "120.202.239.255") {
		return "中国湖北省武汉市江汉区 移动" // fix
	} else if isIPInRange(ip, "120.225.74.0", "120.225.74.255") {
		return "中国山东省泰安市泰山区 移动" // fix
	} else if isIPInRange(ip, "120.229.27.0", "120.229.30.255") {
		return "中国广东省深圳市福田区 移动 城域网" // improve
	} else if isIPInRange(ip, "120.244.44.0", "120.244.47.255") {
		return "中国北京市顺义区 移动" // fix
	} else if isIPInRange(ip, "183.14.28.0", "183.14.31.255") {
		return "中国广东省深圳市南山区 电信 城域网" // improve
	} else if isIPInRange(ip, "183.254.42.0", "183.254.42.255") {
		return "中国海南省海口市美兰区 移动 城域网" // improve
	} else if isIPInRange(ip, "202.85.208.0", "202.85.208.255") {
		return "中国北京市海淀区 电信&联通 数据中心" // improve
	} else if isIPInRange(ip, "220.197.4.0", "220.197.5.255") ||
		isIPInRange(ip, "220.197.13.0", "220.197.31.255") {
		return "中国贵州省贵阳市 联通 移动网络" // improve
	} else if isIPInRange(ip, "220.197.234.0", "220.197.236.255") {
		return "中国云南省昆明市 联通 移动网络"
	} else if isIPInRange(ip, "220.200.107.0", "220.200.107.255") {
		return "中国海南省三亚市 联通" // fix
	} else if isIPInRange(ip, "223.104.24.0", "223.104.24.255") {
		return "中国贵州省毕节市 移动 移动网络" // fix
	} else if isIPInRange(ip, "223.104.197.0", "223.104.197.255") {
		return "中国山西省太原市 移动 移动网络" // improve
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
	} else if isIPInRange(ip, "223.104.227.0", "223.104.228.255") {
		return "中国天津市 移动 城域网" // fix
	} else if isIPInRange(ip, "223.122.0.0", "223.122.204.255") {
		return "中国香港 移动" // fix
	} else if isIPInRange(ip, "223.160.225.0", "223.160.231.255") {
		return "中国广东省广州市 广电 移动网络" // fix
	} else if isIPInRange(ip, "223.198.48.0", "223.198.67.255") {
		return "中国海南省文昌市 电信 城域网" // fix
	}
	return "" // no patch
}

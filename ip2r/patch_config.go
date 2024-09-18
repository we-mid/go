package ip2r

const patchConfig = `
1.53.27.0		1.53.27.255			越南同奈省 FTP
1.116.0.0		1.117.255.255		中国上海市 腾讯云 数据中心
1.202.108.0		1.202.127.255		中国北京市 电信 城域网

14.16.194.0		14.16.201.255		中国广东省东莞市 电信 城域网
14.153.245.0	14.154.26.255
14.154.27.0		14.154.27.255		中国广东省深圳市福田区 电信 城域网
14.155.51.0		14.155.52.255		中国广东省深圳市南山区 电信
14.155.192.0	14.155.195.255		中国广东省深圳市龙岗区 电信

27.38.208.0     27.38.208.255
27.38.251.0     27.38.251.255
163.125.210.0   163.125.210.255     中国广东省深圳市光明区 联通 城域网

27.38.210.0		27.38.210.255
163.125.251.0	163.125.251.255		中国广东省深圳市光明区 联通

27.115.118.0	27.115.124.255		中国上海市 联通 数据中心

27.149.20.0		27.149.59.255		中国福建省泉州市 电信 城域网

35.80.0.0		35.96.15.255
35.96.16.0		35.96.31.255
35.96.160.0		35.111.255.255		美国俄勒冈州波特兰 亚马逊云 数据中心

35.96.32.0		35.96.47.255		加拿大阿尔伯塔卡尔加里 亚马逊云 数据中心
35.96.48.0		35.96.63.255		德国黑森州法兰克福 亚马逊云 数据中心
35.96.64.0		35.96.79.255		法国法兰西岛巴黎 亚马逊云 数据中心
35.96.80.0		35.96.95.255		英国英格兰伦敦 亚马逊云 数据中心
35.96.96.0		35.96.111.255		德国黑森州法兰克福 亚马逊云 数据中心
35.96.112.0		35.96.127.255		意大利伦巴第大区米兰 亚马逊云 数据中心
35.96.128.0		35.96.143.255		澳大利亚新南威尔士悉尼 亚马逊云 数据中心
35.96.144.0		35.96.159.255		美国弗吉尼亚州阿什本 亚马逊云 数据中心

36.98.198.0		36.98.211.255		中国河北省唐山市 电信 移动网络
36.113.210.0	36.113.219.255		中国四川省成都市 电信 移动网络
36.143.135.0	36.143.135.255		中国河北省石家庄市长安区 移动
36.143.182.0	36.143.182.191		中国河北省廊坊市广阳区 移动
36.143.182.192	36.143.182.255		中国河北省廊坊市三河市 移动

36.163.138.0	36.163.138.255		中国陕西省渭南市 移动 城域网
36.163.152.0	36.163.155.255		中国陕西省西安市 移动 城域网
36.163.162.0	36.163.162.255		中国陕西省西安市临潼区 移动
36.163.165.0	36.163.165.255		中国陕西省西安市 移动 城域网

36.101.92.0		36.101.103.255
223.198.48.0    223.198.67.255      中国海南省文昌市 电信 城域网
36.101.140.0	36.101.151.255		中国海南省海口市秀英区 电信 城域网

36.150.8.0		36.150.11.255		中国江苏省盐城市 移动
36.159.223.0    36.159.223.255      中国安徽省芜湖市 移动 城域网

39.100.0.0		39.100.255.255		中国河北省张家口市 阿里云 数据中心

39.144.26.0     39.144.26.127
39.144.27.0     39.144.27.255
39.144.28.0     39.144.28.255
39.144.29.0     39.144.29.255       中国河南省郑州市 移动 移动网络
39.144.30.232	39.144.33.255		中国新疆维吾尔自治区 移动 移动网络

39.144.26.128   39.144.26.255       中国河南省商丘市 移动 移动网络
39.144.59.0		39.144.60.255		中国辽宁省沈阳市 移动 移动网络

39.144.61.0     39.144.66.255       中国广西壮族自治区 移动 移动网络
39.144.67.0		39.144.67.63		中国广西壮族自治区柳州市 移动 移动网络
39.144.67.64	39.144.67.127		中国广西壮族自治区桂林市 移动 移动网络
39.144.67.128	39.144.67.159		中国广西壮族自治区河池市 移动 移动网络
39.144.67.160	39.144.67.191		中国广西壮族自治区百色市 移动 移动网络
39.144.67.192	39.144.67.255		中国广西壮族自治区崇左市 移动 移动网络
39.144.68.0		39.144.68.255		中国广西壮族自治区玉林市 移动 城域网

39.144.69.0     39.144.69.255       中国海南省海口市 移动 移动网络

39.144.73.0		39.144.73.255
39.144.74.0     39.144.74.255
39.144.75.0		39.144.75.255		中国宁夏回族自治区银川市 移动 移动网络
39.144.82.0     39.144.82.255       中国天津市 移动
39.144.84.0		39.144.84.127		中国河北省邢台市 移动 移动网络
39.144.84.128	39.144.84.255		中国河北省邯郸市 移动 移动网络
39.144.86.0    	39.144.86.255       中国河北省沧州市 移动 移动网络
39.144.87.64    39.144.87.127       中国河北省张家口市 移动 移动网络
39.144.87.128   39.144.87.191       中国河北省承德市 移动 移动网络
39.144.92.0     39.144.92.255       中国内蒙古自治区呼和浩特市 移动 移动网络
39.144.95.0     39.144.95.255       中国山西省太原市 移动 移动网络
39.144.99.0		39.144.99.127		中国山西省临汾市 移动 移动网络
39.144.99.128	39.144.99.255		中国山西省运城市 移动 移动网络
39.144.101.0	39.144.101.255		中国吉林省 移动 移动网络
39.144.103.0	39.144.107.255		中国上海市浦东新区 移动 城域网
39.144.110.0	39.144.110.255		中国山东省临沂市 移动 城域网
39.144.123.0    39.144.123.255      中国浙江省宁波市 移动 城域网

39.144.124.0    39.144.124.255
39.144.129.0    39.144.130.255      中国浙江省杭州市 移动 城域网

# 中国广西壮族自治区 移动 移动网络
39.144.133.0    39.144.133.255      中国广西省南宁市 移动 移动网络

39.144.134.0	39.144.134.63		中国广西壮族自治区柳州市 移动 移动网络
39.144.134.64	39.144.134.127		中国广西壮族自治区桂林市 移动 移动网络
39.144.134.128	39.144.134.191		中国广西壮族自治区百色市 移动 移动网络
39.144.134.192	39.144.134.255		中国广西壮族自治区来宾市 移动 移动网络

39.144.135.0	39.144.135.63		中国广西壮族自治区玉林市 移动 移动网络
39.144.135.64	39.144.135.127		中国广西壮族自治区贵港市 移动 移动网络
39.144.135.128	39.144.135.191		中国广西壮族自治区 移动 移动网络
39.144.135.192	39.144.135.255		中国广西壮族自治区南宁市 移动 移动网络

# 中国四川省广元市 移动 城域网
39.144.140.0    39.144.140.255      中国四川省阿坝藏族羌族自治州 移动 城域网

39.144.141.0	39.144.141.63		中国四川省绵阳市 移动 移动网络
39.144.141.64	39.144.141.95		中国四川省南充市 移动 移动网络
39.144.141.96	39.144.141.127		中国四川省攀枝花市 移动 移动网络
39.144.141.128	39.144.141.159		中国四川省雅安市 移动 移动网络
39.144.141.160	39.144.141.191		中国四川省宜宾市 移动 移动网络
39.144.141.192	39.144.141.255		中国四川省自贡市 移动 移动网络

39.144.142.0	39.144.142.127		中国四川省宜宾市 移动 移动网络
39.144.142.128	39.144.142.255		中国四川省遂宁市 移动 移动网络
39.144.143.0    39.144.143.255      中国四川省绵阳市 移动 移动网络

39.144.146.0	39.144.146.63		中国云南省曲靖市 移动 移动网络
39.144.146.64	39.144.146.127		中国云南省红河哈尼族彝族自治州 移动 移动网络
39.144.146.128	39.144.146.159		中国云南省昭通市 移动 移动网络
39.144.146.160	39.144.146.191		中国云南省大理白族自治州 移动 移动网络
39.144.146.192	39.144.146.223		中国云南省文山壮族苗族自治州 移动 移动网络
39.144.146.224	39.144.146.255		中国云南省普洱市 移动 移动网络

39.144.147.0	39.144.147.63		中国云南省玉溪市 移动 移动网络
39.144.147.64	39.144.147.95		中国云南省保山市 移动 移动网络
39.144.147.96	39.144.147.127		中国云南省西双版纳傣族自治州 移动 移动网络
39.144.147.128	39.144.147.191		中国云南省临沧市 移动 移动网络
39.144.147.192	39.144.147.255		中国云南省楚雄彝族自治州 移动 移动网络

39.144.148.0	39.144.148.63		中国云南省德宏傣族景颇族自治州 移动 移动网络
39.144.148.64	39.144.148.127		中国云南省丽江市 移动 移动网络
39.144.148.128	39.144.148.191		中国云南省昭通市 移动 移动网络
39.144.148.192	39.144.150.255		中国云南省 移动 移动网络

39.144.155.0	39.144.155.127		中国江苏省南通市 移动 移动网络
39.144.155.128	39.144.155.255		中国江苏省镇江市 移动 移动网络
39.144.156.128  39.144.156.255      中国江苏省苏州市 移动 移动网络
39.144.157.0	39.144.157.127		中国江苏省无锡市 移动 移动网络
39.144.157.128	39.144.157.191		中国江苏省淮安市 移动 移动网络
39.144.157.192	39.144.157.255		中国江苏省镇江市 移动 移动网络

39.144.161.0	39.144.161.255		中国安徽省 移动 移动网络
39.144.168.0	39.144.168.63		中国江西省南昌市 移动 移动网络
39.144.168.64	39.144.168.127		中国江西省九江市 移动 移动网络
39.144.168.128	39.144.168.191		中国江西省上饶市 移动 移动网络
39.144.168.192	39.144.168.255		中国江西省抚州市 移动 移动网络

39.144.180.64   39.144.180.127      中国河南省许昌市 移动 移动网络
39.144.187.0    39.144.187.127      中国河南省洛阳市 移动 移动网络

39.144.191.0	39.144.191.63		中国湖南省湘潭市 移动 移动网络
39.144.191.64	39.144.191.127		中国湖南省株洲市 移动 移动网络
39.144.191.128	39.144.191.191		中国湖南省衡阳市 移动 移动网络
39.144.191.192	39.144.192.63		中国湖南省 移动 移动网络

39.144.196.0	39.144.196.255
39.144.200.0	39.144.209.255		中国新疆维吾尔自治区 移动 移动网络
39.144.210.0    39.144.210.255      中国甘肃省兰州市 移动 移动网络

39.144.227.0	39.144.227.63		中国贵州省铜仁市 移动 移动网络
39.144.227.64	39.144.227.127		中国贵州省遵义市 移动 移动网络
39.144.227.128	39.144.227.255		中国贵州省黔南布依族苗族自治州 移动 移动网络
39.144.228.0	39.144.228.63		中国贵州省黔东南苗族侗族自治州 移动 移动网络
39.144.228.64	39.144.228.191		中国贵州省毕节市 移动 移动网络
39.144.228.192	39.144.228.255		中国贵州省贵阳市 移动 移动网络
39.144.230.0	39.144.230.255		中国贵州省黔东南苗族侗族自治州 移动 移动网络

39.144.248.0	39.144.248.127		中国福建省宁德市 移动 移动网络
39.144.248.128	39.144.248.255		中国福建省莆田市 移动 移动网络
39.144.249.0    39.144.249.127      中国福建省漳州市 移动 移动网络
39.144.249.128  39.144.249.255      中国福建省龙岩市 移动 移动网络
39.144.250.0	39.144.250.127		中国福建省三明市 移动 移动网络
39.144.250.128	39.144.250.255		中国福建省南平市 移动 移动网络
39.144.251.0    39.144.251.255      中国福建省厦门市 移动 城域网
39.144.252.0	39.144.252.255		中国福建省泉州市 移动 城域网

39.146.32.0		39.146.35.255		中国安徽省淮南市寿县 移动 城域网
39.146.81.0     39.146.85.255       中国安徽省滁州市定远县 移动
39.146.224.0    39.146.224.255      中国安徽省合肥市 移动 城域网

61.141.248.0	61.141.255.255
61.144.172.0	61.144.175.255		中国广东省深圳市南山区 电信 城域网
61.149.164.0	61.149.167.255		中国北京市西城区 联通 城域网
61.186.16.0     61.186.31.255       中国海南省海口市秀英区 电信 城域网
72.143.216.0    72.143.239.255      加拿大不列颠哥伦比亚温哥华 Rogers

42.192.0.0      42.192.255.255
43.142.0.0      43.143.127.255
49.234.0.0      49.235.255.255
81.68.0.0       81.69.255.255
101.34.0.0      101.35.255.255
101.43.0.0      101.43.127.255
110.40.128.0    110.40.255.255
110.42.128.0    110.42.255.255
111.229.0.0     111.229.255.255
111.231.0.0     111.231.187.255
115.159.0.0     115.159.227.255
118.25.0.0      118.25.199.255
118.89.66.0     118.89.209.255
122.51.0.0      122.51.255.255
123.206.96.0    123.206.233.255
123.207.176.0   123.207.207.255
124.220.0.0     124.223.255.255
150.158.0.0     150.158.255.255
175.24.0.0      175.24.255.255
182.254.128.0   182.254.159.255     中国上海市 腾讯云 数据中心

43.224.44.0		43.224.45.255		中国北京市朝阳区
58.37.174.0		58.37.175.255		中国上海市普陀区 电信
59.32.232.0		59.32.232.255		中国广东省韶关市南雄市 电信 城域网
59.37.125.0		59.37.125.255		中国广东省深圳市南山区 电信 城域网

101.82.77.0		101.82.91.255		中国上海市松江区 电信 城域网
101.226.98.128  101.227.139.255     中国上海市 电信 数据中心
106.55.88.0     106.55.255.255      中国广东省广州市 腾讯云 数据中心

106.101.0.0		106.101.0.255		韩国京畿道 Uplus

106.101.96.0    106.101.191.255
124.62.32.0     124.62.63.255       韩国首尔特别市 Uplus

111.18.244.0	111.18.248.255		中国陕西省西安市 移动 城域网
111.29.200.0    111.29.207.255      中国海南省定安县 移动 城域网
111.43.58.0		111.43.58.255		中国黑龙江省绥化市明水县 移动 城域网
111.49.224.0	111.49.231.255		中国宁夏回族自治区吴忠市利通区 移动 城域网

111.55.10.0     111.55.10.255       中国广西壮族自治区柳州市 移动 移动网络
111.55.13.0     111.55.15.255       中国广西壮族自治区 移动 移动网络
111.55.25.0     111.55.27.255       中国河北省保定市 移动
111.55.36.0		111.55.36.255		中国云南省昆明市 移动
111.55.112.0	111.55.114.255		中国黑龙江省哈尔滨市 移动 移动网络
111.55.136.0	111.55.141.255		中国浙江省 移动 移动网络
111.55.145.0    111.55.145.255      中国四川省成都市 移动 移动网络
111.55.146.0    111.55.146.255      中国四川省眉山市 移动 城域网
111.55.149.0	111.55.150.255		中国四川省达州市 移动 城域网
111.55.166.0	111.55.166.255		中国广西壮族自治区南宁市 移动 移动网络
111.55.168.0	111.55.168.255		中国广西壮族自治区崇左市 移动
111.183.0.0		111.183.10.255
111.183.80.0	111.183.94.255		中国湖北省武汉市 电信 移动网络

112.10.202.0	112.10.203.255		中国浙江省杭州市西湖区 移动 城域网
112.66.32.0     112.66.47.255       中国海南省海口市美兰区 电信 城域网
112.96.58.0		112.96.59.255		中国广东省中山市 联通

112.224.67.0	112.224.67.255
112.224.164.0   112.224.168.255     中国山东省青岛市 联通 移动网络

113.88.208.0	113.88.211.255
113.116.160.0	113.116.163.255		中国广东省深圳市光明区 电信 城域网

113.118.88.0    113.118.91.255
113.118.112.0	113.118.115.255
113.248.1.0		113.248.3.255		中国重庆市渝北区 电信

113.201.75.0	113.201.75.255		中国陕西省西安市 联通 移动网络
116.169.80.0	116.169.80.255		中国四川省成都市 联通 移动网络
116.171.4.0		116.171.4.255		中国贵州省遵义市 联通 移动网络

114.92.220.0    114.92.220.255      中国上海市浦东新区 电信
115.44.252.0	115.44.255.255		中国广东省深圳市宝安区 天威视讯 城域网
116.25.236.0	116.25.243.255		中国广东省深圳市罗湖区 电信 城域网
116.229.180.0   116.229.183.255     中国上海市嘉定区 电信 城域网

117.136.2.0		117.136.2.127		中国河北省石家庄市 移动 移动网络
117.136.2.128	117.136.2.191		中国河北省保定市 移动 移动网络
117.136.2.192	117.136.2.255		中国河北省承德市 移动 移动网络

117.136.45.0	117.136.45.31		中国江苏省南京市 移动 移动网络
117.136.45.32	117.136.45.47		中国江苏省苏州市 移动 移动网络
117.136.45.48	117.136.45.159		中国江苏省南京市 移动 移动网络
117.136.45.160	117.136.45.175		中国江苏省扬州市 移动 移动网络
117.136.45.176	117.136.45.191		中国江苏省 移动 移动网络
117.136.45.192	117.136.45.207		中国江苏省淮安市 移动 移动网络
117.136.45.208	117.136.45.223		中国江苏省 移动 移动网络
117.136.45.224	117.136.45.239		中国江苏省连云港市 移动 移动网络
117.136.45.240	117.136.45.255		中国江苏省宿迁市 移动 移动网络

117.136.113.0	117.136.113.255		中国浙江省金华市 移动 城域网
117.136.120.0   117.136.120.255     中国上海市浦东新区 移动 城域网
117.155.88.0    117.155.89.255      中国湖北省孝感市汉川市 移动 城域网

119.123.44.0    119.123.51.255      中国广东省深圳市宝安区 电信 城域网
119.147.10.0	119.147.10.255		中国广东省深圳市南山区 电信 数据中心

120.202.239.0   120.202.239.255     中国湖北省武汉市江汉区 移动
120.225.24.0	120.225.24.255		中国山东省济南市 移动 城域网
120.225.74.0    120.225.74.255      中国山东省泰安市泰山区 移动
120.229.10.128	120.229.11.127		中国广东省深圳市坪山区 移动
120.229.11.128	120.229.11.255		中国广东省深圳市南山区 移动
120.229.27.0    120.229.30.255      中国广东省深圳市福田区 移动 城域网
120.229.86.0	120.229.86.255		中国广东省深圳市罗湖区 移动 城域网
120.231.210.0	120.231.210.255
120.231.214.0	120.231.214.255		中国广东省深圳市宝安区 移动 城域网
120.231.235.0	120.231.236.255		中国广东省深圳市龙岗区 移动 城域网
120.232.54.0	120.232.62.255		中国广东省深圳市 移动

124.240.64.0	124.240.127.255		中国广东省佛山市 广电
120.244.44.0    120.244.47.255      中国北京市顺义区 移动
120.245.126.0	120.245.126.255		中国北京市大兴区 移动
121.58.52.0		121.58.52.255		中国海南省海口市秀英 电信 移动网络

123.6.0.0		123.6.255.255		中国河南省郑州市 联通 数据中心

123.139.60.0    123.139.61.255      中国陕西省西安市 联通 移动网络
123.144.119.0	123.144.119.255
123.146.112.0	123.146.115.255		中国重庆市江北区 联通
125.85.120.0	125.85.120.255		中国重庆市渝北区 电信
125.86.88.0		125.86.103.255		中国重庆市九龙坡区 电信 城域网

139.162.0.0		139.162.63.255		新加坡 Linode 数据中心

139.227.8.0		139.227.8.255		中国上海市浦东新区

171.42.128.0	171.42.159.255		中国湖北省天门市 电信 城域网

183.14.28.0     183.14.31.255       中国广东省深圳市南山区 电信 城域网
183.17.48.0		183.17.63.255		中国广东省深圳市宝安区 电信 城域网
183.156.76.0    183.156.78.255      中国浙江省杭州市西湖区 电信 城域网
183.226.176.0	183.226.179.255		中国重庆市石柱土家族自治县 移动 城域网

183.227.218.0	183.227.218.255
183.227.219.0	183.227.219.255		中国重庆市江北区 移动
183.238.113.0	183.238.113.255		中国广东省深圳市宝安区 移动
183.238.216.0	183.238.216.255		中国广东省深圳市龙华区 移动
183.254.42.0    183.254.42.255      中国海南省海口市美兰区 移动 城域网

202.85.208.0    202.85.208.255      中国北京市海淀区 电信&联通 数据中心
203.6.224.0		203.6.239.255		中国贵州省贵阳市 电信
218.203.172.0	218.203.174.255		中国甘肃省白银市 移动

220.197.4.0     220.197.5.255
220.197.13.0    220.197.31.255      中国贵州省贵阳市 联通 移动网络

220.197.224.0	220.197.232.255
220.197.234.0   220.197.236.255     中国云南省昆明市 联通 移动网络

220.200.107.0   220.200.107.255     中国海南省三亚市 联通
220.205.249.0	220.205.255.255		中国安徽省合肥市 联通 移动网络
223.73.206.0	223.73.206.127		中国广东省深圳市罗湖区 移动 城域网
223.73.206.128	223.73.206.255		中国广东省深圳市龙岗区 移动 城域网
223.73.209.0	223.73.209.255		中国广东省深圳市罗湖区 移动 城域网
223.104.24.0    223.104.24.255      中国贵州省毕节市 移动 移动网络
223.104.38.0	223.104.42.255		中国北京市 移动 城域网
223.104.49.0	223.104.50.255		中国福建省漳州市 移动 移动网络
223.104.53.0	223.104.53.255		中国福建省泉州市 移动 城域网
223.104.68.0	223.104.68.255		中国广东省深圳市 移动 城域网

223.104.71.0	223.104.71.127		中国广东省汕头市 移动 移动网络
223.104.71.128	223.104.71.191		中国广东省汕尾市 移动 移动网络
223.104.71.192	223.104.71.207		中国广东省韶关市 移动 移动网络
223.104.71.208	223.104.71.239		中国广东省茂名市 移动 移动网络
223.104.71.240	223.104.71.255		中国广东省阳江市 移动 移动网络

# 范围太广 中山-珠海
223.104.85.0	223.104.85.255		中国广东省 移动 移动网络

223.104.122.0	223.104.122.255		中国湖北省武汉市 移动 移动网络

# 范围太广 吕梁-太原
223.104.197.0   223.104.197.255     中国山西省 移动 移动网络

223.104.199.0   223.104.199.255     中国山西省太原市 移动

223.104.202.0	223.104.202.255
223.104.204.0   223.104.204.255     中国陕西省西安市 移动 移动网络

223.104.205.0		223.104.205.63		中国陕西省渭南市 移动 移动网络
223.104.205.64		223.104.205.127		中国陕西省咸阳市 移动 移动网络
223.104.205.128		223.104.205.191		中国陕西省榆林市 移动 移动网络
223.104.205.192		223.104.205.255		中国陕西省延安市 移动 移动网络

223.104.206.0		223.104.206.63		中国陕西省宝鸡市 移动 移动网络
223.104.206.64		223.104.206.127		中国陕西省汉中市 移动 移动网络
223.104.206.128		223.104.206.255		中国陕西省安康市 移动 移动网络

223.104.227.0       223.104.228.255     中国天津市 移动 城域网
223.122.0.0         223.122.204.255     中国香港新界 移动
223.160.148.0		223.160.151.255		中国辽宁省沈阳市 广电 移动网络
223.160.172.0		223.160.175.255		中国陕西省西安市 广电 移动网络
223.160.208.0		223.160.215.255		中国浙江省杭州市 广电 移动网络

223.160.224.0		223.160.224.255
223.160.225.0       223.160.231.255     中国广东省广州市 广电 移动网络
223.223.176.0		223.223.183.255		中国北京市海淀区 电信&联通&移动 数据中心
`

# `ip2r` IP归属地离线查询 Go语言SDK

同时集成了两个库，支持IPv4+IPv6：

- Go库（IPv4）：https://github.com/lionsoul2014/ip2region
- Node.js库（可选，+IPv6支持）：https://github.com/yourtion/node-ip2region

## 运行Node.js进程依赖（可选，+IPv6支持，内存占用145MB左右）

当前 [node_ip2r](https://gitee.com/we-mid/node_ip2r) 提供的服务交互，使用的是 Unix Socket IPC 策略实现

```sh
git clone git@gitee.com:we-mid/node_ip2r.git
cd node_ip2r
pnpm install
node server
```

## 下载ip2region.xdb，运行测试

```sh
git clone git@gitee.com:we-mid/go.git
cd go/ip2r
# 下载 ip2region.xdb 到本地
# 参考 https://github.com/lionsoul2014/ip2region
go test
```

## 在你的Go项目中使用ip2r

```go
import "gitee.com/we-mid/go/ip2r"

// 下载 ip2region.xdb 到本地
// 参考 https://github.com/lionsoul2014/ip2region
const dbPath = "./ip2region.xdb"

if err := ip2r.Load(dbPath); err != nil {
	log.Println("[ip2r] 加载失败", err)
	return
}
defer ip2r.Close()

for {
	// ...
	res, err := ip2r.Query(req.IP)
	// >> IPv4
	// res=&{IP:124.220.36.180 Region:中国上海市 腾讯云 数据中心 Took:6.777µs}
	// IPv6
	// res=&{IP:2408:8456:f10c:a4fd:9925:5858:55aa:33af Region:中国广东省中山市 中国联通3GNET网络 Took:689.283µs}
}
```

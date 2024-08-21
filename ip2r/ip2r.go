package ip2r

import (
	"fmt"
	"log"
	"strings"
	"time"

	"gitee.com/we-mid/go/ip2r/core"
	"gitee.com/we-mid/go/ip2r/db"
	"gitee.com/we-mid/go/ip2r/node"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var searcher *xdb.Searcher

// todo IPv6 support (maybe with node-ip2region)
// var ip = "1.2.3.4" // IPv4
// var ip = "2409:8904:a730:1b4c:1d68:6d5d:e915:7d20" // IPv6 ??
func Query(ip string) (*core.Res, error) {
	var err error
	var region string
	var start = time.Now()
	ip = strings.TrimSpace(ip) // 容错
	if isIPv4(ip) {
		// patch
		// "中国上海市 腾讯云 数据中心"
		// "中国广东省深圳市光明区 联通 城域网"
		if region = patch(ip); region == "" {
			// "中国|0|广东省|深圳市|联通"
			// "中国|0|北京|北京市|联通"
			region, err = searcher.SearchByStr(ip)
			if err != nil {
				return nil, fmt.Errorf("failed to search IP(%s): %w", ip, err)
			}
		}
	} else {
		// todo node-ip2region
		// return nil, fmt.Errorf("only IPv4 is supported. IP(%s)", ip)
		if res, err := node.Query(ip); err != nil {
			return nil, err
		} else {
			region = res.Region
		}
	}
	took := time.Since(start)
	return &core.Res{IP: ip, Region: region, Took: took}, nil
}

func Close() {
	if searcher != nil {
		searcher.Close()
	}
	node.Close()
}
func Load(dbPath string) error {
	Close()
	if err := node.Setup(); err != nil {
		// return err
		log.Println("[ip2r] [可选服务] node_ip2r连接失败", err) // 非核心必选依赖
		// skip
	}
	var err error
	// https://github.com/lionsoul2014/ip2region/tree/master/binding/golang
	// 方式三、缓存整个 xdb 数据：
	// 备注：并发使用，用整个 xdb 缓存创建的 searcher 对象可以安全用于并发。
	searcher, err = db.Load03FullCache(dbPath)
	return err
}

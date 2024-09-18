package ip2r

import "testing"

func TestPatch(t *testing.T) {
	// if r := patch("220.197.20.6"); r != "中国贵州省贵阳市 联通 移动网络" {
	if r := patch("123.207.193.82"); r != "中国上海市 腾讯云 数据中心" {
		t.Fatal("got:", r)
	}
}

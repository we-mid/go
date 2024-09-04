package ip2r

import "testing"

func TestPatch(t *testing.T) {
	if r := patch("220.197.20.6"); r != "中国贵州省贵阳市 联通 移动网络" {
		t.Fatal("got:", r)
	}
}

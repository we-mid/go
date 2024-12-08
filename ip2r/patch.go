package ip2r

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// todo fixme 实现合理的树结构查找

type segment struct {
	left, right uint32
	region      string
}

var segs []segment

func patch(ip string) string {
	v := ipToInt(ip)
	for _, s := range segs {
		if v >= s.left && v <= s.right {
			return s.region
		}
	}
	return ""
}

func init() {
	CustomPatchString(patchConfig)
}

func CustomPatch(filename string) error {
	bs, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	CustomPatchString(string(bs))
	return nil
}

const reStrIPv4 = `\d{1,3}(\.\d{1,3}){3}`

var reStrIPv4CIDR = fmt.Sprintf(`%s\/\d{1,2}`, reStrIPv4)
var reMultiIPv4CIDR = regexp.MustCompile(fmt.Sprintf(`^%s(\,%s)*$`, reStrIPv4CIDR, reStrIPv4CIDR))

// todo: config string => []byte
func CustomPatchString(config string) {
	lines := strings.Split(config, "\n")
	var buf [][2]uint32
	for _, line := range lines {
		if strings.HasPrefix(line, "#") { // commenting
			continue // skip
		}
		line = strings.Trim(line, " \t\r\n")
		if line == "" { // empty
			continue // skip
		}
		parts := strings.Fields(line)
		if len(parts) < 2 { // invalid
			continue // skip
		}
		start := strings.Trim(parts[0], " \t\n\r")
		end := strings.Trim(parts[1], " \t\n\r")
		left, right := ipToInt(start), ipToInt(end)
		if len(parts) == 2 {
			buf = append(buf, [2]uint32{left, right})
		} else if len(parts) >= 3 {
			buf = append(buf, [2]uint32{left, right})

			// 扩展：支持末尾追加一个可选的CIDR标记
			if len(parts) >= 4 && reMultiIPv4CIDR.MatchString(parts[len(parts)-1]) {
				parts = parts[:len(parts)-1] // remove last
			}

			region := strings.Join(parts[2:], " ")
			for _, p := range buf {
				segs = append(segs, segment{left: p[0], right: p[1], region: region})
			}
			buf = nil // reset
		}
	}
}

package ip2r

import (
	"os"
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
			region := strings.Join(parts[2:], " ")
			for _, p := range buf {
				segs = append(segs, segment{left: p[0], right: p[1], region: region})
			}
			buf = nil // reset
		}
	}
}

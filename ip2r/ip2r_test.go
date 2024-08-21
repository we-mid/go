package ip2r

import (
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	if err := Load("./ip2region.xdb"); err != nil {
		log.Fatalln("Load:", err)
	}
	m.Run()
}

func TestIP2R(t *testing.T) {
	ips := []string{
		// IPv4
		"124.220.36.180",
		// IPv6
		"2408:8456:f10c:a4fd:9925:5858:55aa:33af",
	}
	for _, ip := range ips {
		if res, err := Query(ip); err != nil {
			t.Fatalf("err should be nil, got %v", err)
		} else {
			log.Printf("res=%+v", res)
		}
	}
}

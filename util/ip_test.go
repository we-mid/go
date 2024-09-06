package util

import (
	"testing"
)

func TestExtractIP(t *testing.T) {
	s := ExtractAddrIP("12.34.56.78")
	if s != "12.34.56.78" {
		t.Fatalf("got: %q", s)
	}
	s = ExtractAddrIP("12.34.56.78:8080")
	if s != "12.34.56.78" {
		t.Fatalf("got: %q", s)
	}
	s = ExtractAddrIP("240e:370:d52:fa11:8944:43e:676c:6fe4")
	if s != "240e:370:d52:fa11:8944:43e:676c:6fe4" {
		t.Fatalf("got: %q", s)
	}
	s = ExtractAddrIP("[240e:370:d52:fa11:8944:43e:676c:6fe4]:7070")
	if s != "240e:370:d52:fa11:8944:43e:676c:6fe4" {
		t.Fatalf("got: %q", s)
	}
}

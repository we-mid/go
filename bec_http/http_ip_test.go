package bec_http

import (
	"net/http/httptest"
	"testing"
)

func TestGetMaybeRealIP(t *testing.T) {
	r := httptest.NewRequest("GET", "/foo", nil)
	r.RemoteAddr = "3.3.3.3:8899"
	r.Header.Add("X-Forwarded-For", "1.1.1.1, 2.2.2.2")
	if s := GetMaybeRealIP(r); s != "1.1.1.1" {
		t.Fatalf("got: %q", s)
	}
}

func TestGetClientAddrAndIP(t *testing.T) {
	r := httptest.NewRequest("GET", "/foo", nil)
	r.RemoteAddr = "3.3.3.3:8899"
	r.Header.Add("X-Forwared-For", "1.1.1.1, 2.2.2.2")
	if s := GetClientAddr(r); s != "3.3.3.3:8899" {
		t.Fatalf("got: %q", s)
	}
	if s := GetClientIP(r); s != "3.3.3.3" {
		t.Fatalf("got: %q", s)
	}
}

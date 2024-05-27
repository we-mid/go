package util

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	isGatewayTrusted = true
)

func init() {
	if os.Getenv("GW_TRUSTED") == "0" {
		isGatewayTrusted = false
	}
}

func HttpLog(r *http.Request, start time.Time, status int, err error) {
	layout := "[req] %s %s %s %dms %d %s\n"
	message := ""
	if err != nil && !errors.Is(err, ErrHandledAndBreak) {
		message = err.Error()
	}
	path := r.URL.Path
	if r.URL.RawQuery != "" {
		path += "?" + r.URL.RawQuery
	}
	params := []any{
		GetClientIP(r),
		r.Method,
		path,
		time.Since(start).Milliseconds(),
		status,
		message,
	}
	fmt.Printf(layout, params...)
	if message != "" {
		log.Printf(layout, params...)
	}
}

// Correct way of getting Client's IP Addresses from http.Request
// https://stackoverflow.com/questions/27234861/correct-way-of-getting-clients-ip-addresses-from-http-request
func GetClientIP(r *http.Request) string {
	IPAddress := ""
	if isGatewayTrusted {
		IPAddress = r.Header.Get("X-Real-Ip")
		if IPAddress == "" {
			IPAddress = r.Header.Get("X-Forwarded-For")
		}
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

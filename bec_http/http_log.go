package bec_http

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	timeFormat = os.Getenv("HTTPLOG_TIMEFORMAT")
)

func init() {
	if timeFormat == "" {
		timeFormat = "01/02 15:04" // 'MM/DD HH:mm'
	}
}

func HttpLog(r *http.Request, start time.Time, status int, err error) {
	layout := "%s [req] %s %s %s %dms %d %s\n"
	message := ""
	if err != nil && !errors.Is(err, ErrHandledAndBreak) {
		message = err.Error()
	}
	path := r.URL.Path
	if r.URL.RawQuery != "" {
		path += "?" + r.URL.RawQuery
	}
	params := []any{
		time.Now().Format(timeFormat),
		FormatIPList(r),
		r.Method,
		path,
		time.Since(start).Milliseconds(),
		status,
		message,
	}
	fmt.Printf(layout, params...) // to stdout
	if message != "" {
		log.Printf(layout[3:], params[1:]...) // to stderr
	}
}

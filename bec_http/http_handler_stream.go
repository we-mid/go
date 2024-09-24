package bec_http

import (
	"context"
	"errors"
	"net/http"
	"time"

	"gitee.com/we-mid/go/util"
)

var StreamEof = errors.New("StreamEof")

func StreamHandlerWrap(contentType string, logic LogicStream) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var headerSet bool
		// 获取 flusher
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming not supported!", http.StatusInternalServerError)
			return
		}
		if err := logic(w, r, func(bytes []byte) {
			if !headerSet {
				// w.WriteHeader(http.StatusOK)
				// 设置响应类型为 text/event-stream，这样客户端会按流的方式处理数据
				// w.Header().Set("Content-Type", "text/event-stream")
				w.Header().Set("Content-Type", contentType)
				// w.Header().Set("Cache-Control", "no-cache")
				w.Header().Set("Connection", "keep-alive")
				headerSet = true
			}
			w.Write(bytes)
			flusher.Flush()
		}); err != nil {
			if errors.Is(err, ErrHandledAndBreak) {
				HttpLog(r, start, 200, err)
				// } else if errors.Is(err, StreamEof) {
			} else if util.IsErrorLike(err, StreamEof) {
				// noop
				HttpLog(r, start, 200, nil)
			} else if errors.Is(err, context.DeadlineExceeded) {
				// ignore
				w.Write([]byte("..."))
				flusher.Flush()
				HttpLog(r, start, 200, nil)
			} else {
				if e, ok := err.(IStatusError); ok && e.Status() < 500 {
					SendErrText(w, err)
					HttpLog(r, start, e.Status(), err)
				} else {
					// protect error messages for 5xx
					SendErrText(w, Err500)
					HttpLog(r, start, 500, err)
				}
			}
		} else {
			// noop
			HttpLog(r, start, 200, nil)
		}
	}
}

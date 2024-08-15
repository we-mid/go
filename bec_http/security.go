package bec_http

import "net/http"

func LimitReqBody(w http.ResponseWriter, r *http.Request, maxBytes int64) {
	// 设置内存限制（这里仅作为示例，根据需要调整）
	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
}

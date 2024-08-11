package bec_http

import (
	"net/http"
)

var corsHeaders = map[string]string{
	// 'Access-Control-Allow-Origin': '*', /* @dev First, read about security */
	"Access-Control-Allow-Methods": "POST, GET",
	"Access-Control-Allow-Headers": "Content-Type, Cookie, X-Crm-Credential",
	// Can not read content-disposition from resposne header #67
	// https://github.com/matthew-andrews/isomorphic-fetch/issues/67#issuecomment-353605695
	"Access-Control-Expose-Headers":    "Content-Disposition, Set-Cookie",
	"Access-Control-Allow-Credentials": "true",
}

// Golang CORS Guide: What It Is and How to Enable It
// https://www.stackhawk.com/blog/golang-cors-guide-what-it-is-and-how-to-enable-it/
func EnableCors(w http.ResponseWriter, r *http.Request) error {
	if len(r.Header["Origin"]) < 1 {
		return nil
	}
	w.Header().Set("Access-Control-Allow-Origin", r.Header["Origin"][0])
	for k, v := range corsHeaders {
		w.Header().Set(k, v)
	}
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return ErrHandledAndBreak
	}
	return nil
}

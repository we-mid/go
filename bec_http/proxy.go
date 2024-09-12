package bec_http

import (
	"io"
	"log"
	"net/http"
	"strings"
)

// urlPath should be like `/foo/`, mind the trailing slash
func ProxyHandler(urlPath, proxyURL string, allowPrefix []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqURI := r.RequestURI
		targetStr := strings.TrimPrefix(reqURI, urlPath)

		var found bool
		for _, p := range allowPrefix {
			if strings.HasPrefix(targetStr, p) {
				found = true
			}
		}
		if !found {
			log.Println("Failed to match allowPrefix", targetStr)
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		targetURL := "https://" + targetStr

		// 创建一个新的请求，并复制原始请求的信息
		proxyReq, err := http.NewRequest(r.Method, proxyURL+targetURL, r.Body)
		if err != nil {
			log.Println("Failed to new request", targetURL)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		proxyReq.Header = r.Header

		// 使用http.Client来发送新的请求
		client := &http.Client{}
		resp, err := client.Do(proxyReq)
		if err != nil {
			log.Println("Failed to send request", targetURL)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// 将响应头复制到原始客户端的响应
		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)

		// 将响应体复制到原始客户端
		io.Copy(w, resp.Body)
	}
}

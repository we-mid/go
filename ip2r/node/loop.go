package node

import (
	"encoding/json"
	"fmt"
	"net"
)

// 当前与 node_ip2r 服务交互，使用 Unix Socket IPC 策略实现
const sock = "/tmp/node_ip2r.unix.sock"
const chanBuffer = 5

// close reqCh to exit loop and teardown
func ipcLoop() (chan<- XReq, <-chan XXRes, error) {
	// 连接到 Unix Domain Socket 服务器
	conn, err := net.Dial("unix", sock)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to connect to socket: %w", err)
	}
	cancel := make(chan struct{}, 1)
	reqCh := make(chan XReq, chanBuffer)
	resCh := make(chan XXRes, chanBuffer)
	go func() {
		// 读取请求
		for {
			select {
			case request, ok := <-reqCh:
				if !ok { // channel is closed
					// teardown
					conn.Close()         // cleanup
					cancel <- struct{}{} // signal
					return               // exit loop
				}
				// 构建请求
				// request := core.Req{IP: "2409:891f:6864:a3ba:100e:5b26:3feb:fc26"}
				// 序列化请求并发送
				jsonRequest, err := json.Marshal(request)
				if err != nil {
					resCh <- XXRes{Error: fmt.Errorf("Failed to marshal request: %w", err)}
					continue // skip
				}
				conn.Write(jsonRequest)
			}
		}
	}()
	go func() {
		// 输出响应
		decoder := json.NewDecoder(conn)
		for {
			select {
			case <-cancel:
				return // exit loop
			default:
				var response XXRes
				if err := decoder.Decode(&response); err != nil {
					resCh <- XXRes{Error: fmt.Errorf("Failed to decode response: %w", err)}
				} else {
					// log.Printf("Received response: %+v\n", response)
					resCh <- response
				}
			}
		}
	}()
	return reqCh, resCh, nil
}

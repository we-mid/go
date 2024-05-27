package util

import (
	"bytes"
	"io"
	"net"
)

// 一次性读取服务器的响应
func ReadFromConnOnce(conn net.Conn) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(conn) // ReadFrom 会一直读取直到遇到 EOF 或错误
	if err != nil && err != io.EOF {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 读取服务器的响应
func ReadFromConn(conn net.Conn, chunkSize int) ([]byte, error) {
	var buf bytes.Buffer
	for {
		// 假设我们简单地按块读取，直到没有更多的数据
		// 在实际情况下，你可能需要一个更复杂的协议来界定消息的边界
		data := make([]byte, chunkSize) // 假设最大消息长度为1024字节
		n, err := conn.Read(data)
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}
		buf.Write(data[:n])
		// 假设服务器在发送完整消息后会关闭连接，或者我们使用其他方式来确定消息何时结束
		// 在这里，我们简单地假设服务器在发送完消息后关闭了连接
	}
	return buf.Bytes(), nil
}

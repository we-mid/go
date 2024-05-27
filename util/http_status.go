package util

import (
	"fmt"
	"net/http"
)

// StatusError 是一个自定义错误类型，包含了 HTTP 状态码
type StatusError struct {
	Status int   // HTTP 状态码
	Err    error // 底层错误
}

// Error 方法实现了 error 接口
func (e *StatusError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%d: %v", e.Status, e.Err)
	}
	if msg := http.StatusText(e.Status); msg != "" {
		return fmt.Sprintf("%d: %s", e.Status, msg)
	}
	return fmt.Sprintf("%d", e.Status)
}

// NewStatusError 创建一个新的 StatusError
func NewStatusError(status int, err error) *StatusError {
	return &StatusError{
		Status: status,
		Err:    err,
	}
}

func NewStatusErrorf(status int, format string, params ...any) *StatusError {
	return NewStatusError(status, fmt.Errorf(format, params...))
}

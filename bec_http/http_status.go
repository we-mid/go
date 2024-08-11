package bec_http

import (
	"fmt"
	"net/http"
)

// StatusError 是一个自定义错误类型，包含了 HTTP 状态码
type IStatusError interface {
	error        // 继承 error 接口
	Status() int // HTTP 状态码
	Err() error  // 底层错误
}
type StatusError struct {
	status int   // HTTP 状态码
	err    error // 底层错误
}

func (e *StatusError) Status() int {
	return e.status
}
func (e *StatusError) Err() error {
	return e.err
}

// Error 方法实现了 error 接口
func (e *StatusError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%d: %v", e.status, e.err)
	}
	if msg := http.StatusText(e.status); msg != "" {
		return fmt.Sprintf("%d: %s", e.status, msg)
	}
	return fmt.Sprintf("%d", e.status)
}

// NewStatusError 创建一个新的 StatusError
func NewStatusError(status int, err error) IStatusError {
	return &StatusError{
		status: status,
		err:    err,
	}
}
func NewStatusErrorf(status int, format string, params ...any) IStatusError {
	return NewStatusError(status, fmt.Errorf(format, params...))
}

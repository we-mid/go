package session_memory

import (
	"errors"
	"sync"
	"time"

	"gitee.com/we-mid/go/session/v2"
	"gitee.com/we-mid/go/util"
)

const (
	lenSessID = 32
	maxTries  = 3
)

// 假设的session存储结构，这里使用map和sync.Map在内存中存储
type SessionStore[T any] struct {
	session.SessionStore[T]
	sync.Map
}

func NewStore[T any]() *SessionStore[T] {
	return &SessionStore[T]{}
}

// 添加一个session
func (s *SessionStore[T]) Set(sessID string, value T, expiry time.Duration) error {
	s.Store(sessID, &session.SessionData[T]{
		Value:  value,
		Expiry: time.Now().Add(expiry),
	})
	return nil
}

// 获取一个session
func (s *SessionStore[T]) Get(sessID string) (T, bool, error) {
	var zero T
	val, ok := s.Load(sessID)
	if !ok {
		return zero, false, nil
	}
	data, ok := val.(*session.SessionData[T])
	if !ok || data.Expired() {
		s.Delete(sessID)
		return zero, false, nil
	}
	return data.Value, true, nil
}

func (s *SessionStore[T]) NewID() (string, error) {
	var err error
	for i := 0; i < maxTries; i++ {
		var exists bool
		var id string
		if id, err = util.RandomBase64(lenSessID); err != nil {
			continue
		}
		_, exists, err = s.Get(id)
		if err != nil {
			continue
		}
		if exists {
			err = errors.New("sessID is duplicated")
			continue
		}
		return id, nil
	}
	return "", err
}

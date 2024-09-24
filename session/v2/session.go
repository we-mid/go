package session

import (
	"time"
)

type SessionStore[T any] interface {
	NewID() (string, error)
	Add(sessID string, value T, expiry time.Duration) error
	Get(sessID string) (T, bool, error)
}

type SessionData[T any] struct {
	Value  T
	Expiry time.Time
}

func (s *SessionData[T]) Expired() bool {
	return time.Now().After(s.Expiry)
}

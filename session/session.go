package session

import (
	"time"
)

type SessionStore interface {
	NewID() (string, error)
	Add(sessID string, value any, expiry time.Duration) error
	Get(sessID string) (any, bool, error)
}

type SessionData struct {
	Value  any
	Expiry time.Time
}

func (s *SessionData) Expired() bool {
	return time.Now().After(s.Expiry)
}

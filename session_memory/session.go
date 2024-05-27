package session_memory

import (
	"errors"
	"sync"
	"time"

	"github.com/we-task/Todo-as-a-Service/x/session"
	"github.com/we-task/Todo-as-a-Service/x/util"
)

const (
	lenSessID = 32
	maxTries  = 3
)

// 假设的session存储结构，这里使用map和sync.Map在内存中存储
type SessionStore struct {
	session.SessionStore
	sync.Map
}

func NewStore() *SessionStore {
	return &SessionStore{}
}

// 添加一个session
func (s *SessionStore) Add(sessID string, value any, expiry time.Duration) error {
	s.Store(sessID, &session.SessionData{
		Value:  value,
		Expiry: time.Now().Add(expiry),
	})
	return nil
}

// 获取一个session
func (s *SessionStore) Get(sessID string) (any, bool, error) {
	val, ok := s.Load(sessID)
	if !ok {
		return nil, false, nil
	}
	data, ok := val.(*session.SessionData)
	if !ok || data.Expired() {
		s.Delete(sessID)
		return nil, false, nil
	}
	return data.Value, true, nil
}

func (s *SessionStore) NewID() (string, error) {
	var err error
	for i := 0; i < maxTries; i++ {
		var exists bool
		var id string
		if id, err = util.RandomString(lenSessID); err != nil {
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

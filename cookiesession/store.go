package cookiesession

import (
	"net/http"
	"time"

	session_memory "gitee.com/we-mid/go/session_memory/v2"
)

type Options struct {
	CookieName   string
	CookiePath   string
	CookieSecure bool
	TTLSession   time.Duration
}

type Store[T any] struct {
	Options
	s *session_memory.SessionStore[T]
}

func NewStore[T any](options Options) *Store[T] {
	s := session_memory.NewStore[T]()
	p := &Store[T]{options, s}
	return p
}

func (p *Store[T]) Get(r *http.Request) (T, bool, error) {
	var zero T
	cookie, err := r.Cookie(p.CookieName)
	if err != nil {
		return zero, false, err
	}
	return p.s.Get(cookie.Value)
}

func (p *Store[T]) Set(w http.ResponseWriter, value T) error {
	sessID, err := p.s.NewID()
	if err != nil {
		return err
	}
	if err := p.s.Set(sessID, value, p.TTLSession); err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		Secure:   p.CookieSecure,
		Path:     p.CookiePath,
		Name:     p.CookieName,
		Value:    sessID,
		MaxAge:   int(p.TTLSession.Seconds()),
	})
	return nil
}

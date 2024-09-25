package cookiesession

import (
	"net/http"
	"time"

	session "gitee.com/we-mid/go/session/v2"
)

type Options[T any] struct {
	SessionStore session.SessionStore[T]
	CookieName   string
	CookiePath   string
	CookieSecure bool
	TTLSession   time.Duration
}

type Store[T any] struct {
	Options[T]
}

func NewStore[T any](options Options[T]) *Store[T] {
	p := &Store[T]{options}
	return p
}

func (p *Store[T]) GetFrom(r *http.Request) (T, bool, error) {
	var zero T
	cookie, err := r.Cookie(p.CookieName)
	if err != nil {
		return zero, false, err
	}
	return p.SessionStore.Get(cookie.Value)
}

func (p *Store[T]) SetTo(w http.ResponseWriter, value T) error {
	sessID, err := p.SessionStore.NewID()
	if err != nil {
		return err
	}
	if err := p.SessionStore.Set(sessID, value, p.TTLSession); err != nil {
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

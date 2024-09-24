package passwordless

import (
	"net/http"

	"gitee.com/we-mid/go/bec_http"
)

type sessionRes struct {
	Email string `json:"email"`
}

func (p *Passwordless) HandleSession(w http.ResponseWriter, r *http.Request) (any, error) {
	var req struct{}
	if err := p.handleReq(w, r, &req); err != nil {
		return nil, err
	}
	email, err := p.GetEmail(r)
	if err != nil {
		return nil, err
	}
	if email == "" {
		return nil, bec_http.Err401
	}
	return &sessionRes{email}, nil
}

func (p *Passwordless) GetEmail(r *http.Request) (string, error) {
	// for type string, ok can be ignored
	email, _, err := p.csStore.Get(r)
	if err != nil && err == http.ErrNoCookie {
		err = nil // no cookie, ignore
	}
	return email, err
}

func (p *Passwordless) bindEmail(w http.ResponseWriter, email string) error {
	return p.csStore.Set(w, email)
}

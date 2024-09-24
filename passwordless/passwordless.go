package passwordless

import (
	"sync"
	"time"

	"gitee.com/we-mid/go/cookiesession"
)

type codeBinding struct {
	code     string
	expireAt time.Time
}

// todo: use chaining option pattern
type Options struct {
	CookieSession cookiesession.Options[string]
	EnableCORS    bool
	RoutePrefix   string
	SaaSURL       string
	SaaSName      string
	// MailSubject   string
	// MailTemplate  string
	LenCode   int
	TTLCode   time.Duration
	OnAttempt func(string)
	OnVerify  func(string, bool)
}

type Passwordless struct {
	Options
	mu      *sync.Mutex
	codeMap map[string]codeBinding
	csStore *cookiesession.Store[string]
}

func New(options Options) *Passwordless {
	var mu sync.Mutex
	codeMap := make(map[string]codeBinding)
	csStore := cookiesession.NewStore[string](options.CookieSession)
	p := &Passwordless{options, &mu, codeMap, csStore}
	go p.cleanExpired()
	return p
}

func (p *Passwordless) cleanExpired() {
	ticker := time.NewTicker(cleanInterval)
	defer ticker.Stop()
	for range ticker.C {
		p.mu.Lock()
		now := time.Now()
		for k, b := range p.codeMap {
			if now.Compare(b.expireAt) >= 0 {
				delete(p.codeMap, k)
			}
		}
		p.mu.Unlock()
	}
}

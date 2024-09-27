package passwordless

import (
	"encoding/json"
	"log"
	"os"
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
	LenCode         int
	TTLCode         time.Duration
	OnAttempt       func(string)
	OnVerify        func(string, bool)
	TestUsersEnvKey string
}

type Passwordless struct {
	Options
	mu        *sync.Mutex
	codeMap   map[string]codeBinding
	csStore   *cookiesession.Store[string]
	testUsers map[string]string
}

func New(options Options) *Passwordless {
	var mu sync.Mutex
	codeMap := make(map[string]codeBinding)
	csStore := cookiesession.NewStore[string](options.CookieSession)

	testUsers := make(map[string]string)
	envKey := options.TestUsersEnvKey
	if envKey != "" {
		str := os.Getenv(envKey)
		if err := json.Unmarshal([]byte(str), &testUsers); err != nil {
			log.Printf("[passwordless] error json.Unmarshal %q: %v\n", envKey, err)
		}
	}
	p := &Passwordless{options, &mu, codeMap, csStore, testUsers}
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
			// if now.Compare(b.expireAt) >= 0 {
			if !b.expireAt.After(now) {
				delete(p.codeMap, k)
			}
		}
		p.mu.Unlock()
	}
}

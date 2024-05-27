package bec_gin

import (
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const (
	rpsLoginApi = 1
)

var (
	// fix: [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
	// Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
	trustedProxies  = []string{"127.0.0.1", "::1"} // default
	_trustedProxies = os.Getenv("TRUSTED_PROXIES")

	cookiePath   = os.Getenv("COOKIE_PATH")
	cookieSecure = os.Getenv("COOKIE_SECURE") == "1"
	cookieSecret = os.Getenv("COOKIE_SECRET")
)

func init() {
	assertBadCookieSecret()

	if _trustedProxies != "" {
		trustedProxies = strings.Split(_trustedProxies, ",")
	}
}

func RatelimitAllReq(rpsAllReq int) gin.HandlerFunc {
	return leakyBucket(rpsAllReq)
}
func ratelimitLoginApi() gin.HandlerFunc {
	return leakyBucket(rpsLoginApi)
}

func SecurityLog(c *gin.Context, layout string, params ...any) {
	if c != nil {
		prefix := "IP=" + c.ClientIP() + " | "
		params = append([]any{prefix}, params...)
	}
	log.Printf("[security] %s"+layout+"\n", params...)
}

func assertBadCookieSecret() {
	ln := 30
	if len(cookieSecret) < ln {
		log.Fatalf("Length of env.COOKIE_SECRET should be gte %d\n", ln)
	}
}
func EnhancedCookieSessions(cookieName string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(cookieSecret))
	// mind security: cookie options
	// https://github.com/gin-contrib/sessions/blob/master/session_options_go1.11.go
	store.Options(sessions.Options{
		Path:     cookiePath,
		Secure:   cookieSecure,
		HttpOnly: true,
	})
	return sessions.Sessions(cookieName, store)
}
func EnhancedGinEngine() *gin.Engine {
	// r := gin.New()
	r := gin.Default() // with default middlewares
	r.SetTrustedProxies(trustedProxies)
	return r
}

// todo: plus most-used-passwords check
// https://github.com/danielmiessler/SecLists/blob/master/Passwords/2023-200_most_used_passwords.txt
func assertBadPassword() {
	ln := 10
	if len(adminPassword) < ln {
		log.Fatalf("Length of env.ADMIN_PASSWORD should be gte %d\n", ln)
	}
}

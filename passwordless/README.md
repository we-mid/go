# Passwordless

> 简单易用的 Passwordless-邮箱登录 SDK

```go
// main.go
import (
	"gitee.com/we-mid/SaaS-FooBar/auth"
	"gitee.com/we-mid/go/cookiesession"
	"gitee.com/we-mid/go/passwordless"
)

const fullHTTPS = false // todo
const needsCORS = true // todo

func init() {
	// Passwordless 邮箱登录
	auth.Init(passwordless.Options{
		CookieSession: cookiesession.Options{
			TTLSession:   3 * 24 * time.Hour,
			CookieSecure: fullHTTPS,
			CookiePath:   "/",
			CookieName:   "saas_foobar",
		},
		SaaSName:    "SaaS-FooBar",
		SaaSURL:     os.Getenv("SAAS_URL"),
		RoutePrefix: "/passwordless",
		EnableCORS:  needsCORS,
		TTLCode:     5 * time.Minute,
		LenCode:     6,
		OnAttempt: func(email string) {
			log.Printf("login attempt: email=%q\n", email)
		},
		OnVerify: func(email string, pass bool) {
			log.Printf("login verify: email=%q, pass=%v\n", email, pass)
		},
	})
}
```

```go
// auth/auth.go
import "gitee.com/we-mid/go/passwordless"

var pwdless *passwordless.Passwordless

func Init(options passwordless.Options) {
	pwdless = passwordless.New(options)
	pwdless.RegisterRoutes(http.DefaultServeMux)
}

func IsLoggedIn(r *http.Request) bool {
	userID := GetUserID(r)
	return userID != ""
}
func GetUserID(r *http.Request) string {
	email, err := pwdless.GetEmail(r)
	if err != nil {
		log.Println("Error while pwdless.GetEmail:", err)
	}
	return email
}
```

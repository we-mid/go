package passwordless

import (
	_ "embed"
	"fmt"
	"net/http"
	"net/mail"
	"time"

	"gitee.com/we-mid/go/bec_http"
	"gitee.com/we-mid/go/mailer"
	"gitee.com/we-mid/go/util"
)

const cleanInterval = 1 * time.Minute

// todo: i18n specify
const subjectEn = "Sign-in Verification"

//go:embed i18n/template_en.txt
var templateEn string

var (
	// 基于IP的安全限流 所有passwordless实例共用
	lAttempt = bec_http.NewIPRateLimit(time.Minute, 3)
	lVerify  = bec_http.NewIPRateLimit(time.Minute, 10)
	myErr429 = bec_http.NewStatusErrorf(429, "Please try again later.")
)

type verifyReq struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}
type verifyRes struct {
	Pass bool `json:"pass"`
}

func (p *Passwordless) HandleVerify(w http.ResponseWriter, r *http.Request) (any, error) {
	var req verifyReq
	if err := p.handleReq(w, r, &req); err != nil {
		return nil, err
	}
	if !lVerify.Allow(r) {
		return nil, myErr429
	}
	email, code := req.Email, req.Code
	var pass bool
	if _, ok := p.testUsers[email]; ok { // is test user
		if c, ok := p.testUsers[email]; ok && c == code {
			pass = true
		}
	} else {
		if b, ok := p.codeMap[email]; ok && b.code == code {
			pass = true
			p.mu.Lock()
			delete(p.codeMap, email)
			p.mu.Unlock()
		}
	}
	if pass {
		if err := p.bindEmail(w, email); err != nil {
			return nil, err
		}
	}
	go p.OnVerify(email, pass)

	return verifyRes{pass}, nil
}

type attemptReq struct {
	Email string `json:"email"`
}

func (p *Passwordless) HandleAttempt(w http.ResponseWriter, r *http.Request) (any, error) {
	var req attemptReq
	if err := p.handleReq(w, r, &req); err != nil {
		return nil, err
	}
	if !lAttempt.Allow(r) {
		return nil, myErr429
	}
	email := req.Email
	go p.OnAttempt(email)

	if _, ok := p.testUsers[email]; ok { // is test user
		return nil, nil // skip to store & send
	}
	var code string
	now := time.Now()
	if b, ok := p.codeMap[email]; ok && b.expireAt.Add(10*time.Second).After(now) {
		code = b.code
	} else {
		code = util.RandomCode(p.LenCode)
		binding := codeBinding{code, now.Add(p.TTLCode)}
		p.mu.Lock()
		p.codeMap[email] = binding
		p.mu.Unlock()
	}
	contact := mailer.SMTP_USER
	from := mail.Address{Name: p.SaaSName, Address: contact}
	to := mail.Address{Address: email}

	// subject := fmt.Sprintf("Sign-in Verification for %s", saasName)
	// subject := util.Ternary(p.MailSubject != "", p.MailSubject, subjectEn)
	subject := subjectEn

	// body := "Dear User,\nYour Verification Code is: 12345678 ."
	// template := util.Ternary(p.MailTemplate != "", p.MailTemplate, templateEn)
	template := templateEn
	body := fmt.Sprintf(template, p.SaaSName, code, p.SaaSURL, contact, p.SaaSName)

	err := mailer.SendMail(body, subject, []mail.Address{to}, &from)
	return nil, err
}

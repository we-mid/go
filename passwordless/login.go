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
	lAttempt = bec_http.NewIPRateLimit(time.Minute, 1)
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
	var pass bool
	if b, ok := p.codeMap[req.Email]; ok && b.code == req.Code {
		pass = true
		delete(p.codeMap, req.Email)
		if err := p.bindEmail(w, req.Email); err != nil {
			return nil, err
		}
	}
	go p.OnVerify(req.Email, pass)

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
	go p.OnAttempt(req.Email)

	code := util.RandomCode(p.LenCode)
	binding := codeBinding{code, time.Now().Add(p.TTLCode)}
	p.mu.Lock()
	p.codeMap[req.Email] = binding
	p.mu.Unlock()

	contact := mailer.SMTP_USER
	from := mail.Address{Name: p.SaaSName, Address: contact}
	to := mail.Address{Address: req.Email}

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

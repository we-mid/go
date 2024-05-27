package mailer

import (
	"crypto/tls"
	"fmt"
	"net/mail"
	"net/smtp"
	"os"
)

// SSL/TLS Email Example
// https://gist.github.com/chrisgillis/10888032
// https://bastengao.com/blog/2019/11/go-smtp-ssl.html

var (
	SMTP_HOST = os.Getenv("SMTP_HOST")
	SMTP_PORT = os.Getenv("SMTP_PORT")
	SMTP_USER = os.Getenv("SMTP_USER")
	SMTP_PASS = os.Getenv("SMTP_PASS")
)

func SendMail(from, to mail.Address, subject, body string) error {
	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	addr := fmt.Sprintf("%s:%s", SMTP_HOST, SMTP_PORT)
	// host, _, _ := net.SplitHostPort(servername)
	host := SMTP_HOST
	auth := smtp.PlainAuth("", SMTP_USER, SMTP_PASS, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return err
	}
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	// Auth
	if err = c.Auth(auth); err != nil {
		return err
	}
	// To && From
	if err = c.Mail(from.Address); err != nil {
		return err
	}
	if err = c.Rcpt(to.Address); err != nil {
		return err
	}
	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

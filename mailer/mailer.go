package mailer

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/mail"
	"net/smtp"
	"os"
	"strings"
)

// SSL/TLS Email Example
// https://gist.github.com/chrisgillis/10888032
// https://portswigger.net/kb/issues/00200800_smtp-header-injection (security)
// https://bastengao.com/blog/2019/11/go-smtp-ssl.html

var (
	// 定义发件人、收件人列表、SMTP服务器信息等
	SMTP_HOST = os.Getenv("SMTP_HOST")
	SMTP_PORT = os.Getenv("SMTP_PORT")
	SMTP_PASS = os.Getenv("SMTP_PASS")
	SMTP_USER = os.Getenv("SMTP_USER")
	SMTP_NICK = os.Getenv("SMTP_NICK")

	SmtpFrom = mail.Address{Name: SMTP_NICK, Address: SMTP_USER}
)

func SendMail(body, subject string, toList []mail.Address, from *mail.Address) error {
	if from == nil {
		from = &SmtpFrom
	}
	var toStrList []string
	for _, to := range toList {
		toStrList = append(toStrList, to.String())
	}
	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = strings.Join(toStrList, ", ")
	headers["Subject"] = subject

	// Setup message
	message := ""
	for k, v := range headers {
		line := fmt.Sprintf("%s: %s", k, v)
		// security assurance: RFC 5321
		if err := validateLine(line); err != nil {
			return err
		}
		message += line + "\r\n"
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
	defer c.Quit()
	// Auth
	if err = c.Auth(auth); err != nil {
		return err
	}
	// To && From
	if err = c.Mail(from.Address); err != nil {
		return err
	}
	for _, to := range toList {
		if err = c.Rcpt(to.Address); err != nil {
			return err
		}
	}
	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}
	defer w.Close()
	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}
	return nil
}

// validateLine checks to see if a line has CR or LF as per RFC 5321.
// https://github.com/golang/go/blob/87ec2c959c73e62bfae230ef7efca11ec2a90804/src/net/smtp/smtp.go#L429
func validateLine(line string) error {
	if strings.ContainsAny(line, "\n\r") {
		return errors.New("[mailer] smtp: A line must not contain CR or LF")
	}
	return nil
}

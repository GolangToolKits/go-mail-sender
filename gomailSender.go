package mailsender

import (
	"strconv"
	"strings"

	"github.com/go-mail/mail"
)

// GoMailSender GoMailSender
type GoMailSender struct {
	User     string
	Password string
	MailHost string
	Port     string
}

// SendMail SendMail
func (m *GoMailSender) SendMail(mailer *Mailer) bool {
	var rtn bool
	mm := mail.NewMessage()
	mm.SetHeader("From", mailer.SenderAddress)
	toAddrs := strings.Join(mailer.Recipients, ",")
	mm.SetHeader("To", toAddrs)
	mm.SetHeader("Subject", mailer.Subject)
	mm.SetBody("text/html", mailer.Body)
	intPort, err := strconv.Atoi(m.Port)
	if err == nil {
		d := mail.NewDialer(m.MailHost, intPort, m.User, m.Password)
		err := d.DialAndSend(mm)
		if err == nil {
			rtn = true
		}
	}
	return rtn
}

// New GetNew
func (m *GoMailSender) New() Sender {
	return m
}
